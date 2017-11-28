// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package sql

import (
	"fmt"
	"math"

	"golang.org/x/net/context"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/types"
)

// limitNode represents a node that limits the number of rows
// returned or only return them past a given number (offset).
type limitNode struct {
	plan       planNode
	countExpr  tree.TypedExpr
	offsetExpr tree.TypedExpr
	evaluated  bool
	count      int64
	offset     int64

	run limitRun
}

// limit constructs a limitNode based on the LIMIT and OFFSET clauses.
func (p *planner) Limit(ctx context.Context, n *tree.Limit) (*limitNode, error) {
	if n == nil || (n.Count == nil && n.Offset == nil) {
		// No LIMIT nor OFFSET; there is nothing special to do.
		return nil, nil
	}

	res := limitNode{}

	data := []struct {
		name string
		src  tree.Expr
		dst  *tree.TypedExpr
	}{
		{"LIMIT", n.Count, &res.countExpr},
		{"OFFSET", n.Offset, &res.offsetExpr},
	}

	for _, datum := range data {
		if datum.src != nil {
			if err := p.txCtx.AssertNoAggregationOrWindowing(
				datum.src, datum.name, p.session.SearchPath,
			); err != nil {
				return nil, err
			}

			normalized, err := p.analyzeExpr(ctx, datum.src, nil, tree.IndexedVarHelper{}, types.Int, true, datum.name)
			if err != nil {
				return nil, err
			}
			*datum.dst = normalized
		}
	}
	return &res, nil
}

// limitRun contains the state of limitNode during local execution.
type limitRun struct {
	rowIndex int64
}

func (n *limitNode) Start(params runParams) error {
	if err := n.plan.Start(params); err != nil {
		return err
	}

	return n.evalLimit(params.evalCtx)
}

func (n *limitNode) Next(params runParams) (bool, error) {
	// n.rowIndex is the 0-based index of the next row.
	// We don't do (n.rowIndex >= n.offset + n.count) to avoid overflow (count can be MaxInt64).
	if n.run.rowIndex-n.offset >= n.count {
		return false, nil
	}

	for {
		if next, err := n.plan.Next(params); !next {
			return false, err
		}

		n.run.rowIndex++
		if n.run.rowIndex > n.offset {
			// Row within limits, return it.
			break
		}

		// Fetch the next row.
	}
	return true, nil
}

func (n *limitNode) Values() tree.Datums { return n.plan.Values() }

func (n *limitNode) Close(ctx context.Context) {
	n.plan.Close(ctx)
}

// estimateLimit pre-computes the count and offset fields if they are constants,
// otherwise predefines them to be MaxInt64, 0. Used by index selection.
// This must be called after type checking and constant folding.
func (n *limitNode) estimateLimit() {
	n.count = math.MaxInt64
	n.offset = 0

	// Use simple integer datum if available.
	// The limit can be a simple DInt here either because it was
	// entered as such in the query, or as a result of constant
	// folding prior to type checking.

	if n.countExpr != nil {
		if i, ok := tree.AsDInt(n.countExpr); ok {
			n.count = int64(i)
		}
	}
	if n.offsetExpr != nil {
		if i, ok := tree.AsDInt(n.offsetExpr); ok {
			n.offset = int64(i)
		}
	}
}

// evalLimit evaluates the Count and Offset fields. If Count is missing, the
// value is MaxInt64. If Offset is missing, the value is 0
func (n *limitNode) evalLimit(evalCtx *tree.EvalContext) error {
	n.count = math.MaxInt64
	n.offset = 0

	data := []struct {
		name string
		src  tree.TypedExpr
		dst  *int64
	}{
		{"LIMIT", n.countExpr, &n.count},
		{"OFFSET", n.offsetExpr, &n.offset},
	}

	for _, datum := range data {
		if datum.src != nil {
			dstDatum, err := datum.src.Eval(evalCtx)
			if err != nil {
				return err
			}

			if dstDatum == tree.DNull {
				// Use the default value.
				continue
			}

			dstDInt := tree.MustBeDInt(dstDatum)
			val := int64(dstDInt)
			if val < 0 {
				return fmt.Errorf("negative value for %s", datum.name)
			}
			*datum.dst = val
		}
	}
	n.evaluated = true
	return nil
}

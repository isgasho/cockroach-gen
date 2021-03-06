// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
)

// TODO(yuzefovich): add benchmarks.

// NewRowNumberOperator creates a new Operator that computes window function
// ROW_NUMBER. outputColIdx specifies in which coldata.Vec the operator should
// put its output (if there is no such column, a new column is appended).
func NewRowNumberOperator(
	allocator *Allocator, input Operator, outputColIdx int, partitionColIdx int,
) Operator {
	base := rowNumberBase{
		OneInputNode:    NewOneInputNode(input),
		allocator:       allocator,
		outputColIdx:    outputColIdx,
		partitionColIdx: partitionColIdx,
	}
	if partitionColIdx == -1 {
		return &rowNumberNoPartitionOp{base}
	}
	return &rowNumberWithPartitionOp{base}
}

// rowNumberBase extracts common fields and common initialization of two
// variations of row number operators. Note that it is not an operator itself
// and should not be used directly.
type rowNumberBase struct {
	OneInputNode
	allocator       *Allocator
	outputColIdx    int
	partitionColIdx int

	rowNumber int64
}

func (r *rowNumberBase) Init() {
	r.Input().Init()
}

type rowNumberNoPartitionOp struct {
	rowNumberBase
}

var _ Operator = &rowNumberNoPartitionOp{}

func (r *rowNumberNoPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)
	if batch.Length() == 0 {
		return coldata.ZeroBatch
	}
	r.allocator.MaybeAddColumn(batch, coltypes.Int64, r.outputColIdx)

	rowNumberCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			r.rowNumber++
			rowNumberCol[sel[i]] = r.rowNumber
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			r.rowNumber++
			rowNumberCol[i] = r.rowNumber
		}
	}
	return batch
}

type rowNumberWithPartitionOp struct {
	rowNumberBase
}

var _ Operator = &rowNumberWithPartitionOp{}

func (r *rowNumberWithPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)
	if batch.Length() == 0 {
		return coldata.ZeroBatch
	}
	r.allocator.MaybeAddColumn(batch, coltypes.Bool, r.partitionColIdx)
	r.allocator.MaybeAddColumn(batch, coltypes.Int64, r.outputColIdx)

	partitionCol := batch.ColVec(r.partitionColIdx).Bool()
	rowNumberCol := batch.ColVec(r.outputColIdx).Int64()
	sel := batch.Selection()
	if sel != nil {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[sel[i]] {
				r.rowNumber = 1
			}
			r.rowNumber++
			rowNumberCol[sel[i]] = r.rowNumber
		}
	} else {
		for i := uint16(0); i < batch.Length(); i++ {
			if partitionCol[i] {
				r.rowNumber = 0
			}
			r.rowNumber++
			rowNumberCol[i] = r.rowNumber
		}
	}
	return batch
}

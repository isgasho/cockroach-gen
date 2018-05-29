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

package storage

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/cockroachdb/cockroach/pkg/config"
	"github.com/cockroachdb/cockroach/pkg/gossip"
	"github.com/cockroachdb/cockroach/pkg/internal/client"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)

const (
	// splitQueueTimerDuration is the duration between splits of queued ranges.
	splitQueueTimerDuration = 0 // zero duration to process splits greedily.

	// splitQueuePurgatoryCheckInterval is the interval at which replicas in
	// purgatory make split attempts. Purgatory is used by the splitQueue to
	// store ranges that are large enough to require a split but are
	// unsplittable because they do not contain a suitable split key. Purgatory
	// prevents them from repeatedly attempting to split at an unbounded rate.
	splitQueuePurgatoryCheckInterval = 1 * time.Minute

	// splits should be relatively isolated, other than requiring expensive
	// RocksDB scans over part of the splitting range to recompute stats. We
	// allow a limitted number of splits to be processed at once.
	splitQueueConcurrency = 4
)

// splitQueue manages a queue of ranges slated to be split due to size
// or along intersecting zone config boundaries.
type splitQueue struct {
	*baseQueue
	db       *client.DB
	purgChan <-chan time.Time
}

// newSplitQueue returns a new instance of splitQueue.
func newSplitQueue(store *Store, db *client.DB, gossip *gossip.Gossip) *splitQueue {
	var purgChan <-chan time.Time
	if c := store.TestingKnobs().SplitQueuePurgatoryChan; c != nil {
		purgChan = c
	} else {
		purgTicker := time.NewTicker(splitQueuePurgatoryCheckInterval)
		purgChan = purgTicker.C
	}

	sq := &splitQueue{
		db:       db,
		purgChan: purgChan,
	}
	sq.baseQueue = newBaseQueue(
		"split", sq, store, gossip,
		queueConfig{
			maxSize:              defaultQueueMaxSize,
			maxConcurrency:       splitQueueConcurrency,
			needsLease:           true,
			needsSystemConfig:    true,
			acceptsUnsplitRanges: true,
			successes:            store.metrics.SplitQueueSuccesses,
			failures:             store.metrics.SplitQueueFailures,
			pending:              store.metrics.SplitQueuePending,
			processingNanos:      store.metrics.SplitQueueProcessingNanos,
			purgatory:            store.metrics.SplitQueuePurgatory,
		},
	)
	return sq
}

// shouldQueue determines whether a range should be queued for
// splitting. This is true if the range is intersected by a zone config
// prefix or if the range's size in bytes exceeds the limit for the zone.
func (sq *splitQueue) shouldQueue(
	ctx context.Context, now hlc.Timestamp, repl *Replica, sysCfg config.SystemConfig,
) (shouldQ bool, priority float64) {
	desc := repl.Desc()
	if sysCfg.NeedsSplit(desc.StartKey, desc.EndKey) {
		// Set priority to 1 in the event the range is split by zone configs.
		priority = 1
		shouldQ = true
	}

	// Add priority based on the size of range compared to the max
	// size for the zone it's in.
	if ratio := float64(repl.GetMVCCStats().Total()) / float64(repl.GetMaxBytes()); ratio > 1 {
		priority += ratio
		shouldQ = true
	}
	return
}

// unsplittableRangeError indicates that a split attempt failed because a no
// suitable split key could be found.
type unsplittableRangeError struct{}

func (unsplittableRangeError) Error() string         { return "could not find valid split key" }
func (unsplittableRangeError) purgatoryErrorMarker() {}

var _ purgatoryError = unsplittableRangeError{}

// process synchronously invokes admin split for each proposed split key.
func (sq *splitQueue) process(ctx context.Context, r *Replica, sysCfg config.SystemConfig) error {
	err := sq.processAttempt(ctx, r, sysCfg)
	switch errors.Cause(err).(type) {
	case nil:
	case *roachpb.ConditionFailedError:
		// ConditionFailedErrors are an expected outcome for range split
		// attempts because splits can race with other descriptor modifications.
		// On seeing a ConditionFailedError, don't return an error and enqueue
		// this replica again in case it still needs to be split.
		log.Infof(ctx, "split saw concurrent descriptor modification; maybe retrying")
		sq.MaybeAdd(r, sq.store.Clock().Now())
	default:
		return err
	}
	return nil
}

func (sq *splitQueue) processAttempt(
	ctx context.Context, r *Replica, sysCfg config.SystemConfig,
) error {
	// First handle case of splitting due to zone config maps.
	desc := r.Desc()
	if splitKey := sysCfg.ComputeSplitKey(desc.StartKey, desc.EndKey); splitKey != nil {
		if _, err := r.adminSplitWithDescriptor(
			ctx,
			roachpb.AdminSplitRequest{
				RequestHeader: roachpb.RequestHeader{
					Key: splitKey.AsRawKey(),
				},
				SplitKey: splitKey.AsRawKey(),
			},
			desc,
		); err != nil {
			return errors.Wrapf(err, "unable to split %s at key %q", r, splitKey)
		}
		return nil
	}

	// Next handle case of splitting due to size. Note that we don't perform
	// size-based splitting if maxBytes is 0 (happens in certain test
	// situations).
	size := r.GetMVCCStats().Total()
	maxBytes := r.GetMaxBytes()
	if maxBytes > 0 && float64(size)/float64(maxBytes) > 1 {
		_, err := r.adminSplitWithDescriptor(
			ctx,
			roachpb.AdminSplitRequest{},
			desc,
		)
		return err
	}
	return nil
}

// timer returns interval between processing successive queued splits.
func (*splitQueue) timer(_ time.Duration) time.Duration {
	return splitQueueTimerDuration
}

// purgatoryChan returns the split queue's purgatory channel.
func (sq *splitQueue) purgatoryChan() <-chan time.Time {
	return sq.purgChan
}

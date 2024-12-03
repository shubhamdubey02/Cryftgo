// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowmantest

import (
	"context"

	"github.com/shubhamdubey02/cryftgo/database"
	"github.com/shubhamdubey02/cryftgo/ids"
	"github.com/shubhamdubey02/cryftgo/snow/snowtest"
)

func MakeLastAcceptedBlockF(blks ...[]*Block) func(context.Context) (ids.ID, error) {
	return func(context.Context) (ids.ID, error) {
		var (
			highestHeight uint64
			highestID     ids.ID
		)
		for _, blkSlice := range blks {
			for _, blk := range blkSlice {
				if blk.Status != snowtest.Accepted {
					continue
				}

				if height := blk.Height(); height >= highestHeight {
					highestHeight = height
					highestID = blk.ID()
				}
			}
		}
		return highestID, nil
	}
}

func MakeGetBlockIDAtHeightF(blks ...[]*Block) func(context.Context, uint64) (ids.ID, error) {
	return func(_ context.Context, height uint64) (ids.ID, error) {
		for _, blkSlice := range blks {
			for _, blk := range blkSlice {
				if blk.Status != snowtest.Accepted {
					continue
				}

				if height == blk.Height() {
					return blk.ID(), nil
				}
			}
		}
		return ids.Empty, database.ErrNotFound
	}
}

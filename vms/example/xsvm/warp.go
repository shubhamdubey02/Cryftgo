// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package xsvm

import (
	"context"

	"github.com/shubhamdubey02/cryftgo/network/p2p/acp118"
	"github.com/shubhamdubey02/cryftgo/snow/engine/common"
	"github.com/shubhamdubey02/cryftgo/vms/platformvm/warp"
)

var _ acp118.Verifier = (*acp118Verifier)(nil)

// acp118Verifier allows signing all warp messages
type acp118Verifier struct{}

func (acp118Verifier) Verify(context.Context, *warp.UnsignedMessage, []byte) *common.AppError {
	return nil
}

// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"math"

	"github.com/shubhamdubey02/cryftgo/codec"
	"github.com/shubhamdubey02/cryftgo/codec/linearcodec"
)

const CodecVersion = 0

var Codec codec.Manager

func init() {
	lc := linearcodec.NewDefault()
	Codec = codec.NewManager(math.MaxInt32)

	err := Codec.RegisterCodec(CodecVersion, lc)
	if err != nil {
		panic(err)
	}
}

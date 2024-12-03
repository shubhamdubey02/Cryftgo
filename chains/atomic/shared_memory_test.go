// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic_test

import (
	"testing"

	"github.com/shubhamdubey02/cryftgo/chains/atomic/atomictest"
	"github.com/shubhamdubey02/cryftgo/database/memdb"
	"github.com/shubhamdubey02/cryftgo/database/prefixdb"
	"github.com/shubhamdubey02/cryftgo/ids"

	. "github.com/shubhamdubey02/cryftgo/chains/atomic"
)

func TestSharedMemory(t *testing.T) {
	chainID0 := ids.GenerateTestID()
	chainID1 := ids.GenerateTestID()

	for _, test := range atomictest.SharedMemoryTests {
		baseDB := memdb.New()

		memoryDB := prefixdb.New([]byte{0}, baseDB)
		testDB := prefixdb.New([]byte{1}, baseDB)

		m := NewMemory(memoryDB)

		sm0 := m.NewSharedMemory(chainID0)
		sm1 := m.NewSharedMemory(chainID1)

		test(t, chainID0, chainID1, sm0, sm1, testDB)
	}
}
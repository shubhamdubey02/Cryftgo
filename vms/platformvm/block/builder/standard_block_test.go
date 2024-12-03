// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/shubhamdubey02/cryftgo/chains/atomic"
	"github.com/shubhamdubey02/cryftgo/database/prefixdb"
	"github.com/shubhamdubey02/cryftgo/ids"
	"github.com/shubhamdubey02/cryftgo/upgrade/upgradetest"
	"github.com/shubhamdubey02/cryftgo/utils/units"
	"github.com/shubhamdubey02/cryftgo/vms/components/avax"
	"github.com/shubhamdubey02/cryftgo/vms/platformvm/genesis/genesistest"
	"github.com/shubhamdubey02/cryftgo/vms/platformvm/status"
	"github.com/shubhamdubey02/cryftgo/vms/platformvm/txs"
	"github.com/shubhamdubey02/cryftgo/vms/secp256k1fx"
)

func TestAtomicTxImports(t *testing.T) {
	require := require.New(t)

	env := newEnvironment(t, upgradetest.Latest)
	env.ctx.Lock.Lock()
	defer env.ctx.Lock.Unlock()

	addr := genesistest.DefaultFundedKeys[0].Address()
	owner := &secp256k1fx.OutputOwners{
		Threshold: 1,
		Addrs:     []ids.ShortID{addr},
	}

	m := atomic.NewMemory(prefixdb.New([]byte{5}, env.baseDB))

	env.msm.SharedMemory = m.NewSharedMemory(env.ctx.ChainID)
	peerSharedMemory := m.NewSharedMemory(env.ctx.XChainID)
	utxo := &avax.UTXO{
		UTXOID: avax.UTXOID{
			TxID:        ids.GenerateTestID(),
			OutputIndex: 1,
		},
		Asset: avax.Asset{ID: env.ctx.AVAXAssetID},
		Out: &secp256k1fx.TransferOutput{
			Amt:          70 * units.MicroAvax,
			OutputOwners: *owner,
		},
	}
	utxoBytes, err := txs.Codec.Marshal(txs.CodecVersion, utxo)
	require.NoError(err)

	inputID := utxo.InputID()
	require.NoError(peerSharedMemory.Apply(map[ids.ID]*atomic.Requests{
		env.ctx.ChainID: {PutRequests: []*atomic.Element{{
			Key:   inputID[:],
			Value: utxoBytes,
			Traits: [][]byte{
				addr.Bytes(),
			},
		}}},
	}))

	wallet := newWallet(t, env, walletConfig{})

	tx, err := wallet.IssueImportTx(
		env.ctx.XChainID,
		owner,
	)
	require.NoError(err)

	require.NoError(env.Builder.Add(tx))
	b, err := env.Builder.BuildBlock(context.Background())
	require.NoError(err)
	// Test multiple verify calls work
	require.NoError(b.Verify(context.Background()))
	require.NoError(b.Accept(context.Background()))
	_, txStatus, err := env.state.GetTx(tx.ID())
	require.NoError(err)
	// Ensure transaction is in the committed state
	require.Equal(status.Committed, txStatus)
}
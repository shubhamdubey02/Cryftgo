// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/shubhamdubey02/cryftgo/codec"
	"github.com/shubhamdubey02/cryftgo/codec/codecmock"
	"github.com/shubhamdubey02/cryftgo/database/memdb"
	"github.com/shubhamdubey02/cryftgo/database/versiondb"
	"github.com/shubhamdubey02/cryftgo/ids"
	"github.com/shubhamdubey02/cryftgo/snow"
	"github.com/shubhamdubey02/cryftgo/snow/consensus/snowman"
	"github.com/shubhamdubey02/cryftgo/snow/engine/common"
	"github.com/shubhamdubey02/cryftgo/utils/constants"
	"github.com/shubhamdubey02/cryftgo/utils/crypto/secp256k1"
	"github.com/shubhamdubey02/cryftgo/utils/logging"
	"github.com/shubhamdubey02/cryftgo/utils/timer/mockable"
	"github.com/shubhamdubey02/cryftgo/vms/avm/block"
	"github.com/shubhamdubey02/cryftgo/vms/avm/block/executor/executormock"
	"github.com/shubhamdubey02/cryftgo/vms/avm/fxs"
	"github.com/shubhamdubey02/cryftgo/vms/avm/metrics"
	"github.com/shubhamdubey02/cryftgo/vms/avm/state"
	"github.com/shubhamdubey02/cryftgo/vms/avm/state/statemock"
	"github.com/shubhamdubey02/cryftgo/vms/avm/txs"
	"github.com/shubhamdubey02/cryftgo/vms/avm/txs/mempool"
	"github.com/shubhamdubey02/cryftgo/vms/avm/txs/mempool/mempoolmock"
	"github.com/shubhamdubey02/cryftgo/vms/avm/txs/txsmock"
	"github.com/shubhamdubey02/cryftgo/vms/components/avax"
	"github.com/shubhamdubey02/cryftgo/vms/secp256k1fx"

	blkexecutor "github.com/shubhamdubey02/cryftgo/vms/avm/block/executor"
	txexecutor "github.com/shubhamdubey02/cryftgo/vms/avm/txs/executor"
)

const trackChecksums = false

var (
	errTest = errors.New("test error")
	chainID = ids.GenerateTestID()
	keys    = secp256k1.TestKeys()
)

func TestBuilderBuildBlock(t *testing.T) {
	type test struct {
		name        string
		builderFunc func(*gomock.Controller) Builder
		expectedErr error
	}

	tests := []test{
		{
			name: "can't get stateless block",
			builderFunc: func(ctrl *gomock.Controller) Builder {
				preferredID := ids.GenerateTestID()
				manager := executormock.NewManager(ctrl)
				manager.EXPECT().Preferred().Return(preferredID)
				manager.EXPECT().GetStatelessBlock(preferredID).Return(nil, errTest)

				mempool := mempoolmock.NewMempool(ctrl)
				mempool.EXPECT().RequestBuildBlock()

				return New(
					&txexecutor.Backend{
						Ctx: &snow.Context{
							Log: logging.NoLog{},
						},
					},
					manager,
					&mockable.Clock{},
					mempool,
				)
			},
			expectedErr: errTest,
		},
		{
			name: "can't get preferred diff",
			builderFunc: func(ctrl *gomock.Controller) Builder {
				preferredID := ids.GenerateTestID()
				preferredHeight := uint64(1337)
				preferredTimestamp := time.Now()
				preferredBlock := block.NewMockBlock(ctrl)
				preferredBlock.EXPECT().Height().Return(preferredHeight)
				preferredBlock.EXPECT().Timestamp().Return(preferredTimestamp)

				manager := executormock.NewManager(ctrl)
				manager.EXPECT().Preferred().Return(preferredID)
				manager.EXPECT().GetStatelessBlock(preferredID).Return(preferredBlock, nil)
				manager.EXPECT().GetState(preferredID).Return(nil, false)

				mempool := mempoolmock.NewMempool(ctrl)
				mempool.EXPECT().RequestBuildBlock()

				return New(
					&txexecutor.Backend{
						Ctx: &snow.Context{
							Log: logging.NoLog{},
						},
					},
					manager,
					&mockable.Clock{},
					mempool,
				)
			},
			expectedErr: state.ErrMissingParentState,
		},
		{
			name: "tx fails semantic verification",
			builderFunc: func(ctrl *gomock.Controller) Builder {
				preferredID := ids.GenerateTestID()
				preferredHeight := uint64(1337)
				preferredTimestamp := time.Now()
				preferredBlock := block.NewMockBlock(ctrl)
				preferredBlock.EXPECT().Height().Return(preferredHeight)
				preferredBlock.EXPECT().Timestamp().Return(preferredTimestamp)

				preferredState := statemock.NewChain(ctrl)
				preferredState.EXPECT().GetLastAccepted().Return(preferredID)
				preferredState.EXPECT().GetTimestamp().Return(preferredTimestamp)

				manager := executormock.NewManager(ctrl)
				manager.EXPECT().Preferred().Return(preferredID)
				manager.EXPECT().GetStatelessBlock(preferredID).Return(preferredBlock, nil)
				manager.EXPECT().GetState(preferredID).Return(preferredState, true)

				unsignedTx := txsmock.NewUnsignedTx(ctrl)
				unsignedTx.EXPECT().Visit(gomock.Any()).Return(errTest) // Fail semantic verification
				tx := &txs.Tx{Unsigned: unsignedTx}

				mempool := mempoolmock.NewMempool(ctrl)
				mempool.EXPECT().Peek().Return(tx, true)
				mempool.EXPECT().Remove([]*txs.Tx{tx})
				mempool.EXPECT().MarkDropped(tx.ID(), errTest)
				// Second loop iteration
				mempool.EXPECT().Peek().Return(nil, false)
				mempool.EXPECT().RequestBuildBlock()

				return New(
					&txexecutor.Backend{
						Ctx: &snow.Context{
							Log: logging.NoLog{},
						},
					},
					manager,
					&mockable.Clock{},
					mempool,
				)
			},
			expectedErr: ErrNoTransactions, // The only tx was invalid
		},
		{
			name: "tx fails execution",
			builderFunc: func(ctrl *gomock.Controller) Builder {
				preferredID := ids.GenerateTestID()
				preferredHeight := uint64(1337)
				preferredTimestamp := time.Now()
				preferredBlock := block.NewMockBlock(ctrl)
				preferredBlock.EXPECT().Height().Return(preferredHeight)
				preferredBlock.EXPECT().Timestamp().Return(preferredTimestamp)

				preferredState := statemock.NewChain(ctrl)
				preferredState.EXPECT().GetLastAccepted().Return(preferredID)
				preferredState.EXPECT().GetTimestamp().Return(preferredTimestamp)

				manager := executormock.NewManager(ctrl)
				manager.EXPECT().Preferred().Return(preferredID)
				manager.EXPECT().GetStatelessBlock(preferredID).Return(preferredBlock, nil)
				manager.EXPECT().GetState(preferredID).Return(preferredState, true)

				unsignedTx := txsmock.NewUnsignedTx(ctrl)
				unsignedTx.EXPECT().Visit(gomock.Any()).Return(nil)     // Pass semantic verification
				unsignedTx.EXPECT().Visit(gomock.Any()).Return(errTest) // Fail execution
				tx := &txs.Tx{Unsigned: unsignedTx}

				mempool := mempoolmock.NewMempool(ctrl)
				mempool.EXPECT().Peek().Return(tx, true)
				mempool.EXPECT().Remove([]*txs.Tx{tx})
				mempool.EXPECT().MarkDropped(tx.ID(), errTest)
				// Second loop iteration
				mempool.EXPECT().Peek().Return(nil, false)
				mempool.EXPECT().RequestBuildBlock()

				return New(
					&txexecutor.Backend{
						Ctx: &snow.Context{
							Log: logging.NoLog{},
						},
					},
					manager,
					&mockable.Clock{},
					mempool,
				)
			},
			expectedErr: ErrNoTransactions, // The only tx was invalid
		},
		{
			name: "tx has non-unique inputs",
			builderFunc: func(ctrl *gomock.Controller) Builder {
				preferredID := ids.GenerateTestID()
				preferredHeight := uint64(1337)
				preferredTimestamp := time.Now()
				preferredBlock := block.NewMockBlock(ctrl)
				preferredBlock.EXPECT().Height().Return(preferredHeight)
				preferredBlock.EXPECT().Timestamp().Return(preferredTimestamp)

				preferredState := statemock.NewChain(ctrl)
				preferredState.EXPECT().GetLastAccepted().Return(preferredID)
				preferredState.EXPECT().GetTimestamp().Return(preferredTimestamp)

				manager := executormock.NewManager(ctrl)
				manager.EXPECT().Preferred().Return(preferredID)
				manager.EXPECT().GetStatelessBlock(preferredID).Return(preferredBlock, nil)
				manager.EXPECT().GetState(preferredID).Return(preferredState, true)
				manager.EXPECT().VerifyUniqueInputs(preferredID, gomock.Any()).Return(errTest)

				unsignedTx := txsmock.NewUnsignedTx(ctrl)
				unsignedTx.EXPECT().Visit(gomock.Any()).Return(nil) // Pass semantic verification
				unsignedTx.EXPECT().Visit(gomock.Any()).Return(nil) // Pass execution
				tx := &txs.Tx{Unsigned: unsignedTx}

				mempool := mempoolmock.NewMempool(ctrl)
				mempool.EXPECT().Peek().Return(tx, true)
				mempool.EXPECT().Remove([]*txs.Tx{tx})
				mempool.EXPECT().MarkDropped(tx.ID(), errTest)
				// Second loop iteration
				mempool.EXPECT().Peek().Return(nil, false)
				mempool.EXPECT().RequestBuildBlock()

				return New(
					&txexecutor.Backend{
						Ctx: &snow.Context{
							Log: logging.NoLog{},
						},
					},
					manager,
					&mockable.Clock{},
					mempool,
				)
			},
			expectedErr: ErrNoTransactions, // The only tx was invalid
		},
		{
			name: "txs consume same input",
			builderFunc: func(ctrl *gomock.Controller) Builder {
				preferredID := ids.GenerateTestID()
				preferredHeight := uint64(1337)
				preferredTimestamp := time.Now()
				preferredBlock := block.NewMockBlock(ctrl)
				preferredBlock.EXPECT().Height().Return(preferredHeight)
				preferredBlock.EXPECT().Timestamp().Return(preferredTimestamp)

				preferredState := statemock.NewChain(ctrl)
				preferredState.EXPECT().GetLastAccepted().Return(preferredID)
				preferredState.EXPECT().GetTimestamp().Return(preferredTimestamp)

				// tx1 and tx2 both consume [inputID].
				// tx1 is added to the block first, so tx2 should be dropped.
				inputID := ids.GenerateTestID()
				unsignedTx1 := txsmock.NewUnsignedTx(ctrl)
				unsignedTx1.EXPECT().Visit(gomock.Any()).Return(nil)  // Pass semantic verification
				unsignedTx1.EXPECT().Visit(gomock.Any()).DoAndReturn( // Pass execution
					func(visitor txs.Visitor) error {
						require.IsType(t, &txexecutor.Executor{}, visitor)
						executor := visitor.(*txexecutor.Executor)
						executor.Inputs.Add(inputID)
						return nil
					},
				)
				unsignedTx1.EXPECT().SetBytes(gomock.Any()).AnyTimes()
				tx1 := &txs.Tx{Unsigned: unsignedTx1}
				// Set the bytes of tx1 to something other than nil
				// so we can check that the remainingSize is updated
				tx1Bytes := []byte{1, 2, 3}
				tx1.SetBytes(nil, tx1Bytes)

				unsignedTx2 := txsmock.NewUnsignedTx(ctrl)
				unsignedTx2.EXPECT().Visit(gomock.Any()).Return(nil)  // Pass semantic verification
				unsignedTx2.EXPECT().Visit(gomock.Any()).DoAndReturn( // Pass execution
					func(visitor txs.Visitor) error {
						require.IsType(t, &txexecutor.Executor{}, visitor)
						executor := visitor.(*txexecutor.Executor)
						executor.Inputs.Add(inputID)
						return nil
					},
				)
				tx2 := &txs.Tx{Unsigned: unsignedTx2}

				manager := executormock.NewManager(ctrl)
				manager.EXPECT().Preferred().Return(preferredID)
				manager.EXPECT().GetStatelessBlock(preferredID).Return(preferredBlock, nil)
				manager.EXPECT().GetState(preferredID).Return(preferredState, true)
				manager.EXPECT().VerifyUniqueInputs(preferredID, gomock.Any()).Return(nil)
				// Assert created block has one tx, tx1,
				// and other fields are set correctly.
				manager.EXPECT().NewBlock(gomock.Any()).DoAndReturn(
					func(block *block.StandardBlock) snowman.Block {
						require.Len(t, block.Transactions, 1)
						require.Equal(t, tx1, block.Transactions[0])
						require.Equal(t, preferredHeight+1, block.Height())
						require.Equal(t, preferredID, block.Parent())
						return nil
					},
				)

				mempool := mempoolmock.NewMempool(ctrl)
				mempool.EXPECT().Peek().Return(tx1, true)
				mempool.EXPECT().Remove([]*txs.Tx{tx1})
				// Second loop iteration
				mempool.EXPECT().Peek().Return(tx2, true)
				mempool.EXPECT().Remove([]*txs.Tx{tx2})
				mempool.EXPECT().MarkDropped(tx2.ID(), blkexecutor.ErrConflictingBlockTxs)
				// Third loop iteration
				mempool.EXPECT().Peek().Return(nil, false)
				mempool.EXPECT().RequestBuildBlock()

				// To marshal the tx/block
				codec := codecmock.NewManager(ctrl)
				codec.EXPECT().Marshal(gomock.Any(), gomock.Any()).Return([]byte{1, 2, 3}, nil).AnyTimes()
				codec.EXPECT().Size(gomock.Any(), gomock.Any()).Return(2, nil).AnyTimes()

				return New(
					&txexecutor.Backend{
						Codec: codec,
						Ctx: &snow.Context{
							Log: logging.NoLog{},
						},
					},
					manager,
					&mockable.Clock{},
					mempool,
				)
			},
			expectedErr: nil,
		},
		{
			name: "preferred timestamp after now",
			builderFunc: func(ctrl *gomock.Controller) Builder {
				preferredID := ids.GenerateTestID()
				preferredHeight := uint64(1337)
				preferredTimestamp := time.Now()
				preferredBlock := block.NewMockBlock(ctrl)
				preferredBlock.EXPECT().Height().Return(preferredHeight)
				preferredBlock.EXPECT().Timestamp().Return(preferredTimestamp)

				// Clock reads just before the preferred timestamp.
				// Created block should have the preferred timestamp since it's later.
				clock := &mockable.Clock{}
				clock.Set(preferredTimestamp.Add(-2 * time.Second))

				preferredState := statemock.NewChain(ctrl)
				preferredState.EXPECT().GetLastAccepted().Return(preferredID)
				preferredState.EXPECT().GetTimestamp().Return(preferredTimestamp)

				manager := executormock.NewManager(ctrl)
				manager.EXPECT().Preferred().Return(preferredID)
				manager.EXPECT().GetStatelessBlock(preferredID).Return(preferredBlock, nil)
				manager.EXPECT().GetState(preferredID).Return(preferredState, true)
				manager.EXPECT().VerifyUniqueInputs(preferredID, gomock.Any()).Return(nil)
				// Assert that the created block has the right timestamp
				manager.EXPECT().NewBlock(gomock.Any()).DoAndReturn(
					func(block *block.StandardBlock) snowman.Block {
						require.Equal(t, preferredTimestamp.Unix(), block.Timestamp().Unix())
						return nil
					},
				)

				inputID := ids.GenerateTestID()
				unsignedTx := txsmock.NewUnsignedTx(ctrl)
				unsignedTx.EXPECT().Visit(gomock.Any()).Return(nil)  // Pass semantic verification
				unsignedTx.EXPECT().Visit(gomock.Any()).DoAndReturn( // Pass execution
					func(visitor txs.Visitor) error {
						require.IsType(t, &txexecutor.Executor{}, visitor)
						executor := visitor.(*txexecutor.Executor)
						executor.Inputs.Add(inputID)
						return nil
					},
				)
				unsignedTx.EXPECT().SetBytes(gomock.Any()).AnyTimes()
				tx := &txs.Tx{Unsigned: unsignedTx}

				mempool := mempoolmock.NewMempool(ctrl)
				mempool.EXPECT().Peek().Return(tx, true)
				mempool.EXPECT().Remove([]*txs.Tx{tx})
				// Second loop iteration
				mempool.EXPECT().Peek().Return(nil, false)
				mempool.EXPECT().RequestBuildBlock()

				// To marshal the tx/block
				codec := codecmock.NewManager(ctrl)
				codec.EXPECT().Marshal(gomock.Any(), gomock.Any()).Return([]byte{1, 2, 3}, nil).AnyTimes()
				codec.EXPECT().Size(gomock.Any(), gomock.Any()).Return(2, nil).AnyTimes()

				return New(
					&txexecutor.Backend{
						Codec: codec,
						Ctx: &snow.Context{
							Log: logging.NoLog{},
						},
					},
					manager,
					clock,
					mempool,
				)
			},
			expectedErr: nil,
		},
		{
			name: "preferred timestamp before now",
			builderFunc: func(ctrl *gomock.Controller) Builder {
				preferredID := ids.GenerateTestID()
				preferredHeight := uint64(1337)
				// preferred block's timestamp is after the time reported by clock
				now := time.Now()
				preferredTimestamp := now.Add(-2 * time.Second)
				preferredBlock := block.NewMockBlock(ctrl)
				preferredBlock.EXPECT().Height().Return(preferredHeight)
				preferredBlock.EXPECT().Timestamp().Return(preferredTimestamp)

				// Clock reads after the preferred timestamp.
				// Created block should have [now] timestamp since it's later.
				clock := &mockable.Clock{}
				clock.Set(now)

				preferredState := statemock.NewChain(ctrl)
				preferredState.EXPECT().GetLastAccepted().Return(preferredID)
				preferredState.EXPECT().GetTimestamp().Return(preferredTimestamp)

				manager := executormock.NewManager(ctrl)
				manager.EXPECT().Preferred().Return(preferredID)
				manager.EXPECT().GetStatelessBlock(preferredID).Return(preferredBlock, nil)
				manager.EXPECT().GetState(preferredID).Return(preferredState, true)
				manager.EXPECT().VerifyUniqueInputs(preferredID, gomock.Any()).Return(nil)
				// Assert that the created block has the right timestamp
				manager.EXPECT().NewBlock(gomock.Any()).DoAndReturn(
					func(block *block.StandardBlock) snowman.Block {
						require.Equal(t, now.Unix(), block.Timestamp().Unix())
						return nil
					},
				)

				inputID := ids.GenerateTestID()
				unsignedTx := txsmock.NewUnsignedTx(ctrl)
				unsignedTx.EXPECT().Visit(gomock.Any()).Return(nil)  // Pass semantic verification
				unsignedTx.EXPECT().Visit(gomock.Any()).DoAndReturn( // Pass execution
					func(visitor txs.Visitor) error {
						require.IsType(t, &txexecutor.Executor{}, visitor)
						executor := visitor.(*txexecutor.Executor)
						executor.Inputs.Add(inputID)
						return nil
					},
				)
				unsignedTx.EXPECT().SetBytes(gomock.Any()).AnyTimes()
				tx := &txs.Tx{Unsigned: unsignedTx}

				mempool := mempoolmock.NewMempool(ctrl)
				mempool.EXPECT().Peek().Return(tx, true)
				mempool.EXPECT().Remove([]*txs.Tx{tx})
				// Second loop iteration
				mempool.EXPECT().Peek().Return(nil, false)
				mempool.EXPECT().RequestBuildBlock()

				// To marshal the tx/block
				codec := codecmock.NewManager(ctrl)
				codec.EXPECT().Marshal(gomock.Any(), gomock.Any()).Return([]byte{1, 2, 3}, nil).AnyTimes()
				codec.EXPECT().Size(gomock.Any(), gomock.Any()).Return(2, nil).AnyTimes()

				return New(
					&txexecutor.Backend{
						Codec: codec,
						Ctx: &snow.Context{
							Log: logging.NoLog{},
						},
					},
					manager,
					clock,
					mempool,
				)
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			builder := tt.builderFunc(ctrl)
			_, err := builder.BuildBlock(context.Background())
			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func TestBlockBuilderAddLocalTx(t *testing.T) {
	transactions := createTxs()

	require := require.New(t)

	registerer := prometheus.NewRegistry()
	toEngine := make(chan common.Message, 100)
	mempool, err := mempool.New("mempool", registerer, toEngine)
	require.NoError(err)
	// add a tx to the mempool
	tx := transactions[0]
	txID := tx.ID()
	require.NoError(mempool.Add(tx))

	_, ok := mempool.Get(txID)
	require.True(ok)

	parser, err := block.NewParser(
		[]fxs.Fx{
			&secp256k1fx.Fx{},
		},
	)
	require.NoError(err)

	backend := &txexecutor.Backend{
		Ctx: &snow.Context{
			Log: logging.NoLog{},
		},
		Codec: parser.Codec(),
	}

	baseDB := versiondb.New(memdb.New())

	state, err := state.New(baseDB, parser, registerer, trackChecksums)
	require.NoError(err)

	clk := &mockable.Clock{}
	onAccept := func(*txs.Tx) error { return nil }
	now := time.Now()
	parentTimestamp := now.Add(-2 * time.Second)
	parentID := ids.GenerateTestID()
	cm := parser.Codec()
	txs, err := createParentTxs(cm)
	require.NoError(err)
	parentBlk, err := block.NewStandardBlock(parentID, 0, parentTimestamp, txs, cm)
	require.NoError(err)
	state.AddBlock(parentBlk)
	state.SetLastAccepted(parentBlk.ID())

	metrics, err := metrics.New(registerer)
	require.NoError(err)

	manager := blkexecutor.NewManager(mempool, metrics, state, backend, clk, onAccept)

	manager.SetPreference(parentBlk.ID())

	builder := New(backend, manager, clk, mempool)

	// show that build block fails if tx is invalid
	_, err = builder.BuildBlock(context.Background())
	require.ErrorIs(err, ErrNoTransactions)
}

func createTxs() []*txs.Tx {
	return []*txs.Tx{{
		Unsigned: &txs.BaseTx{BaseTx: avax.BaseTx{
			NetworkID:    constants.UnitTestID,
			BlockchainID: ids.GenerateTestID(),
			Outs: []*avax.TransferableOutput{{
				Asset: avax.Asset{ID: ids.GenerateTestID()},
				Out: &secp256k1fx.TransferOutput{
					OutputOwners: secp256k1fx.OutputOwners{
						Addrs: []ids.ShortID{ids.GenerateTestShortID()},
					},
				},
			}},
			Ins: []*avax.TransferableInput{{
				UTXOID: avax.UTXOID{
					TxID:        ids.ID{'t', 'x', 'I', 'D'},
					OutputIndex: 1,
				},
				Asset: avax.Asset{ID: ids.GenerateTestID()},
				In: &secp256k1fx.TransferInput{
					Amt: uint64(54321),
					Input: secp256k1fx.Input{
						SigIndices: []uint32{2},
					},
				},
			}},
			Memo: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		}},
		Creds: []*fxs.FxCredential{
			{
				Credential: &secp256k1fx.Credential{},
			},
		},
	}}
}

func createParentTxs(cm codec.Manager) ([]*txs.Tx, error) {
	countTxs := 1
	testTxs := make([]*txs.Tx, 0, countTxs)
	for i := 0; i < countTxs; i++ {
		// Create the tx
		tx := &txs.Tx{Unsigned: &txs.BaseTx{BaseTx: avax.BaseTx{
			NetworkID:    constants.UnitTestID,
			BlockchainID: chainID,
			Outs: []*avax.TransferableOutput{{
				Asset: avax.Asset{ID: ids.ID{1, 2, 3}},
				Out: &secp256k1fx.TransferOutput{
					Amt: uint64(12345),
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: 1,
						Addrs:     []ids.ShortID{keys[0].PublicKey().Address()},
					},
				},
			}},
			Ins: []*avax.TransferableInput{{
				UTXOID: avax.UTXOID{
					TxID:        ids.ID{'t', 'x', 'p', 'a', 'r', 'e', 'n', 't'},
					OutputIndex: 1,
				},
				Asset: avax.Asset{ID: ids.ID{1, 2, 3}},
				In: &secp256k1fx.TransferInput{
					Amt: uint64(54321),
					Input: secp256k1fx.Input{
						SigIndices: []uint32{2},
					},
				},
			}},
			Memo: []byte{1, 2, 9, 4, 5, 6, 7, 8},
		}}}
		if err := tx.SignSECP256K1Fx(cm, [][]*secp256k1.PrivateKey{{keys[0]}}); err != nil {
			return nil, err
		}
		testTxs = append(testTxs, tx)
	}
	return testTxs, nil
}
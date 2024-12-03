// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shubhamdubey02/cryftgo/vms/avm/state (interfaces: Chain)
//
// Generated by this command:
//
//	mockgen -package=statemock -destination=vms/avm/state/statemock/chain.go -mock_names=Chain=Chain github.com/shubhamdubey02/cryftgo/vms/avm/state Chain
//

// Package statemock is a generated GoMock package.
package statemock

import (
	reflect "reflect"
	time "time"

	ids "github.com/shubhamdubey02/cryftgo/ids"
	block "github.com/shubhamdubey02/cryftgo/vms/avm/block"
	txs "github.com/shubhamdubey02/cryftgo/vms/avm/txs"
	avax "github.com/shubhamdubey02/cryftgo/vms/components/avax"
	gomock "go.uber.org/mock/gomock"
)

// Chain is a mock of Chain interface.
type Chain struct {
	ctrl     *gomock.Controller
	recorder *ChainMockRecorder
}

// ChainMockRecorder is the mock recorder for Chain.
type ChainMockRecorder struct {
	mock *Chain
}

// NewChain creates a new mock instance.
func NewChain(ctrl *gomock.Controller) *Chain {
	mock := &Chain{ctrl: ctrl}
	mock.recorder = &ChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Chain) EXPECT() *ChainMockRecorder {
	return m.recorder
}

// AddBlock mocks base method.
func (m *Chain) AddBlock(arg0 block.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBlock", arg0)
}

// AddBlock indicates an expected call of AddBlock.
func (mr *ChainMockRecorder) AddBlock(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlock", reflect.TypeOf((*Chain)(nil).AddBlock), arg0)
}

// AddTx mocks base method.
func (m *Chain) AddTx(arg0 *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTx", arg0)
}

// AddTx indicates an expected call of AddTx.
func (mr *ChainMockRecorder) AddTx(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTx", reflect.TypeOf((*Chain)(nil).AddTx), arg0)
}

// AddUTXO mocks base method.
func (m *Chain) AddUTXO(arg0 *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddUTXO", arg0)
}

// AddUTXO indicates an expected call of AddUTXO.
func (mr *ChainMockRecorder) AddUTXO(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUTXO", reflect.TypeOf((*Chain)(nil).AddUTXO), arg0)
}

// DeleteUTXO mocks base method.
func (m *Chain) DeleteUTXO(arg0 ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUTXO", arg0)
}

// DeleteUTXO indicates an expected call of DeleteUTXO.
func (mr *ChainMockRecorder) DeleteUTXO(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUTXO", reflect.TypeOf((*Chain)(nil).DeleteUTXO), arg0)
}

// GetBlock mocks base method.
func (m *Chain) GetBlock(arg0 ids.ID) (block.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0)
	ret0, _ := ret[0].(block.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *ChainMockRecorder) GetBlock(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*Chain)(nil).GetBlock), arg0)
}

// GetBlockIDAtHeight mocks base method.
func (m *Chain) GetBlockIDAtHeight(arg0 uint64) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockIDAtHeight", arg0)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockIDAtHeight indicates an expected call of GetBlockIDAtHeight.
func (mr *ChainMockRecorder) GetBlockIDAtHeight(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockIDAtHeight", reflect.TypeOf((*Chain)(nil).GetBlockIDAtHeight), arg0)
}

// GetLastAccepted mocks base method.
func (m *Chain) GetLastAccepted() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastAccepted")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// GetLastAccepted indicates an expected call of GetLastAccepted.
func (mr *ChainMockRecorder) GetLastAccepted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastAccepted", reflect.TypeOf((*Chain)(nil).GetLastAccepted))
}

// GetTimestamp mocks base method.
func (m *Chain) GetTimestamp() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimestamp")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimestamp indicates an expected call of GetTimestamp.
func (mr *ChainMockRecorder) GetTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimestamp", reflect.TypeOf((*Chain)(nil).GetTimestamp))
}

// GetTx mocks base method.
func (m *Chain) GetTx(arg0 ids.ID) (*txs.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTx", arg0)
	ret0, _ := ret[0].(*txs.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTx indicates an expected call of GetTx.
func (mr *ChainMockRecorder) GetTx(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*Chain)(nil).GetTx), arg0)
}

// GetUTXO mocks base method.
func (m *Chain) GetUTXO(arg0 ids.ID) (*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUTXO", arg0)
	ret0, _ := ret[0].(*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUTXO indicates an expected call of GetUTXO.
func (mr *ChainMockRecorder) GetUTXO(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUTXO", reflect.TypeOf((*Chain)(nil).GetUTXO), arg0)
}

// SetLastAccepted mocks base method.
func (m *Chain) SetLastAccepted(arg0 ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastAccepted", arg0)
}

// SetLastAccepted indicates an expected call of SetLastAccepted.
func (mr *ChainMockRecorder) SetLastAccepted(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastAccepted", reflect.TypeOf((*Chain)(nil).SetLastAccepted), arg0)
}

// SetTimestamp mocks base method.
func (m *Chain) SetTimestamp(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimestamp", arg0)
}

// SetTimestamp indicates an expected call of SetTimestamp.
func (mr *ChainMockRecorder) SetTimestamp(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimestamp", reflect.TypeOf((*Chain)(nil).SetTimestamp), arg0)
}

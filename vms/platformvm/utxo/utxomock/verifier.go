// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shubhamdubey02/cryftgo/vms/platformvm/utxo (interfaces: Verifier)
//
// Generated by this command:
//
//	mockgen -package=utxomock -destination=vms/platformvm/utxo/utxomock/verifier.go -mock_names=Verifier=Verifier github.com/shubhamdubey02/cryftgo/vms/platformvm/utxo Verifier
//

// Package utxomock is a generated GoMock package.
package utxomock

import (
	reflect "reflect"

	ids "github.com/shubhamdubey02/cryftgo/ids"
	avax "github.com/shubhamdubey02/cryftgo/vms/components/avax"
	verify "github.com/shubhamdubey02/cryftgo/vms/components/verify"
	txs "github.com/shubhamdubey02/cryftgo/vms/platformvm/txs"
	gomock "go.uber.org/mock/gomock"
)

// Verifier is a mock of Verifier interface.
type Verifier struct {
	ctrl     *gomock.Controller
	recorder *VerifierMockRecorder
}

// VerifierMockRecorder is the mock recorder for Verifier.
type VerifierMockRecorder struct {
	mock *Verifier
}

// NewVerifier creates a new mock instance.
func NewVerifier(ctrl *gomock.Controller) *Verifier {
	mock := &Verifier{ctrl: ctrl}
	mock.recorder = &VerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Verifier) EXPECT() *VerifierMockRecorder {
	return m.recorder
}

// VerifySpend mocks base method.
func (m *Verifier) VerifySpend(arg0 txs.UnsignedTx, arg1 avax.UTXOGetter, arg2 []*avax.TransferableInput, arg3 []*avax.TransferableOutput, arg4 []verify.Verifiable, arg5 map[ids.ID]uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySpend", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySpend indicates an expected call of VerifySpend.
func (mr *VerifierMockRecorder) VerifySpend(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySpend", reflect.TypeOf((*Verifier)(nil).VerifySpend), arg0, arg1, arg2, arg3, arg4, arg5)
}

// VerifySpendUTXOs mocks base method.
func (m *Verifier) VerifySpendUTXOs(arg0 txs.UnsignedTx, arg1 []*avax.UTXO, arg2 []*avax.TransferableInput, arg3 []*avax.TransferableOutput, arg4 []verify.Verifiable, arg5 map[ids.ID]uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySpendUTXOs", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySpendUTXOs indicates an expected call of VerifySpendUTXOs.
func (mr *VerifierMockRecorder) VerifySpendUTXOs(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySpendUTXOs", reflect.TypeOf((*Verifier)(nil).VerifySpendUTXOs), arg0, arg1, arg2, arg3, arg4, arg5)
}
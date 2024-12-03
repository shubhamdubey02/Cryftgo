// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shubhamdubey02/cryftgo/vms/platformvm/fx (interfaces: Owner)
//
// Generated by this command:
//
//	mockgen -package=fxmock -destination=vms/platformvm/fx/fxmock/owner.go -mock_names=Owner=Owner github.com/shubhamdubey02/cryftgo/vms/platformvm/fx Owner
//

// Package fxmock is a generated GoMock package.
package fxmock

import (
	reflect "reflect"

	snow "github.com/shubhamdubey02/cryftgo/snow"
	verify "github.com/shubhamdubey02/cryftgo/vms/components/verify"
	gomock "go.uber.org/mock/gomock"
)

// Owner is a mock of Owner interface.
type Owner struct {
	verify.IsNotState

	ctrl     *gomock.Controller
	recorder *OwnerMockRecorder
}

// OwnerMockRecorder is the mock recorder for Owner.
type OwnerMockRecorder struct {
	mock *Owner
}

// NewOwner creates a new mock instance.
func NewOwner(ctrl *gomock.Controller) *Owner {
	mock := &Owner{ctrl: ctrl}
	mock.recorder = &OwnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Owner) EXPECT() *OwnerMockRecorder {
	return m.recorder
}

// InitCtx mocks base method.
func (m *Owner) InitCtx(arg0 *snow.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitCtx", arg0)
}

// InitCtx indicates an expected call of InitCtx.
func (mr *OwnerMockRecorder) InitCtx(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitCtx", reflect.TypeOf((*Owner)(nil).InitCtx), arg0)
}

// Verify mocks base method.
func (m *Owner) Verify() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify")
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *OwnerMockRecorder) Verify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*Owner)(nil).Verify))
}

// isState mocks base method.
func (m *Owner) isState() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isState")
	ret0, _ := ret[0].(error)
	return ret0
}

// isState indicates an expected call of isState.
func (mr *OwnerMockRecorder) isState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isState", reflect.TypeOf((*Owner)(nil).isState))
}

// Code generated by MockGen. DO NOT EDIT.
// Source: vms/platformvm/state/diff.go
//
// Generated by this command:
//
//	mockgen -source=vms/platformvm/state/diff.go -destination=vms/platformvm/state/mock_diff.go -package=state -exclude_interfaces= -mock_names=MockDiff=MockDiff
//

// Package state is a generated GoMock package.
package state

import (
	reflect "reflect"
	time "time"

	ids "github.com/shubhamdubey02/cryftgo/ids"
	iterator "github.com/shubhamdubey02/cryftgo/utils/iterator"
	avax "github.com/shubhamdubey02/cryftgo/vms/components/avax"
	gas "github.com/shubhamdubey02/cryftgo/vms/components/gas"
	fx "github.com/shubhamdubey02/cryftgo/vms/platformvm/fx"
	status "github.com/shubhamdubey02/cryftgo/vms/platformvm/status"
	txs "github.com/shubhamdubey02/cryftgo/vms/platformvm/txs"
	gomock "go.uber.org/mock/gomock"
)

// MockDiff is a mock of Diff interface.
type MockDiff struct {
	ctrl     *gomock.Controller
	recorder *MockDiffMockRecorder
}

// MockDiffMockRecorder is the mock recorder for MockDiff.
type MockDiffMockRecorder struct {
	mock *MockDiff
}

// NewMockDiff creates a new mock instance.
func NewMockDiff(ctrl *gomock.Controller) *MockDiff {
	mock := &MockDiff{ctrl: ctrl}
	mock.recorder = &MockDiffMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiff) EXPECT() *MockDiffMockRecorder {
	return m.recorder
}

// AddChain mocks base method.
func (m *MockDiff) AddChain(createChainTx *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChain", createChainTx)
}

// AddChain indicates an expected call of AddChain.
func (mr *MockDiffMockRecorder) AddChain(createChainTx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChain", reflect.TypeOf((*MockDiff)(nil).AddChain), createChainTx)
}

// AddRewardUTXO mocks base method.
func (m *MockDiff) AddRewardUTXO(txID ids.ID, utxo *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRewardUTXO", txID, utxo)
}

// AddRewardUTXO indicates an expected call of AddRewardUTXO.
func (mr *MockDiffMockRecorder) AddRewardUTXO(txID, utxo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRewardUTXO", reflect.TypeOf((*MockDiff)(nil).AddRewardUTXO), txID, utxo)
}

// AddSubnet mocks base method.
func (m *MockDiff) AddSubnet(subnetID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSubnet", subnetID)
}

// AddSubnet indicates an expected call of AddSubnet.
func (mr *MockDiffMockRecorder) AddSubnet(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnet", reflect.TypeOf((*MockDiff)(nil).AddSubnet), subnetID)
}

// AddSubnetTransformation mocks base method.
func (m *MockDiff) AddSubnetTransformation(transformSubnetTx *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSubnetTransformation", transformSubnetTx)
}

// AddSubnetTransformation indicates an expected call of AddSubnetTransformation.
func (mr *MockDiffMockRecorder) AddSubnetTransformation(transformSubnetTx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnetTransformation", reflect.TypeOf((*MockDiff)(nil).AddSubnetTransformation), transformSubnetTx)
}

// AddTx mocks base method.
func (m *MockDiff) AddTx(tx *txs.Tx, status status.Status) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTx", tx, status)
}

// AddTx indicates an expected call of AddTx.
func (mr *MockDiffMockRecorder) AddTx(tx, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTx", reflect.TypeOf((*MockDiff)(nil).AddTx), tx, status)
}

// AddUTXO mocks base method.
func (m *MockDiff) AddUTXO(utxo *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddUTXO", utxo)
}

// AddUTXO indicates an expected call of AddUTXO.
func (mr *MockDiffMockRecorder) AddUTXO(utxo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUTXO", reflect.TypeOf((*MockDiff)(nil).AddUTXO), utxo)
}

// Apply mocks base method.
func (m *MockDiff) Apply(arg0 Chain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockDiffMockRecorder) Apply(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockDiff)(nil).Apply), arg0)
}

// DeleteCurrentDelegator mocks base method.
func (m *MockDiff) DeleteCurrentDelegator(staker *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCurrentDelegator", staker)
}

// DeleteCurrentDelegator indicates an expected call of DeleteCurrentDelegator.
func (mr *MockDiffMockRecorder) DeleteCurrentDelegator(staker any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrentDelegator", reflect.TypeOf((*MockDiff)(nil).DeleteCurrentDelegator), staker)
}

// DeleteCurrentValidator mocks base method.
func (m *MockDiff) DeleteCurrentValidator(staker *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCurrentValidator", staker)
}

// DeleteCurrentValidator indicates an expected call of DeleteCurrentValidator.
func (mr *MockDiffMockRecorder) DeleteCurrentValidator(staker any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrentValidator", reflect.TypeOf((*MockDiff)(nil).DeleteCurrentValidator), staker)
}

// DeleteExpiry mocks base method.
func (m *MockDiff) DeleteExpiry(arg0 ExpiryEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteExpiry", arg0)
}

// DeleteExpiry indicates an expected call of DeleteExpiry.
func (mr *MockDiffMockRecorder) DeleteExpiry(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiry", reflect.TypeOf((*MockDiff)(nil).DeleteExpiry), arg0)
}

// DeletePendingDelegator mocks base method.
func (m *MockDiff) DeletePendingDelegator(staker *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePendingDelegator", staker)
}

// DeletePendingDelegator indicates an expected call of DeletePendingDelegator.
func (mr *MockDiffMockRecorder) DeletePendingDelegator(staker any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePendingDelegator", reflect.TypeOf((*MockDiff)(nil).DeletePendingDelegator), staker)
}

// DeletePendingValidator mocks base method.
func (m *MockDiff) DeletePendingValidator(staker *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePendingValidator", staker)
}

// DeletePendingValidator indicates an expected call of DeletePendingValidator.
func (mr *MockDiffMockRecorder) DeletePendingValidator(staker any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePendingValidator", reflect.TypeOf((*MockDiff)(nil).DeletePendingValidator), staker)
}

// DeleteUTXO mocks base method.
func (m *MockDiff) DeleteUTXO(utxoID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUTXO", utxoID)
}

// DeleteUTXO indicates an expected call of DeleteUTXO.
func (mr *MockDiffMockRecorder) DeleteUTXO(utxoID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUTXO", reflect.TypeOf((*MockDiff)(nil).DeleteUTXO), utxoID)
}

// GetAccruedFees mocks base method.
func (m *MockDiff) GetAccruedFees() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccruedFees")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetAccruedFees indicates an expected call of GetAccruedFees.
func (mr *MockDiffMockRecorder) GetAccruedFees() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccruedFees", reflect.TypeOf((*MockDiff)(nil).GetAccruedFees))
}

// GetCurrentDelegatorIterator mocks base method.
func (m *MockDiff) GetCurrentDelegatorIterator(subnetID ids.ID, nodeID ids.NodeID) (iterator.Iterator[*Staker], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentDelegatorIterator", subnetID, nodeID)
	ret0, _ := ret[0].(iterator.Iterator[*Staker])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentDelegatorIterator indicates an expected call of GetCurrentDelegatorIterator.
func (mr *MockDiffMockRecorder) GetCurrentDelegatorIterator(subnetID, nodeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentDelegatorIterator", reflect.TypeOf((*MockDiff)(nil).GetCurrentDelegatorIterator), subnetID, nodeID)
}

// GetCurrentStakerIterator mocks base method.
func (m *MockDiff) GetCurrentStakerIterator() (iterator.Iterator[*Staker], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentStakerIterator")
	ret0, _ := ret[0].(iterator.Iterator[*Staker])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentStakerIterator indicates an expected call of GetCurrentStakerIterator.
func (mr *MockDiffMockRecorder) GetCurrentStakerIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentStakerIterator", reflect.TypeOf((*MockDiff)(nil).GetCurrentStakerIterator))
}

// GetCurrentSupply mocks base method.
func (m *MockDiff) GetCurrentSupply(subnetID ids.ID) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSupply", subnetID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentSupply indicates an expected call of GetCurrentSupply.
func (mr *MockDiffMockRecorder) GetCurrentSupply(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSupply", reflect.TypeOf((*MockDiff)(nil).GetCurrentSupply), subnetID)
}

// GetCurrentValidator mocks base method.
func (m *MockDiff) GetCurrentValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentValidator", subnetID, nodeID)
	ret0, _ := ret[0].(*Staker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentValidator indicates an expected call of GetCurrentValidator.
func (mr *MockDiffMockRecorder) GetCurrentValidator(subnetID, nodeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentValidator", reflect.TypeOf((*MockDiff)(nil).GetCurrentValidator), subnetID, nodeID)
}

// GetDelegateeReward mocks base method.
func (m *MockDiff) GetDelegateeReward(subnetID ids.ID, nodeID ids.NodeID) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDelegateeReward", subnetID, nodeID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDelegateeReward indicates an expected call of GetDelegateeReward.
func (mr *MockDiffMockRecorder) GetDelegateeReward(subnetID, nodeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDelegateeReward", reflect.TypeOf((*MockDiff)(nil).GetDelegateeReward), subnetID, nodeID)
}

// GetExpiryIterator mocks base method.
func (m *MockDiff) GetExpiryIterator() (iterator.Iterator[ExpiryEntry], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpiryIterator")
	ret0, _ := ret[0].(iterator.Iterator[ExpiryEntry])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpiryIterator indicates an expected call of GetExpiryIterator.
func (mr *MockDiffMockRecorder) GetExpiryIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpiryIterator", reflect.TypeOf((*MockDiff)(nil).GetExpiryIterator))
}

// GetFeeState mocks base method.
func (m *MockDiff) GetFeeState() gas.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeeState")
	ret0, _ := ret[0].(gas.State)
	return ret0
}

// GetFeeState indicates an expected call of GetFeeState.
func (mr *MockDiffMockRecorder) GetFeeState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeeState", reflect.TypeOf((*MockDiff)(nil).GetFeeState))
}

// GetPendingDelegatorIterator mocks base method.
func (m *MockDiff) GetPendingDelegatorIterator(subnetID ids.ID, nodeID ids.NodeID) (iterator.Iterator[*Staker], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingDelegatorIterator", subnetID, nodeID)
	ret0, _ := ret[0].(iterator.Iterator[*Staker])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingDelegatorIterator indicates an expected call of GetPendingDelegatorIterator.
func (mr *MockDiffMockRecorder) GetPendingDelegatorIterator(subnetID, nodeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingDelegatorIterator", reflect.TypeOf((*MockDiff)(nil).GetPendingDelegatorIterator), subnetID, nodeID)
}

// GetPendingStakerIterator mocks base method.
func (m *MockDiff) GetPendingStakerIterator() (iterator.Iterator[*Staker], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingStakerIterator")
	ret0, _ := ret[0].(iterator.Iterator[*Staker])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingStakerIterator indicates an expected call of GetPendingStakerIterator.
func (mr *MockDiffMockRecorder) GetPendingStakerIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingStakerIterator", reflect.TypeOf((*MockDiff)(nil).GetPendingStakerIterator))
}

// GetPendingValidator mocks base method.
func (m *MockDiff) GetPendingValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingValidator", subnetID, nodeID)
	ret0, _ := ret[0].(*Staker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingValidator indicates an expected call of GetPendingValidator.
func (mr *MockDiffMockRecorder) GetPendingValidator(subnetID, nodeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingValidator", reflect.TypeOf((*MockDiff)(nil).GetPendingValidator), subnetID, nodeID)
}

// GetSubnetManager mocks base method.
func (m *MockDiff) GetSubnetManager(subnetID ids.ID) (ids.ID, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetManager", subnetID)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSubnetManager indicates an expected call of GetSubnetManager.
func (mr *MockDiffMockRecorder) GetSubnetManager(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetManager", reflect.TypeOf((*MockDiff)(nil).GetSubnetManager), subnetID)
}

// GetSubnetOwner mocks base method.
func (m *MockDiff) GetSubnetOwner(subnetID ids.ID) (fx.Owner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetOwner", subnetID)
	ret0, _ := ret[0].(fx.Owner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnetOwner indicates an expected call of GetSubnetOwner.
func (mr *MockDiffMockRecorder) GetSubnetOwner(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetOwner", reflect.TypeOf((*MockDiff)(nil).GetSubnetOwner), subnetID)
}

// GetSubnetTransformation mocks base method.
func (m *MockDiff) GetSubnetTransformation(subnetID ids.ID) (*txs.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetTransformation", subnetID)
	ret0, _ := ret[0].(*txs.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnetTransformation indicates an expected call of GetSubnetTransformation.
func (mr *MockDiffMockRecorder) GetSubnetTransformation(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetTransformation", reflect.TypeOf((*MockDiff)(nil).GetSubnetTransformation), subnetID)
}

// GetTimestamp mocks base method.
func (m *MockDiff) GetTimestamp() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimestamp")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimestamp indicates an expected call of GetTimestamp.
func (mr *MockDiffMockRecorder) GetTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimestamp", reflect.TypeOf((*MockDiff)(nil).GetTimestamp))
}

// GetTx mocks base method.
func (m *MockDiff) GetTx(txID ids.ID) (*txs.Tx, status.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTx", txID)
	ret0, _ := ret[0].(*txs.Tx)
	ret1, _ := ret[1].(status.Status)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTx indicates an expected call of GetTx.
func (mr *MockDiffMockRecorder) GetTx(txID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*MockDiff)(nil).GetTx), txID)
}

// GetUTXO mocks base method.
func (m *MockDiff) GetUTXO(utxoID ids.ID) (*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUTXO", utxoID)
	ret0, _ := ret[0].(*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUTXO indicates an expected call of GetUTXO.
func (mr *MockDiffMockRecorder) GetUTXO(utxoID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUTXO", reflect.TypeOf((*MockDiff)(nil).GetUTXO), utxoID)
}

// HasExpiry mocks base method.
func (m *MockDiff) HasExpiry(arg0 ExpiryEntry) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasExpiry", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasExpiry indicates an expected call of HasExpiry.
func (mr *MockDiffMockRecorder) HasExpiry(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasExpiry", reflect.TypeOf((*MockDiff)(nil).HasExpiry), arg0)
}

// PutCurrentDelegator mocks base method.
func (m *MockDiff) PutCurrentDelegator(staker *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutCurrentDelegator", staker)
}

// PutCurrentDelegator indicates an expected call of PutCurrentDelegator.
func (mr *MockDiffMockRecorder) PutCurrentDelegator(staker any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCurrentDelegator", reflect.TypeOf((*MockDiff)(nil).PutCurrentDelegator), staker)
}

// PutCurrentValidator mocks base method.
func (m *MockDiff) PutCurrentValidator(staker *Staker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutCurrentValidator", staker)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutCurrentValidator indicates an expected call of PutCurrentValidator.
func (mr *MockDiffMockRecorder) PutCurrentValidator(staker any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCurrentValidator", reflect.TypeOf((*MockDiff)(nil).PutCurrentValidator), staker)
}

// PutExpiry mocks base method.
func (m *MockDiff) PutExpiry(arg0 ExpiryEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutExpiry", arg0)
}

// PutExpiry indicates an expected call of PutExpiry.
func (mr *MockDiffMockRecorder) PutExpiry(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutExpiry", reflect.TypeOf((*MockDiff)(nil).PutExpiry), arg0)
}

// PutPendingDelegator mocks base method.
func (m *MockDiff) PutPendingDelegator(staker *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPendingDelegator", staker)
}

// PutPendingDelegator indicates an expected call of PutPendingDelegator.
func (mr *MockDiffMockRecorder) PutPendingDelegator(staker any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPendingDelegator", reflect.TypeOf((*MockDiff)(nil).PutPendingDelegator), staker)
}

// PutPendingValidator mocks base method.
func (m *MockDiff) PutPendingValidator(staker *Staker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPendingValidator", staker)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutPendingValidator indicates an expected call of PutPendingValidator.
func (mr *MockDiffMockRecorder) PutPendingValidator(staker any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPendingValidator", reflect.TypeOf((*MockDiff)(nil).PutPendingValidator), staker)
}

// SetAccruedFees mocks base method.
func (m *MockDiff) SetAccruedFees(f uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccruedFees", f)
}

// SetAccruedFees indicates an expected call of SetAccruedFees.
func (mr *MockDiffMockRecorder) SetAccruedFees(f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccruedFees", reflect.TypeOf((*MockDiff)(nil).SetAccruedFees), f)
}

// SetCurrentSupply mocks base method.
func (m *MockDiff) SetCurrentSupply(subnetID ids.ID, cs uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentSupply", subnetID, cs)
}

// SetCurrentSupply indicates an expected call of SetCurrentSupply.
func (mr *MockDiffMockRecorder) SetCurrentSupply(subnetID, cs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentSupply", reflect.TypeOf((*MockDiff)(nil).SetCurrentSupply), subnetID, cs)
}

// SetDelegateeReward mocks base method.
func (m *MockDiff) SetDelegateeReward(subnetID ids.ID, nodeID ids.NodeID, amount uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDelegateeReward", subnetID, nodeID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDelegateeReward indicates an expected call of SetDelegateeReward.
func (mr *MockDiffMockRecorder) SetDelegateeReward(subnetID, nodeID, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDelegateeReward", reflect.TypeOf((*MockDiff)(nil).SetDelegateeReward), subnetID, nodeID, amount)
}

// SetFeeState mocks base method.
func (m *MockDiff) SetFeeState(f gas.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFeeState", f)
}

// SetFeeState indicates an expected call of SetFeeState.
func (mr *MockDiffMockRecorder) SetFeeState(f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFeeState", reflect.TypeOf((*MockDiff)(nil).SetFeeState), f)
}

// SetSubnetManager mocks base method.
func (m *MockDiff) SetSubnetManager(subnetID, chainID ids.ID, addr []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSubnetManager", subnetID, chainID, addr)
}

// SetSubnetManager indicates an expected call of SetSubnetManager.
func (mr *MockDiffMockRecorder) SetSubnetManager(subnetID, chainID, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubnetManager", reflect.TypeOf((*MockDiff)(nil).SetSubnetManager), subnetID, chainID, addr)
}

// SetSubnetOwner mocks base method.
func (m *MockDiff) SetSubnetOwner(subnetID ids.ID, owner fx.Owner) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSubnetOwner", subnetID, owner)
}

// SetSubnetOwner indicates an expected call of SetSubnetOwner.
func (mr *MockDiffMockRecorder) SetSubnetOwner(subnetID, owner any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubnetOwner", reflect.TypeOf((*MockDiff)(nil).SetSubnetOwner), subnetID, owner)
}

// SetTimestamp mocks base method.
func (m *MockDiff) SetTimestamp(tm time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimestamp", tm)
}

// SetTimestamp indicates an expected call of SetTimestamp.
func (mr *MockDiffMockRecorder) SetTimestamp(tm any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimestamp", reflect.TypeOf((*MockDiff)(nil).SetTimestamp), tm)
}

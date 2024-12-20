// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shubhamdubey02/cryftgo/vms/avm/metrics (interfaces: Metrics)
//
// Generated by this command:
//
//	mockgen -package=metricsmock -destination=vms/avm/metrics/metricsmock/metrics.go -mock_names=Metrics=Metrics github.com/shubhamdubey02/cryftgo/vms/avm/metrics Metrics
//

// Package metricsmock is a generated GoMock package.
package metricsmock

import (
	http "net/http"
	reflect "reflect"

	block "github.com/shubhamdubey02/cryftgo/vms/avm/block"
	txs "github.com/shubhamdubey02/cryftgo/vms/avm/txs"
	rpc "github.com/gorilla/rpc/v2"
	gomock "go.uber.org/mock/gomock"
)

// Metrics is a mock of Metrics interface.
type Metrics struct {
	ctrl     *gomock.Controller
	recorder *MetricsMockRecorder
}

// MetricsMockRecorder is the mock recorder for Metrics.
type MetricsMockRecorder struct {
	mock *Metrics
}

// NewMetrics creates a new mock instance.
func NewMetrics(ctrl *gomock.Controller) *Metrics {
	mock := &Metrics{ctrl: ctrl}
	mock.recorder = &MetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Metrics) EXPECT() *MetricsMockRecorder {
	return m.recorder
}

// AfterRequest mocks base method.
func (m *Metrics) AfterRequest(arg0 *rpc.RequestInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterRequest", arg0)
}

// AfterRequest indicates an expected call of AfterRequest.
func (mr *MetricsMockRecorder) AfterRequest(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterRequest", reflect.TypeOf((*Metrics)(nil).AfterRequest), arg0)
}

// IncTxRefreshHits mocks base method.
func (m *Metrics) IncTxRefreshHits() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncTxRefreshHits")
}

// IncTxRefreshHits indicates an expected call of IncTxRefreshHits.
func (mr *MetricsMockRecorder) IncTxRefreshHits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncTxRefreshHits", reflect.TypeOf((*Metrics)(nil).IncTxRefreshHits))
}

// IncTxRefreshMisses mocks base method.
func (m *Metrics) IncTxRefreshMisses() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncTxRefreshMisses")
}

// IncTxRefreshMisses indicates an expected call of IncTxRefreshMisses.
func (mr *MetricsMockRecorder) IncTxRefreshMisses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncTxRefreshMisses", reflect.TypeOf((*Metrics)(nil).IncTxRefreshMisses))
}

// IncTxRefreshes mocks base method.
func (m *Metrics) IncTxRefreshes() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncTxRefreshes")
}

// IncTxRefreshes indicates an expected call of IncTxRefreshes.
func (mr *MetricsMockRecorder) IncTxRefreshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncTxRefreshes", reflect.TypeOf((*Metrics)(nil).IncTxRefreshes))
}

// InterceptRequest mocks base method.
func (m *Metrics) InterceptRequest(arg0 *rpc.RequestInfo) *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterceptRequest", arg0)
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// InterceptRequest indicates an expected call of InterceptRequest.
func (mr *MetricsMockRecorder) InterceptRequest(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterceptRequest", reflect.TypeOf((*Metrics)(nil).InterceptRequest), arg0)
}

// MarkBlockAccepted mocks base method.
func (m *Metrics) MarkBlockAccepted(arg0 block.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkBlockAccepted", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkBlockAccepted indicates an expected call of MarkBlockAccepted.
func (mr *MetricsMockRecorder) MarkBlockAccepted(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkBlockAccepted", reflect.TypeOf((*Metrics)(nil).MarkBlockAccepted), arg0)
}

// MarkTxAccepted mocks base method.
func (m *Metrics) MarkTxAccepted(arg0 *txs.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkTxAccepted", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkTxAccepted indicates an expected call of MarkTxAccepted.
func (mr *MetricsMockRecorder) MarkTxAccepted(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkTxAccepted", reflect.TypeOf((*Metrics)(nil).MarkTxAccepted), arg0)
}

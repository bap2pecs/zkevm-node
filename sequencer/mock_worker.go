// Code generated by mockery v2.22.1. DO NOT EDIT.

package sequencer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	state "github.com/0xPolygonHermez/zkevm-node/state"

	types "github.com/ethereum/go-ethereum/core/types"
)

// WorkerMock is an autogenerated mock type for the workerInterface type
type WorkerMock struct {
	mock.Mock
}

// AddTxTracker provides a mock function with given fields: ctx, txTracker
func (_m *WorkerMock) AddTxTracker(ctx context.Context, txTracker *TxTracker) (error, bool) {
	ret := _m.Called(ctx, txTracker)

	var r0 error
	var r1 bool
	if rf, ok := ret.Get(0).(func(context.Context, *TxTracker) (error, bool)); ok {
		return rf(ctx, txTracker)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *TxTracker) error); ok {
		r0 = rf(ctx, txTracker)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *TxTracker) bool); ok {
		r1 = rf(ctx, txTracker)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// DeleteTx provides a mock function with given fields: txHash, from
func (_m *WorkerMock) DeleteTx(txHash common.Hash, from common.Address) {
	_m.Called(txHash, from)
}

// GetBestFittingTx provides a mock function with given fields: resources
func (_m *WorkerMock) GetBestFittingTx(resources state.BatchResources) *TxTracker {
	ret := _m.Called(resources)

	var r0 *TxTracker
	if rf, ok := ret.Get(0).(func(state.BatchResources) *TxTracker); ok {
		r0 = rf(resources)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TxTracker)
		}
	}

	return r0
}

// HandleL2Reorg provides a mock function with given fields: txHashes
func (_m *WorkerMock) HandleL2Reorg(txHashes []common.Hash) {
	_m.Called(txHashes)
}

// MoveTxToNotReady provides a mock function with given fields: txHash, from, actualNonce, actualBalance
func (_m *WorkerMock) MoveTxToNotReady(txHash common.Hash, from common.Address, actualNonce *uint64, actualBalance *big.Int) []*TxTracker {
	ret := _m.Called(txHash, from, actualNonce, actualBalance)

	var r0 []*TxTracker
	if rf, ok := ret.Get(0).(func(common.Hash, common.Address, *uint64, *big.Int) []*TxTracker); ok {
		r0 = rf(txHash, from, actualNonce, actualBalance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TxTracker)
		}
	}

	return r0
}

// NewTxTracker provides a mock function with given fields: tx, isClaim, counters, ip
func (_m *WorkerMock) NewTxTracker(tx types.Transaction, isClaim bool, counters state.ZKCounters, ip string) (*TxTracker, error) {
	ret := _m.Called(tx, isClaim, counters, ip)

	var r0 *TxTracker
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Transaction, bool, state.ZKCounters, string) (*TxTracker, error)); ok {
		return rf(tx, isClaim, counters, ip)
	}
	if rf, ok := ret.Get(0).(func(types.Transaction, bool, state.ZKCounters, string) *TxTracker); ok {
		r0 = rf(tx, isClaim, counters, ip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TxTracker)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Transaction, bool, state.ZKCounters, string) error); ok {
		r1 = rf(tx, isClaim, counters, ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAfterSingleSuccessfulTxExecution provides a mock function with given fields: from, touchedAddresses
func (_m *WorkerMock) UpdateAfterSingleSuccessfulTxExecution(from common.Address, touchedAddresses map[common.Address]*state.InfoReadWrite) []*TxTracker {
	ret := _m.Called(from, touchedAddresses)

	var r0 []*TxTracker
	if rf, ok := ret.Get(0).(func(common.Address, map[common.Address]*state.InfoReadWrite) []*TxTracker); ok {
		r0 = rf(from, touchedAddresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*TxTracker)
		}
	}

	return r0
}

// UpdateTx provides a mock function with given fields: txHash, from, ZKCounters
func (_m *WorkerMock) UpdateTx(txHash common.Hash, from common.Address, ZKCounters state.ZKCounters) {
	_m.Called(txHash, from, ZKCounters)
}

type mockConstructorTestingTNewWorkerMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewWorkerMock creates a new instance of WorkerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWorkerMock(t mockConstructorTestingTNewWorkerMock) *WorkerMock {
	mock := &WorkerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

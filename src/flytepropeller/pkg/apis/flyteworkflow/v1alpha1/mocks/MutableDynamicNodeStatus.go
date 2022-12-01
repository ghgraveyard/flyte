// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flyte/src/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// MutableDynamicNodeStatus is an autogenerated mock type for the MutableDynamicNodeStatus type
type MutableDynamicNodeStatus struct {
	mock.Mock
}

type MutableDynamicNodeStatus_GetDynamicNodePhase struct {
	*mock.Call
}

func (_m MutableDynamicNodeStatus_GetDynamicNodePhase) Return(_a0 v1alpha1.DynamicNodePhase) *MutableDynamicNodeStatus_GetDynamicNodePhase {
	return &MutableDynamicNodeStatus_GetDynamicNodePhase{Call: _m.Call.Return(_a0)}
}

func (_m *MutableDynamicNodeStatus) OnGetDynamicNodePhase() *MutableDynamicNodeStatus_GetDynamicNodePhase {
	c_call := _m.On("GetDynamicNodePhase")
	return &MutableDynamicNodeStatus_GetDynamicNodePhase{Call: c_call}
}

func (_m *MutableDynamicNodeStatus) OnGetDynamicNodePhaseMatch(matchers ...interface{}) *MutableDynamicNodeStatus_GetDynamicNodePhase {
	c_call := _m.On("GetDynamicNodePhase", matchers...)
	return &MutableDynamicNodeStatus_GetDynamicNodePhase{Call: c_call}
}

// GetDynamicNodePhase provides a mock function with given fields:
func (_m *MutableDynamicNodeStatus) GetDynamicNodePhase() v1alpha1.DynamicNodePhase {
	ret := _m.Called()

	var r0 v1alpha1.DynamicNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.DynamicNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.DynamicNodePhase)
	}

	return r0
}

type MutableDynamicNodeStatus_GetDynamicNodeReason struct {
	*mock.Call
}

func (_m MutableDynamicNodeStatus_GetDynamicNodeReason) Return(_a0 string) *MutableDynamicNodeStatus_GetDynamicNodeReason {
	return &MutableDynamicNodeStatus_GetDynamicNodeReason{Call: _m.Call.Return(_a0)}
}

func (_m *MutableDynamicNodeStatus) OnGetDynamicNodeReason() *MutableDynamicNodeStatus_GetDynamicNodeReason {
	c_call := _m.On("GetDynamicNodeReason")
	return &MutableDynamicNodeStatus_GetDynamicNodeReason{Call: c_call}
}

func (_m *MutableDynamicNodeStatus) OnGetDynamicNodeReasonMatch(matchers ...interface{}) *MutableDynamicNodeStatus_GetDynamicNodeReason {
	c_call := _m.On("GetDynamicNodeReason", matchers...)
	return &MutableDynamicNodeStatus_GetDynamicNodeReason{Call: c_call}
}

// GetDynamicNodeReason provides a mock function with given fields:
func (_m *MutableDynamicNodeStatus) GetDynamicNodeReason() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type MutableDynamicNodeStatus_GetExecutionError struct {
	*mock.Call
}

func (_m MutableDynamicNodeStatus_GetExecutionError) Return(_a0 *core.ExecutionError) *MutableDynamicNodeStatus_GetExecutionError {
	return &MutableDynamicNodeStatus_GetExecutionError{Call: _m.Call.Return(_a0)}
}

func (_m *MutableDynamicNodeStatus) OnGetExecutionError() *MutableDynamicNodeStatus_GetExecutionError {
	c_call := _m.On("GetExecutionError")
	return &MutableDynamicNodeStatus_GetExecutionError{Call: c_call}
}

func (_m *MutableDynamicNodeStatus) OnGetExecutionErrorMatch(matchers ...interface{}) *MutableDynamicNodeStatus_GetExecutionError {
	c_call := _m.On("GetExecutionError", matchers...)
	return &MutableDynamicNodeStatus_GetExecutionError{Call: c_call}
}

// GetExecutionError provides a mock function with given fields:
func (_m *MutableDynamicNodeStatus) GetExecutionError() *core.ExecutionError {
	ret := _m.Called()

	var r0 *core.ExecutionError
	if rf, ok := ret.Get(0).(func() *core.ExecutionError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ExecutionError)
		}
	}

	return r0
}

type MutableDynamicNodeStatus_IsDirty struct {
	*mock.Call
}

func (_m MutableDynamicNodeStatus_IsDirty) Return(_a0 bool) *MutableDynamicNodeStatus_IsDirty {
	return &MutableDynamicNodeStatus_IsDirty{Call: _m.Call.Return(_a0)}
}

func (_m *MutableDynamicNodeStatus) OnIsDirty() *MutableDynamicNodeStatus_IsDirty {
	c_call := _m.On("IsDirty")
	return &MutableDynamicNodeStatus_IsDirty{Call: c_call}
}

func (_m *MutableDynamicNodeStatus) OnIsDirtyMatch(matchers ...interface{}) *MutableDynamicNodeStatus_IsDirty {
	c_call := _m.On("IsDirty", matchers...)
	return &MutableDynamicNodeStatus_IsDirty{Call: c_call}
}

// IsDirty provides a mock function with given fields:
func (_m *MutableDynamicNodeStatus) IsDirty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetDynamicNodePhase provides a mock function with given fields: phase
func (_m *MutableDynamicNodeStatus) SetDynamicNodePhase(phase v1alpha1.DynamicNodePhase) {
	_m.Called(phase)
}

// SetDynamicNodeReason provides a mock function with given fields: reason
func (_m *MutableDynamicNodeStatus) SetDynamicNodeReason(reason string) {
	_m.Called(reason)
}

// SetExecutionError provides a mock function with given fields: executionError
func (_m *MutableDynamicNodeStatus) SetExecutionError(executionError *core.ExecutionError) {
	_m.Called(executionError)
}

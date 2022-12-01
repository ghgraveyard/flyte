// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/flyteorg/flyte/src/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MutableGateNodeStatus is an autogenerated mock type for the MutableGateNodeStatus type
type MutableGateNodeStatus struct {
	mock.Mock
}

type MutableGateNodeStatus_GetGateNodePhase struct {
	*mock.Call
}

func (_m MutableGateNodeStatus_GetGateNodePhase) Return(_a0 v1alpha1.GateNodePhase) *MutableGateNodeStatus_GetGateNodePhase {
	return &MutableGateNodeStatus_GetGateNodePhase{Call: _m.Call.Return(_a0)}
}

func (_m *MutableGateNodeStatus) OnGetGateNodePhase() *MutableGateNodeStatus_GetGateNodePhase {
	c_call := _m.On("GetGateNodePhase")
	return &MutableGateNodeStatus_GetGateNodePhase{Call: c_call}
}

func (_m *MutableGateNodeStatus) OnGetGateNodePhaseMatch(matchers ...interface{}) *MutableGateNodeStatus_GetGateNodePhase {
	c_call := _m.On("GetGateNodePhase", matchers...)
	return &MutableGateNodeStatus_GetGateNodePhase{Call: c_call}
}

// GetGateNodePhase provides a mock function with given fields:
func (_m *MutableGateNodeStatus) GetGateNodePhase() v1alpha1.GateNodePhase {
	ret := _m.Called()

	var r0 v1alpha1.GateNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.GateNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.GateNodePhase)
	}

	return r0
}

type MutableGateNodeStatus_IsDirty struct {
	*mock.Call
}

func (_m MutableGateNodeStatus_IsDirty) Return(_a0 bool) *MutableGateNodeStatus_IsDirty {
	return &MutableGateNodeStatus_IsDirty{Call: _m.Call.Return(_a0)}
}

func (_m *MutableGateNodeStatus) OnIsDirty() *MutableGateNodeStatus_IsDirty {
	c_call := _m.On("IsDirty")
	return &MutableGateNodeStatus_IsDirty{Call: c_call}
}

func (_m *MutableGateNodeStatus) OnIsDirtyMatch(matchers ...interface{}) *MutableGateNodeStatus_IsDirty {
	c_call := _m.On("IsDirty", matchers...)
	return &MutableGateNodeStatus_IsDirty{Call: c_call}
}

// IsDirty provides a mock function with given fields:
func (_m *MutableGateNodeStatus) IsDirty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetGateNodePhase provides a mock function with given fields: phase
func (_m *MutableGateNodeStatus) SetGateNodePhase(phase v1alpha1.GateNodePhase) {
	_m.Called(phase)
}

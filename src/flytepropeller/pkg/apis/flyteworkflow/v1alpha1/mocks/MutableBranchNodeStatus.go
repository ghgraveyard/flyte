// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/flyteorg/flyte/src/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MutableBranchNodeStatus is an autogenerated mock type for the MutableBranchNodeStatus type
type MutableBranchNodeStatus struct {
	mock.Mock
}

type MutableBranchNodeStatus_GetFinalizedNode struct {
	*mock.Call
}

func (_m MutableBranchNodeStatus_GetFinalizedNode) Return(_a0 *string) *MutableBranchNodeStatus_GetFinalizedNode {
	return &MutableBranchNodeStatus_GetFinalizedNode{Call: _m.Call.Return(_a0)}
}

func (_m *MutableBranchNodeStatus) OnGetFinalizedNode() *MutableBranchNodeStatus_GetFinalizedNode {
	c_call := _m.On("GetFinalizedNode")
	return &MutableBranchNodeStatus_GetFinalizedNode{Call: c_call}
}

func (_m *MutableBranchNodeStatus) OnGetFinalizedNodeMatch(matchers ...interface{}) *MutableBranchNodeStatus_GetFinalizedNode {
	c_call := _m.On("GetFinalizedNode", matchers...)
	return &MutableBranchNodeStatus_GetFinalizedNode{Call: c_call}
}

// GetFinalizedNode provides a mock function with given fields:
func (_m *MutableBranchNodeStatus) GetFinalizedNode() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

type MutableBranchNodeStatus_GetPhase struct {
	*mock.Call
}

func (_m MutableBranchNodeStatus_GetPhase) Return(_a0 v1alpha1.BranchNodePhase) *MutableBranchNodeStatus_GetPhase {
	return &MutableBranchNodeStatus_GetPhase{Call: _m.Call.Return(_a0)}
}

func (_m *MutableBranchNodeStatus) OnGetPhase() *MutableBranchNodeStatus_GetPhase {
	c_call := _m.On("GetPhase")
	return &MutableBranchNodeStatus_GetPhase{Call: c_call}
}

func (_m *MutableBranchNodeStatus) OnGetPhaseMatch(matchers ...interface{}) *MutableBranchNodeStatus_GetPhase {
	c_call := _m.On("GetPhase", matchers...)
	return &MutableBranchNodeStatus_GetPhase{Call: c_call}
}

// GetPhase provides a mock function with given fields:
func (_m *MutableBranchNodeStatus) GetPhase() v1alpha1.BranchNodePhase {
	ret := _m.Called()

	var r0 v1alpha1.BranchNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.BranchNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.BranchNodePhase)
	}

	return r0
}

type MutableBranchNodeStatus_IsDirty struct {
	*mock.Call
}

func (_m MutableBranchNodeStatus_IsDirty) Return(_a0 bool) *MutableBranchNodeStatus_IsDirty {
	return &MutableBranchNodeStatus_IsDirty{Call: _m.Call.Return(_a0)}
}

func (_m *MutableBranchNodeStatus) OnIsDirty() *MutableBranchNodeStatus_IsDirty {
	c_call := _m.On("IsDirty")
	return &MutableBranchNodeStatus_IsDirty{Call: c_call}
}

func (_m *MutableBranchNodeStatus) OnIsDirtyMatch(matchers ...interface{}) *MutableBranchNodeStatus_IsDirty {
	c_call := _m.On("IsDirty", matchers...)
	return &MutableBranchNodeStatus_IsDirty{Call: c_call}
}

// IsDirty provides a mock function with given fields:
func (_m *MutableBranchNodeStatus) IsDirty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetBranchNodeError provides a mock function with given fields:
func (_m *MutableBranchNodeStatus) SetBranchNodeError() {
	_m.Called()
}

// SetBranchNodeSuccess provides a mock function with given fields: id
func (_m *MutableBranchNodeStatus) SetBranchNodeSuccess(id string) {
	_m.Called(id)
}

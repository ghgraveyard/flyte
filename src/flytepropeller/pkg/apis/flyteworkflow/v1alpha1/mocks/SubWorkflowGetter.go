// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/flyteorg/flyte/src/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// SubWorkflowGetter is an autogenerated mock type for the SubWorkflowGetter type
type SubWorkflowGetter struct {
	mock.Mock
}

type SubWorkflowGetter_FindSubWorkflow struct {
	*mock.Call
}

func (_m SubWorkflowGetter_FindSubWorkflow) Return(_a0 v1alpha1.ExecutableSubWorkflow) *SubWorkflowGetter_FindSubWorkflow {
	return &SubWorkflowGetter_FindSubWorkflow{Call: _m.Call.Return(_a0)}
}

func (_m *SubWorkflowGetter) OnFindSubWorkflow(subID string) *SubWorkflowGetter_FindSubWorkflow {
	c_call := _m.On("FindSubWorkflow", subID)
	return &SubWorkflowGetter_FindSubWorkflow{Call: c_call}
}

func (_m *SubWorkflowGetter) OnFindSubWorkflowMatch(matchers ...interface{}) *SubWorkflowGetter_FindSubWorkflow {
	c_call := _m.On("FindSubWorkflow", matchers...)
	return &SubWorkflowGetter_FindSubWorkflow{Call: c_call}
}

// FindSubWorkflow provides a mock function with given fields: subID
func (_m *SubWorkflowGetter) FindSubWorkflow(subID string) v1alpha1.ExecutableSubWorkflow {
	ret := _m.Called(subID)

	var r0 v1alpha1.ExecutableSubWorkflow
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableSubWorkflow); ok {
		r0 = rf(subID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableSubWorkflow)
		}
	}

	return r0
}

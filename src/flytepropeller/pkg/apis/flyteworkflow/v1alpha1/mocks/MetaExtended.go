// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/flyteorg/flyte/src/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// MetaExtended is an autogenerated mock type for the MetaExtended type
type MetaExtended struct {
	mock.Mock
}

type MetaExtended_FindSubWorkflow struct {
	*mock.Call
}

func (_m MetaExtended_FindSubWorkflow) Return(_a0 v1alpha1.ExecutableSubWorkflow) *MetaExtended_FindSubWorkflow {
	return &MetaExtended_FindSubWorkflow{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnFindSubWorkflow(subID string) *MetaExtended_FindSubWorkflow {
	c_call := _m.On("FindSubWorkflow", subID)
	return &MetaExtended_FindSubWorkflow{Call: c_call}
}

func (_m *MetaExtended) OnFindSubWorkflowMatch(matchers ...interface{}) *MetaExtended_FindSubWorkflow {
	c_call := _m.On("FindSubWorkflow", matchers...)
	return &MetaExtended_FindSubWorkflow{Call: c_call}
}

// FindSubWorkflow provides a mock function with given fields: subID
func (_m *MetaExtended) FindSubWorkflow(subID string) v1alpha1.ExecutableSubWorkflow {
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

type MetaExtended_GetAnnotations struct {
	*mock.Call
}

func (_m MetaExtended_GetAnnotations) Return(_a0 map[string]string) *MetaExtended_GetAnnotations {
	return &MetaExtended_GetAnnotations{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetAnnotations() *MetaExtended_GetAnnotations {
	c_call := _m.On("GetAnnotations")
	return &MetaExtended_GetAnnotations{Call: c_call}
}

func (_m *MetaExtended) OnGetAnnotationsMatch(matchers ...interface{}) *MetaExtended_GetAnnotations {
	c_call := _m.On("GetAnnotations", matchers...)
	return &MetaExtended_GetAnnotations{Call: c_call}
}

// GetAnnotations provides a mock function with given fields:
func (_m *MetaExtended) GetAnnotations() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type MetaExtended_GetCreationTimestamp struct {
	*mock.Call
}

func (_m MetaExtended_GetCreationTimestamp) Return(_a0 v1.Time) *MetaExtended_GetCreationTimestamp {
	return &MetaExtended_GetCreationTimestamp{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetCreationTimestamp() *MetaExtended_GetCreationTimestamp {
	c_call := _m.On("GetCreationTimestamp")
	return &MetaExtended_GetCreationTimestamp{Call: c_call}
}

func (_m *MetaExtended) OnGetCreationTimestampMatch(matchers ...interface{}) *MetaExtended_GetCreationTimestamp {
	c_call := _m.On("GetCreationTimestamp", matchers...)
	return &MetaExtended_GetCreationTimestamp{Call: c_call}
}

// GetCreationTimestamp provides a mock function with given fields:
func (_m *MetaExtended) GetCreationTimestamp() v1.Time {
	ret := _m.Called()

	var r0 v1.Time
	if rf, ok := ret.Get(0).(func() v1.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.Time)
	}

	return r0
}

type MetaExtended_GetDefinitionVersion struct {
	*mock.Call
}

func (_m MetaExtended_GetDefinitionVersion) Return(_a0 v1alpha1.WorkflowDefinitionVersion) *MetaExtended_GetDefinitionVersion {
	return &MetaExtended_GetDefinitionVersion{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetDefinitionVersion() *MetaExtended_GetDefinitionVersion {
	c_call := _m.On("GetDefinitionVersion")
	return &MetaExtended_GetDefinitionVersion{Call: c_call}
}

func (_m *MetaExtended) OnGetDefinitionVersionMatch(matchers ...interface{}) *MetaExtended_GetDefinitionVersion {
	c_call := _m.On("GetDefinitionVersion", matchers...)
	return &MetaExtended_GetDefinitionVersion{Call: c_call}
}

// GetDefinitionVersion provides a mock function with given fields:
func (_m *MetaExtended) GetDefinitionVersion() v1alpha1.WorkflowDefinitionVersion {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowDefinitionVersion
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowDefinitionVersion); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowDefinitionVersion)
	}

	return r0
}

type MetaExtended_GetEventVersion struct {
	*mock.Call
}

func (_m MetaExtended_GetEventVersion) Return(_a0 v1alpha1.EventVersion) *MetaExtended_GetEventVersion {
	return &MetaExtended_GetEventVersion{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetEventVersion() *MetaExtended_GetEventVersion {
	c_call := _m.On("GetEventVersion")
	return &MetaExtended_GetEventVersion{Call: c_call}
}

func (_m *MetaExtended) OnGetEventVersionMatch(matchers ...interface{}) *MetaExtended_GetEventVersion {
	c_call := _m.On("GetEventVersion", matchers...)
	return &MetaExtended_GetEventVersion{Call: c_call}
}

// GetEventVersion provides a mock function with given fields:
func (_m *MetaExtended) GetEventVersion() v1alpha1.EventVersion {
	ret := _m.Called()

	var r0 v1alpha1.EventVersion
	if rf, ok := ret.Get(0).(func() v1alpha1.EventVersion); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.EventVersion)
	}

	return r0
}

type MetaExtended_GetExecutionID struct {
	*mock.Call
}

func (_m MetaExtended_GetExecutionID) Return(_a0 v1alpha1.WorkflowExecutionIdentifier) *MetaExtended_GetExecutionID {
	return &MetaExtended_GetExecutionID{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetExecutionID() *MetaExtended_GetExecutionID {
	c_call := _m.On("GetExecutionID")
	return &MetaExtended_GetExecutionID{Call: c_call}
}

func (_m *MetaExtended) OnGetExecutionIDMatch(matchers ...interface{}) *MetaExtended_GetExecutionID {
	c_call := _m.On("GetExecutionID", matchers...)
	return &MetaExtended_GetExecutionID{Call: c_call}
}

// GetExecutionID provides a mock function with given fields:
func (_m *MetaExtended) GetExecutionID() v1alpha1.WorkflowExecutionIdentifier {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowExecutionIdentifier
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowExecutionIdentifier); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowExecutionIdentifier)
	}

	return r0
}

type MetaExtended_GetExecutionStatus struct {
	*mock.Call
}

func (_m MetaExtended_GetExecutionStatus) Return(_a0 v1alpha1.ExecutableWorkflowStatus) *MetaExtended_GetExecutionStatus {
	return &MetaExtended_GetExecutionStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetExecutionStatus() *MetaExtended_GetExecutionStatus {
	c_call := _m.On("GetExecutionStatus")
	return &MetaExtended_GetExecutionStatus{Call: c_call}
}

func (_m *MetaExtended) OnGetExecutionStatusMatch(matchers ...interface{}) *MetaExtended_GetExecutionStatus {
	c_call := _m.On("GetExecutionStatus", matchers...)
	return &MetaExtended_GetExecutionStatus{Call: c_call}
}

// GetExecutionStatus provides a mock function with given fields:
func (_m *MetaExtended) GetExecutionStatus() v1alpha1.ExecutableWorkflowStatus {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableWorkflowStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableWorkflowStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableWorkflowStatus)
		}
	}

	return r0
}

type MetaExtended_GetK8sWorkflowID struct {
	*mock.Call
}

func (_m MetaExtended_GetK8sWorkflowID) Return(_a0 types.NamespacedName) *MetaExtended_GetK8sWorkflowID {
	return &MetaExtended_GetK8sWorkflowID{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetK8sWorkflowID() *MetaExtended_GetK8sWorkflowID {
	c_call := _m.On("GetK8sWorkflowID")
	return &MetaExtended_GetK8sWorkflowID{Call: c_call}
}

func (_m *MetaExtended) OnGetK8sWorkflowIDMatch(matchers ...interface{}) *MetaExtended_GetK8sWorkflowID {
	c_call := _m.On("GetK8sWorkflowID", matchers...)
	return &MetaExtended_GetK8sWorkflowID{Call: c_call}
}

// GetK8sWorkflowID provides a mock function with given fields:
func (_m *MetaExtended) GetK8sWorkflowID() types.NamespacedName {
	ret := _m.Called()

	var r0 types.NamespacedName
	if rf, ok := ret.Get(0).(func() types.NamespacedName); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.NamespacedName)
	}

	return r0
}

type MetaExtended_GetLabels struct {
	*mock.Call
}

func (_m MetaExtended_GetLabels) Return(_a0 map[string]string) *MetaExtended_GetLabels {
	return &MetaExtended_GetLabels{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetLabels() *MetaExtended_GetLabels {
	c_call := _m.On("GetLabels")
	return &MetaExtended_GetLabels{Call: c_call}
}

func (_m *MetaExtended) OnGetLabelsMatch(matchers ...interface{}) *MetaExtended_GetLabels {
	c_call := _m.On("GetLabels", matchers...)
	return &MetaExtended_GetLabels{Call: c_call}
}

// GetLabels provides a mock function with given fields:
func (_m *MetaExtended) GetLabels() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type MetaExtended_GetName struct {
	*mock.Call
}

func (_m MetaExtended_GetName) Return(_a0 string) *MetaExtended_GetName {
	return &MetaExtended_GetName{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetName() *MetaExtended_GetName {
	c_call := _m.On("GetName")
	return &MetaExtended_GetName{Call: c_call}
}

func (_m *MetaExtended) OnGetNameMatch(matchers ...interface{}) *MetaExtended_GetName {
	c_call := _m.On("GetName", matchers...)
	return &MetaExtended_GetName{Call: c_call}
}

// GetName provides a mock function with given fields:
func (_m *MetaExtended) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type MetaExtended_GetNamespace struct {
	*mock.Call
}

func (_m MetaExtended_GetNamespace) Return(_a0 string) *MetaExtended_GetNamespace {
	return &MetaExtended_GetNamespace{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetNamespace() *MetaExtended_GetNamespace {
	c_call := _m.On("GetNamespace")
	return &MetaExtended_GetNamespace{Call: c_call}
}

func (_m *MetaExtended) OnGetNamespaceMatch(matchers ...interface{}) *MetaExtended_GetNamespace {
	c_call := _m.On("GetNamespace", matchers...)
	return &MetaExtended_GetNamespace{Call: c_call}
}

// GetNamespace provides a mock function with given fields:
func (_m *MetaExtended) GetNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type MetaExtended_GetOwnerReference struct {
	*mock.Call
}

func (_m MetaExtended_GetOwnerReference) Return(_a0 v1.OwnerReference) *MetaExtended_GetOwnerReference {
	return &MetaExtended_GetOwnerReference{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetOwnerReference() *MetaExtended_GetOwnerReference {
	c_call := _m.On("GetOwnerReference")
	return &MetaExtended_GetOwnerReference{Call: c_call}
}

func (_m *MetaExtended) OnGetOwnerReferenceMatch(matchers ...interface{}) *MetaExtended_GetOwnerReference {
	c_call := _m.On("GetOwnerReference", matchers...)
	return &MetaExtended_GetOwnerReference{Call: c_call}
}

// GetOwnerReference provides a mock function with given fields:
func (_m *MetaExtended) GetOwnerReference() v1.OwnerReference {
	ret := _m.Called()

	var r0 v1.OwnerReference
	if rf, ok := ret.Get(0).(func() v1.OwnerReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.OwnerReference)
	}

	return r0
}

type MetaExtended_GetRawOutputDataConfig struct {
	*mock.Call
}

func (_m MetaExtended_GetRawOutputDataConfig) Return(_a0 v1alpha1.RawOutputDataConfig) *MetaExtended_GetRawOutputDataConfig {
	return &MetaExtended_GetRawOutputDataConfig{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetRawOutputDataConfig() *MetaExtended_GetRawOutputDataConfig {
	c_call := _m.On("GetRawOutputDataConfig")
	return &MetaExtended_GetRawOutputDataConfig{Call: c_call}
}

func (_m *MetaExtended) OnGetRawOutputDataConfigMatch(matchers ...interface{}) *MetaExtended_GetRawOutputDataConfig {
	c_call := _m.On("GetRawOutputDataConfig", matchers...)
	return &MetaExtended_GetRawOutputDataConfig{Call: c_call}
}

// GetRawOutputDataConfig provides a mock function with given fields:
func (_m *MetaExtended) GetRawOutputDataConfig() v1alpha1.RawOutputDataConfig {
	ret := _m.Called()

	var r0 v1alpha1.RawOutputDataConfig
	if rf, ok := ret.Get(0).(func() v1alpha1.RawOutputDataConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.RawOutputDataConfig)
	}

	return r0
}

type MetaExtended_GetSecurityContext struct {
	*mock.Call
}

func (_m MetaExtended_GetSecurityContext) Return(_a0 core.SecurityContext) *MetaExtended_GetSecurityContext {
	return &MetaExtended_GetSecurityContext{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetSecurityContext() *MetaExtended_GetSecurityContext {
	c_call := _m.On("GetSecurityContext")
	return &MetaExtended_GetSecurityContext{Call: c_call}
}

func (_m *MetaExtended) OnGetSecurityContextMatch(matchers ...interface{}) *MetaExtended_GetSecurityContext {
	c_call := _m.On("GetSecurityContext", matchers...)
	return &MetaExtended_GetSecurityContext{Call: c_call}
}

// GetSecurityContext provides a mock function with given fields:
func (_m *MetaExtended) GetSecurityContext() core.SecurityContext {
	ret := _m.Called()

	var r0 core.SecurityContext
	if rf, ok := ret.Get(0).(func() core.SecurityContext); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(core.SecurityContext)
	}

	return r0
}

type MetaExtended_GetServiceAccountName struct {
	*mock.Call
}

func (_m MetaExtended_GetServiceAccountName) Return(_a0 string) *MetaExtended_GetServiceAccountName {
	return &MetaExtended_GetServiceAccountName{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnGetServiceAccountName() *MetaExtended_GetServiceAccountName {
	c_call := _m.On("GetServiceAccountName")
	return &MetaExtended_GetServiceAccountName{Call: c_call}
}

func (_m *MetaExtended) OnGetServiceAccountNameMatch(matchers ...interface{}) *MetaExtended_GetServiceAccountName {
	c_call := _m.On("GetServiceAccountName", matchers...)
	return &MetaExtended_GetServiceAccountName{Call: c_call}
}

// GetServiceAccountName provides a mock function with given fields:
func (_m *MetaExtended) GetServiceAccountName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type MetaExtended_GetTask struct {
	*mock.Call
}

func (_m MetaExtended_GetTask) Return(_a0 v1alpha1.ExecutableTask, _a1 error) *MetaExtended_GetTask {
	return &MetaExtended_GetTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *MetaExtended) OnGetTask(id string) *MetaExtended_GetTask {
	c_call := _m.On("GetTask", id)
	return &MetaExtended_GetTask{Call: c_call}
}

func (_m *MetaExtended) OnGetTaskMatch(matchers ...interface{}) *MetaExtended_GetTask {
	c_call := _m.On("GetTask", matchers...)
	return &MetaExtended_GetTask{Call: c_call}
}

// GetTask provides a mock function with given fields: id
func (_m *MetaExtended) GetTask(id string) (v1alpha1.ExecutableTask, error) {
	ret := _m.Called(id)

	var r0 v1alpha1.ExecutableTask
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableTask); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableTask)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type MetaExtended_IsInterruptible struct {
	*mock.Call
}

func (_m MetaExtended_IsInterruptible) Return(_a0 bool) *MetaExtended_IsInterruptible {
	return &MetaExtended_IsInterruptible{Call: _m.Call.Return(_a0)}
}

func (_m *MetaExtended) OnIsInterruptible() *MetaExtended_IsInterruptible {
	c_call := _m.On("IsInterruptible")
	return &MetaExtended_IsInterruptible{Call: c_call}
}

func (_m *MetaExtended) OnIsInterruptibleMatch(matchers ...interface{}) *MetaExtended_IsInterruptible {
	c_call := _m.On("IsInterruptible", matchers...)
	return &MetaExtended_IsInterruptible{Call: c_call}
}

// IsInterruptible provides a mock function with given fields:
func (_m *MetaExtended) IsInterruptible() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

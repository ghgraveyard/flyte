// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	datacatalog "github.com/flyteorg/flyte/gen/pb-go/flyteidl/datacatalog"

	mock "github.com/stretchr/testify/mock"
)

// TagManager is an autogenerated mock type for the TagManager type
type TagManager struct {
	mock.Mock
}

type TagManager_AddTag struct {
	*mock.Call
}

func (_m TagManager_AddTag) Return(_a0 *datacatalog.AddTagResponse, _a1 error) *TagManager_AddTag {
	return &TagManager_AddTag{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *TagManager) OnAddTag(ctx context.Context, request *datacatalog.AddTagRequest) *TagManager_AddTag {
	c_call := _m.On("AddTag", ctx, request)
	return &TagManager_AddTag{Call: c_call}
}

func (_m *TagManager) OnAddTagMatch(matchers ...interface{}) *TagManager_AddTag {
	c_call := _m.On("AddTag", matchers...)
	return &TagManager_AddTag{Call: c_call}
}

// AddTag provides a mock function with given fields: ctx, request
func (_m *TagManager) AddTag(ctx context.Context, request *datacatalog.AddTagRequest) (*datacatalog.AddTagResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *datacatalog.AddTagResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.AddTagRequest) *datacatalog.AddTagResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.AddTagResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.AddTagRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

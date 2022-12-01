// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/flyteorg/flyte/src/datacatalog/pkg/repositories/models"
)

// TagRepo is an autogenerated mock type for the TagRepo type
type TagRepo struct {
	mock.Mock
}

type TagRepo_Create struct {
	*mock.Call
}

func (_m TagRepo_Create) Return(_a0 error) *TagRepo_Create {
	return &TagRepo_Create{Call: _m.Call.Return(_a0)}
}

func (_m *TagRepo) OnCreate(ctx context.Context, in models.Tag) *TagRepo_Create {
	c_call := _m.On("Create", ctx, in)
	return &TagRepo_Create{Call: c_call}
}

func (_m *TagRepo) OnCreateMatch(matchers ...interface{}) *TagRepo_Create {
	c_call := _m.On("Create", matchers...)
	return &TagRepo_Create{Call: c_call}
}

// Create provides a mock function with given fields: ctx, in
func (_m *TagRepo) Create(ctx context.Context, in models.Tag) error {
	ret := _m.Called(ctx, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Tag) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type TagRepo_Get struct {
	*mock.Call
}

func (_m TagRepo_Get) Return(_a0 models.Tag, _a1 error) *TagRepo_Get {
	return &TagRepo_Get{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *TagRepo) OnGet(ctx context.Context, in models.TagKey) *TagRepo_Get {
	c_call := _m.On("Get", ctx, in)
	return &TagRepo_Get{Call: c_call}
}

func (_m *TagRepo) OnGetMatch(matchers ...interface{}) *TagRepo_Get {
	c_call := _m.On("Get", matchers...)
	return &TagRepo_Get{Call: c_call}
}

// Get provides a mock function with given fields: ctx, in
func (_m *TagRepo) Get(ctx context.Context, in models.TagKey) (models.Tag, error) {
	ret := _m.Called(ctx, in)

	var r0 models.Tag
	if rf, ok := ret.Get(0).(func(context.Context, models.TagKey) models.Tag); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(models.Tag)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.TagKey) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

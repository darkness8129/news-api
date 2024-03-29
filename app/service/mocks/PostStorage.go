// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "darkness8129/news-api/app/entity"

	mock "github.com/stretchr/testify/mock"
)

// PostStorage is an autogenerated mock type for the PostStorage type
type PostStorage struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, post
func (_m *PostStorage) Create(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	ret := _m.Called(ctx, post)

	var r0 *entity.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Post) (*entity.Post, error)); ok {
		return rf(ctx, post)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Post) *entity.Post); ok {
		r0 = rf(ctx, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Post) error); ok {
		r1 = rf(ctx, post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *PostStorage) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *PostStorage) Get(ctx context.Context, id string) (*entity.Post, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Post, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Post); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *PostStorage) List(ctx context.Context) ([]entity.Post, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Post, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Post); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, post
func (_m *PostStorage) Update(ctx context.Context, id string, post *entity.Post) (*entity.Post, error) {
	ret := _m.Called(ctx, id, post)

	var r0 *entity.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *entity.Post) (*entity.Post, error)); ok {
		return rf(ctx, id, post)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *entity.Post) *entity.Post); ok {
		r0 = rf(ctx, id, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *entity.Post) error); ok {
		r1 = rf(ctx, id, post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostStorage creates a new instance of PostStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostStorage(t mockConstructorTestingTNewPostStorage) *PostStorage {
	mock := &PostStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

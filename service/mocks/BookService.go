// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	domain "books_service/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	web "books_service/web"
)

// BookService is an autogenerated mock type for the BookService type
type BookService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *BookService) Create(ctx context.Context, request *web.Request) (*domain.Domain, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) (*domain.Domain, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) *domain.Domain); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *web.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *BookService) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx
func (_m *BookService) FindAll(ctx context.Context) ([]*domain.Domain, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*domain.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.Domain, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: ctx, id
func (_m *BookService) FindById(ctx context.Context, id int) (*domain.Domain, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *domain.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*domain.Domain, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *domain.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, request
func (_m *BookService) Update(ctx context.Context, request *web.Request) (*domain.Domain, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *domain.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) (*domain.Domain, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *web.Request) *domain.Domain); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *web.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBookService creates a new instance of BookService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookService(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookService {
	mock := &BookService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

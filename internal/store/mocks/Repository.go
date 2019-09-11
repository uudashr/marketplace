// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import store "github.com/uudashr/marketplace/internal/store"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Store provides a mock function with given fields: _a0
func (_m *Repository) Store(_a0 *store.Store) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*store.Store) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreByID provides a mock function with given fields: id
func (_m *Repository) StoreByID(id string) (*store.Store, error) {
	ret := _m.Called(id)

	var r0 *store.Store
	if rf, ok := ret.Get(0).(func(string) *store.Store); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.Store)
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

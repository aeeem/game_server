// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "game_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// PlayerRepository is an autogenerated mock type for the PlayerRepository type
type PlayerRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx
func (_m *PlayerRepository) Fetch(ctx context.Context) (domain.Player, error) {
	ret := _m.Called(ctx)

	var r0 domain.Player
	if rf, ok := ret.Get(0).(func(context.Context) domain.Player); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(domain.Player)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx
func (_m *PlayerRepository) Store(ctx context.Context) (domain.Player, error) {
	ret := _m.Called(ctx)

	var r0 domain.Player
	if rf, ok := ret.Get(0).(func(context.Context) domain.Player); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(domain.Player)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

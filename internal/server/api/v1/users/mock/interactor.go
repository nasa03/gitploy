// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/hanjunlee/gitploy/ent"
	vo "github.com/hanjunlee/gitploy/vo"
)

// MockInteractor is a mock of Interactor interface.
type MockInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockInteractorMockRecorder
}

// MockInteractorMockRecorder is the mock recorder for MockInteractor.
type MockInteractorMockRecorder struct {
	mock *MockInteractor
}

// NewMockInteractor creates a new mock instance.
func NewMockInteractor(ctrl *gomock.Controller) *MockInteractor {
	mock := &MockInteractor{ctrl: ctrl}
	mock.recorder = &MockInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractor) EXPECT() *MockInteractorMockRecorder {
	return m.recorder
}

// DeleteUser mocks base method.
func (m *MockInteractor) DeleteUser(ctx context.Context, u *ent.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockInteractorMockRecorder) DeleteUser(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockInteractor)(nil).DeleteUser), ctx, u)
}

// FindUserByID mocks base method.
func (m *MockInteractor) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", ctx, id)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockInteractorMockRecorder) FindUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockInteractor)(nil).FindUserByID), ctx, id)
}

// GetRateLimit mocks base method.
func (m *MockInteractor) GetRateLimit(ctx context.Context, u *ent.User) (*vo.RateLimit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRateLimit", ctx, u)
	ret0, _ := ret[0].(*vo.RateLimit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateLimit indicates an expected call of GetRateLimit.
func (mr *MockInteractorMockRecorder) GetRateLimit(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateLimit", reflect.TypeOf((*MockInteractor)(nil).GetRateLimit), ctx, u)
}

// ListUsers mocks base method.
func (m *MockInteractor) ListUsers(ctx context.Context, login string, page, perPage int) ([]*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, login, page, perPage)
	ret0, _ := ret[0].([]*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockInteractorMockRecorder) ListUsers(ctx, login, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockInteractor)(nil).ListUsers), ctx, login, page, perPage)
}

// UpdateUser mocks base method.
func (m *MockInteractor) UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, u)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockInteractorMockRecorder) UpdateUser(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockInteractor)(nil).UpdateUser), ctx, u)
}
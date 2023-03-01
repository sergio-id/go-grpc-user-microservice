// Code generated by MockGen. DO NOT EDIT.
// Source: redis_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/sergio-id/go-grpc-user-microservice/internal/user/domain"
)

// MockRedisRepository is a mock of RedisRepository interface.
type MockRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisRepositoryMockRecorder
}

// MockRedisRepositoryMockRecorder is the mock recorder for MockRedisRepository.
type MockRedisRepositoryMockRecorder struct {
	mock *MockRedisRepository
}

// NewMockRedisRepository creates a new mock instance.
func NewMockRedisRepository(ctrl *gomock.Controller) *MockRedisRepository {
	mock := &MockRedisRepository{ctrl: ctrl}
	mock.recorder = &MockRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisRepository) EXPECT() *MockRedisRepositoryMockRecorder {
	return m.recorder
}

// DeleteUserCtx mocks base method.
func (m *MockRedisRepository) DeleteUserCtx(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserCtx", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserCtx indicates an expected call of DeleteUserCtx.
func (mr *MockRedisRepositoryMockRecorder) DeleteUserCtx(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserCtx", reflect.TypeOf((*MockRedisRepository)(nil).DeleteUserCtx), ctx, id)
}

// GetByIDCtx mocks base method.
func (m *MockRedisRepository) GetByIDCtx(ctx context.Context, id uint64) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDCtx", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDCtx indicates an expected call of GetByIDCtx.
func (mr *MockRedisRepositoryMockRecorder) GetByIDCtx(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDCtx", reflect.TypeOf((*MockRedisRepository)(nil).GetByIDCtx), ctx, id)
}

// SetUserCtx mocks base method.
func (m *MockRedisRepository) SetUserCtx(ctx context.Context, duration time.Duration, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserCtx", ctx, duration, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserCtx indicates an expected call of SetUserCtx.
func (mr *MockRedisRepositoryMockRecorder) SetUserCtx(ctx, duration, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserCtx", reflect.TypeOf((*MockRedisRepository)(nil).SetUserCtx), ctx, duration, user)
}

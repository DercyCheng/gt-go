// Code generated by MockGen. DO NOT EDIT.
// Source: ./oauth2.go

// Package wechatmocks is a generated GoMock package.
package wechatmocks

import (
	context "context"
	"gitee.com/geekbang/basic-go/webook/oauth2/domain/wechat"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AuthURL mocks base method.
func (m *MockService) AuthURL(ctx context.Context, state string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthURL", ctx, state)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthURL indicates an expected call of AuthURL.
func (mr *MockServiceMockRecorder) AuthURL(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthURL", reflect.TypeOf((*MockService)(nil).AuthURL), ctx, state)
}

// VerifyCode mocks base method.
func (m *MockService) VerifyCode(ctx context.Context, code string) (wechat.WechatInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyCode", ctx, code)
	ret0, _ := ret[0].(wechat.WechatInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyCode indicates an expected call of VerifyCode.
func (mr *MockServiceMockRecorder) VerifyCode(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyCode", reflect.TypeOf((*MockService)(nil).VerifyCode), ctx, code)
}

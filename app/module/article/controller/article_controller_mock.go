// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bangadam/go-fiber-starter/app/module/article/controller (interfaces: IArticleController)

// Package controller is a generated GoMock package.
package controller

import (
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockIArticleController is a mock of IArticleController interface.
type MockIArticleController struct {
	ctrl     *gomock.Controller
	recorder *MockIArticleControllerMockRecorder
}

// MockIArticleControllerMockRecorder is the mock recorder for MockIArticleController.
type MockIArticleControllerMockRecorder struct {
	mock *MockIArticleController
}

// NewMockIArticleController creates a new mock instance.
func NewMockIArticleController(ctrl *gomock.Controller) *MockIArticleController {
	mock := &MockIArticleController{ctrl: ctrl}
	mock.recorder = &MockIArticleControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArticleController) EXPECT() *MockIArticleControllerMockRecorder {
	return m.recorder
}

// Index mocks base method.
func (m *MockIArticleController) Index(arg0 *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Index indicates an expected call of Index.
func (mr *MockIArticleControllerMockRecorder) Index(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockIArticleController)(nil).Index), arg0)
}

// Show mocks base method.
func (m *MockIArticleController) Show(arg0 *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Show indicates an expected call of Show.
func (mr *MockIArticleControllerMockRecorder) Show(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockIArticleController)(nil).Show), arg0)
}

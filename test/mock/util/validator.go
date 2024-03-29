// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go

// Package mock_util is a generated GoMock package.
package mock_util

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// ValidateStruct mocks base method.
func (m *MockValidator) ValidateStruct(i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateStruct", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateStruct indicates an expected call of ValidateStruct.
func (mr *MockValidatorMockRecorder) ValidateStruct(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateStruct", reflect.TypeOf((*MockValidator)(nil).ValidateStruct), i)
}

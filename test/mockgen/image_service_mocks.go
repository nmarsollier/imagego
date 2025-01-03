// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/image/service.go

// Package mockgen is a generated GoMock package.
package mockgen

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	image "github.com/nmarsollier/imagego/internal/image"
)

// MockImageService is a mock of ImageService interface.
type MockImageService struct {
	ctrl     *gomock.Controller
	recorder *MockImageServiceMockRecorder
}

// MockImageServiceMockRecorder is the mock recorder for MockImageService.
type MockImageServiceMockRecorder struct {
	mock *MockImageService
}

// NewMockImageService creates a new mock instance.
func NewMockImageService(ctrl *gomock.Controller) *MockImageService {
	mock := &MockImageService{ctrl: ctrl}
	mock.recorder = &MockImageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageService) EXPECT() *MockImageServiceMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockImageService) Find(imageID string, size int) (*image.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", imageID, size)
	ret0, _ := ret[0].(*image.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockImageServiceMockRecorder) Find(imageID, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockImageService)(nil).Find), imageID, size)
}

// Insert mocks base method.
func (m *MockImageService) Insert(image *image.Image) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", image)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockImageServiceMockRecorder) Insert(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockImageService)(nil).Insert), image)
}

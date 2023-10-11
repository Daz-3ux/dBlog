// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Daz-3ux/dBlog/internal/dazBlog/biz/post (interfaces: PostBiz)

// Package post is a generated GoMock package.
package post

import (
	context "context"
	reflect "reflect"

	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockPostBiz is a mock of PostBiz interface.
type MockPostBiz struct {
	ctrl     *gomock.Controller
	recorder *MockPostBizMockRecorder
}

// MockPostBizMockRecorder is the mock recorder for MockPostBiz.
type MockPostBizMockRecorder struct {
	mock *MockPostBiz
}

// NewMockPostBiz creates a new mock instance.
func NewMockPostBiz(ctrl *gomock.Controller) *MockPostBiz {
	mock := &MockPostBiz{ctrl: ctrl}
	mock.recorder = &MockPostBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostBiz) EXPECT() *MockPostBizMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPostBiz) Create(arg0 context.Context, arg1 string, arg2 *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CreatePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostBizMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostBiz)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockPostBiz) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostBizMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostBiz)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockPostBiz) DeleteCollection(arg0 context.Context, arg1 string, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockPostBizMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockPostBiz)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockPostBiz) Get(arg0 context.Context, arg1, arg2 string) (*v1.GetPostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.GetPostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPostBizMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostBiz)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockPostBiz) List(arg0 context.Context, arg1 string, arg2, arg3 int) (*v1.ListPostsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.ListPostsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPostBizMockRecorder) List(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPostBiz)(nil).List), arg0, arg1, arg2, arg3)
}

// Update mocks base method.
func (m *MockPostBiz) Update(arg0 context.Context, arg1, arg2 string, arg3 *v1.UpdatePostRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPostBizMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostBiz)(nil).Update), arg0, arg1, arg2, arg3)
}

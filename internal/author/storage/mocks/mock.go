// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/xlzpm/internal/author/model"
	storage "github.com/xlzpm/internal/author/storage"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, author *model.Author) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, author)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, author)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, id)
}

// FindAll mocks base method.
func (m *MockRepository) FindAll(ctx context.Context, sortOptions storage.SortOptions) ([]model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, sortOptions)
	ret0, _ := ret[0].([]model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(ctx, sortOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), ctx, sortOptions)
}

// FindOne mocks base method.
func (m *MockRepository) FindOne(ctx context.Context, id string) (model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, id)
	ret0, _ := ret[0].(model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockRepositoryMockRecorder) FindOne(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockRepository)(nil).FindOne), ctx, id)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, author model.Author) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, author)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, author)
}

// MockSortOptions is a mock of SortOptions interface.
type MockSortOptions struct {
	ctrl     *gomock.Controller
	recorder *MockSortOptionsMockRecorder
}

// MockSortOptionsMockRecorder is the mock recorder for MockSortOptions.
type MockSortOptionsMockRecorder struct {
	mock *MockSortOptions
}

// NewMockSortOptions creates a new mock instance.
func NewMockSortOptions(ctrl *gomock.Controller) *MockSortOptions {
	mock := &MockSortOptions{ctrl: ctrl}
	mock.recorder = &MockSortOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSortOptions) EXPECT() *MockSortOptionsMockRecorder {
	return m.recorder
}

// GetOrderBy mocks base method.
func (m *MockSortOptions) GetOrderBy() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderBy")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOrderBy indicates an expected call of GetOrderBy.
func (mr *MockSortOptionsMockRecorder) GetOrderBy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderBy", reflect.TypeOf((*MockSortOptions)(nil).GetOrderBy))
}

// MockFilterOptions is a mock of FilterOptions interface.
type MockFilterOptions struct {
	ctrl     *gomock.Controller
	recorder *MockFilterOptionsMockRecorder
}

// MockFilterOptionsMockRecorder is the mock recorder for MockFilterOptions.
type MockFilterOptionsMockRecorder struct {
	mock *MockFilterOptions
}

// NewMockFilterOptions creates a new mock instance.
func NewMockFilterOptions(ctrl *gomock.Controller) *MockFilterOptions {
	mock := &MockFilterOptions{ctrl: ctrl}
	mock.recorder = &MockFilterOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilterOptions) EXPECT() *MockFilterOptionsMockRecorder {
	return m.recorder
}
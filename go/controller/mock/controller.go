// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go
//
// Generated by this command:
//
//	mockgen -package mock -destination mock/controller.go --source=controller.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	pgtype "github.com/jackc/pgx/v5/pgtype"
	notifications "github.com/nhost/hasura-auth/go/notifications"
	sql "github.com/nhost/hasura-auth/go/sql"
	gomock "go.uber.org/mock/gomock"
)

// MockEmailer is a mock of Emailer interface.
type MockEmailer struct {
	ctrl     *gomock.Controller
	recorder *MockEmailerMockRecorder
}

// MockEmailerMockRecorder is the mock recorder for MockEmailer.
type MockEmailerMockRecorder struct {
	mock *MockEmailer
}

// NewMockEmailer creates a new mock instance.
func NewMockEmailer(ctrl *gomock.Controller) *MockEmailer {
	mock := &MockEmailer{ctrl: ctrl}
	mock.recorder = &MockEmailerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailer) EXPECT() *MockEmailerMockRecorder {
	return m.recorder
}

// SendEmailVerify mocks base method.
func (m *MockEmailer) SendEmailVerify(to, locale string, data notifications.EmailVerifyData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmailVerify", to, locale, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEmailVerify indicates an expected call of SendEmailVerify.
func (mr *MockEmailerMockRecorder) SendEmailVerify(to, locale, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmailVerify", reflect.TypeOf((*MockEmailer)(nil).SendEmailVerify), to, locale, data)
}

// MockDBClient is a mock of DBClient interface.
type MockDBClient struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientMockRecorder
}

// MockDBClientMockRecorder is the mock recorder for MockDBClient.
type MockDBClientMockRecorder struct {
	mock *MockDBClient
}

// NewMockDBClient creates a new mock instance.
func NewMockDBClient(ctrl *gomock.Controller) *MockDBClient {
	mock := &MockDBClient{ctrl: ctrl}
	mock.recorder = &MockDBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBClient) EXPECT() *MockDBClientMockRecorder {
	return m.recorder
}

// GetUserByEmail mocks base method.
func (m *MockDBClient) GetUserByEmail(ctx context.Context, email pgtype.Text) (sql.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(sql.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockDBClientMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockDBClient)(nil).GetUserByEmail), ctx, email)
}

// InsertUser mocks base method.
func (m *MockDBClient) InsertUser(ctx context.Context, arg sql.InsertUserParams) (sql.InsertUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", ctx, arg)
	ret0, _ := ret[0].(sql.InsertUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockDBClientMockRecorder) InsertUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockDBClient)(nil).InsertUser), ctx, arg)
}

// InsertUserWithRefreshToken mocks base method.
func (m *MockDBClient) InsertUserWithRefreshToken(ctx context.Context, arg sql.InsertUserWithRefreshTokenParams) (sql.InsertUserWithRefreshTokenRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserWithRefreshToken", ctx, arg)
	ret0, _ := ret[0].(sql.InsertUserWithRefreshTokenRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserWithRefreshToken indicates an expected call of InsertUserWithRefreshToken.
func (mr *MockDBClientMockRecorder) InsertUserWithRefreshToken(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserWithRefreshToken", reflect.TypeOf((*MockDBClient)(nil).InsertUserWithRefreshToken), ctx, arg)
}
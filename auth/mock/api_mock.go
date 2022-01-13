// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	auth "github.com/polarismesh/polaris-server/auth"
	cache "github.com/polarismesh/polaris-server/cache"
	v1 "github.com/polarismesh/polaris-server/common/api/v1"
	model "github.com/polarismesh/polaris-server/common/model"
	reflect "reflect"
)

// MockAuthManager is a mock of AuthManager interface
type MockAuthManager struct {
	ctrl     *gomock.Controller
	recorder *MockAuthManagerMockRecorder
}

// MockAuthManagerMockRecorder is the mock recorder for MockAuthManager
type MockAuthManagerMockRecorder struct {
	mock *MockAuthManager
}

// NewMockAuthManager creates a new mock instance
func NewMockAuthManager(ctrl *gomock.Controller) *MockAuthManager {
	mock := &MockAuthManager{ctrl: ctrl}
	mock.recorder = &MockAuthManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthManager) EXPECT() *MockAuthManagerMockRecorder {
	return m.recorder
}

// Initialize mocks base method
func (m *MockAuthManager) Initialize(options *auth.Config, cacheMgn *cache.NamingCache) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", options, cacheMgn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockAuthManagerMockRecorder) Initialize(options, cacheMgn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockAuthManager)(nil).Initialize), options, cacheMgn)
}

// Login mocks base method
func (m *MockAuthManager) Login(req *v1.LoginRequest) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", req)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockAuthManagerMockRecorder) Login(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthManager)(nil).Login), req)
}

// HasPermission mocks base method
func (m *MockAuthManager) HasPermission(preCtx *model.AcquireContext) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermission", preCtx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasPermission indicates an expected call of HasPermission
func (mr *MockAuthManagerMockRecorder) HasPermission(preCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermission", reflect.TypeOf((*MockAuthManager)(nil).HasPermission), preCtx)
}

// ChangeOpenStatus mocks base method
func (m *MockAuthManager) ChangeOpenStatus(status auth.AuthStatus) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOpenStatus", status)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ChangeOpenStatus indicates an expected call of ChangeOpenStatus
func (mr *MockAuthManagerMockRecorder) ChangeOpenStatus(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOpenStatus", reflect.TypeOf((*MockAuthManager)(nil).ChangeOpenStatus), status)
}

// IsOpenAuth mocks base method
func (m *MockAuthManager) IsOpenAuth() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOpenAuth")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOpenAuth indicates an expected call of IsOpenAuth
func (mr *MockAuthManagerMockRecorder) IsOpenAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOpenAuth", reflect.TypeOf((*MockAuthManager)(nil).IsOpenAuth))
}

// Name mocks base method
func (m *MockAuthManager) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockAuthManagerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAuthManager)(nil).Name))
}

// GetUserServer mocks base method
func (m *MockAuthManager) GetUserServer() auth.UserServer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserServer")
	ret0, _ := ret[0].(auth.UserServer)
	return ret0
}

// GetUserServer indicates an expected call of GetUserServer
func (mr *MockAuthManagerMockRecorder) GetUserServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserServer", reflect.TypeOf((*MockAuthManager)(nil).GetUserServer))
}

// GetAuthStrategyServer mocks base method
func (m *MockAuthManager) GetAuthStrategyServer() auth.AuthStrategyServer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthStrategyServer")
	ret0, _ := ret[0].(auth.AuthStrategyServer)
	return ret0
}

// GetAuthStrategyServer indicates an expected call of GetAuthStrategyServer
func (mr *MockAuthManagerMockRecorder) GetAuthStrategyServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthStrategyServer", reflect.TypeOf((*MockAuthManager)(nil).GetAuthStrategyServer))
}

// AfterResourceOperation mocks base method
func (m *MockAuthManager) AfterResourceOperation(afterCtx *model.AcquireContext) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterResourceOperation", afterCtx)
}

// AfterResourceOperation indicates an expected call of AfterResourceOperation
func (mr *MockAuthManagerMockRecorder) AfterResourceOperation(afterCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterResourceOperation", reflect.TypeOf((*MockAuthManager)(nil).AfterResourceOperation), afterCtx)
}

// MockUserServer is a mock of UserServer interface
type MockUserServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServerMockRecorder
}

// MockUserServerMockRecorder is the mock recorder for MockUserServer
type MockUserServerMockRecorder struct {
	mock *MockUserServer
}

// NewMockUserServer creates a new mock instance
func NewMockUserServer(ctrl *gomock.Controller) *MockUserServer {
	mock := &MockUserServer{ctrl: ctrl}
	mock.recorder = &MockUserServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserServer) EXPECT() *MockUserServerMockRecorder {
	return m.recorder
}

// CreateUsers mocks base method
func (m *MockUserServer) CreateUsers(ctx context.Context, users []*v1.User) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUsers", ctx, users)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// CreateUsers indicates an expected call of CreateUsers
func (mr *MockUserServerMockRecorder) CreateUsers(ctx, users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUsers", reflect.TypeOf((*MockUserServer)(nil).CreateUsers), ctx, users)
}

// UpdateUser mocks base method
func (m *MockUserServer) UpdateUser(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserServerMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserServer)(nil).UpdateUser), ctx, user)
}

// DeleteUser mocks base method
func (m *MockUserServer) DeleteUser(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockUserServerMockRecorder) DeleteUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserServer)(nil).DeleteUser), ctx, user)
}

// ListUsers mocks base method
func (m *MockUserServer) ListUsers(ctx context.Context, query map[string]string) *v1.BatchQueryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, query)
	ret0, _ := ret[0].(*v1.BatchQueryResponse)
	return ret0
}

// ListUsers indicates an expected call of ListUsers
func (mr *MockUserServerMockRecorder) ListUsers(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockUserServer)(nil).ListUsers), ctx, query)
}

// GetUserToken mocks base method
func (m *MockUserServer) GetUserToken(ctx context.Context, filter map[string]string) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserToken", ctx, filter)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetUserToken indicates an expected call of GetUserToken
func (mr *MockUserServerMockRecorder) GetUserToken(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserToken", reflect.TypeOf((*MockUserServer)(nil).GetUserToken), ctx, filter)
}

// ChangeUserTokenStatus mocks base method
func (m *MockUserServer) ChangeUserTokenStatus(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserTokenStatus", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// ChangeUserTokenStatus indicates an expected call of ChangeUserTokenStatus
func (mr *MockUserServerMockRecorder) ChangeUserTokenStatus(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserTokenStatus", reflect.TypeOf((*MockUserServer)(nil).ChangeUserTokenStatus), ctx, user)
}

// RefreshUserToken mocks base method
func (m *MockUserServer) RefreshUserToken(ctx context.Context, user *v1.User) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshUserToken", ctx, user)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// RefreshUserToken indicates an expected call of RefreshUserToken
func (mr *MockUserServerMockRecorder) RefreshUserToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshUserToken", reflect.TypeOf((*MockUserServer)(nil).RefreshUserToken), ctx, user)
}

// CreateUserGroup mocks base method
func (m *MockUserServer) CreateUserGroup(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserGroup", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// CreateUserGroup indicates an expected call of CreateUserGroup
func (mr *MockUserServerMockRecorder) CreateUserGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserGroup", reflect.TypeOf((*MockUserServer)(nil).CreateUserGroup), ctx, group)
}

// UpdateUserGroup mocks base method
func (m *MockUserServer) UpdateUserGroup(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserGroup", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateUserGroup indicates an expected call of UpdateUserGroup
func (mr *MockUserServerMockRecorder) UpdateUserGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserGroup", reflect.TypeOf((*MockUserServer)(nil).UpdateUserGroup), ctx, group)
}

// DeleteUserGroup mocks base method
func (m *MockUserServer) DeleteUserGroup(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserGroup", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// DeleteUserGroup indicates an expected call of DeleteUserGroup
func (mr *MockUserServerMockRecorder) DeleteUserGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserGroup", reflect.TypeOf((*MockUserServer)(nil).DeleteUserGroup), ctx, group)
}

// ListUserGroups mocks base method
func (m *MockUserServer) ListUserGroups(ctx context.Context, query map[string]string) *v1.BatchQueryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserGroups", ctx, query)
	ret0, _ := ret[0].(*v1.BatchQueryResponse)
	return ret0
}

// ListUserGroups indicates an expected call of ListUserGroups
func (mr *MockUserServerMockRecorder) ListUserGroups(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserGroups", reflect.TypeOf((*MockUserServer)(nil).ListUserGroups), ctx, query)
}

// GetUserGroupToken mocks base method
func (m *MockUserServer) GetUserGroupToken(ctx context.Context, filter map[string]string) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserGroupToken", ctx, filter)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// GetUserGroupToken indicates an expected call of GetUserGroupToken
func (mr *MockUserServerMockRecorder) GetUserGroupToken(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserGroupToken", reflect.TypeOf((*MockUserServer)(nil).GetUserGroupToken), ctx, filter)
}

// ChangeUserGroupTokenStatus mocks base method
func (m *MockUserServer) ChangeUserGroupTokenStatus(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserGroupTokenStatus", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// ChangeUserGroupTokenStatus indicates an expected call of ChangeUserGroupTokenStatus
func (mr *MockUserServerMockRecorder) ChangeUserGroupTokenStatus(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserGroupTokenStatus", reflect.TypeOf((*MockUserServer)(nil).ChangeUserGroupTokenStatus), ctx, group)
}

// RefreshUserGroupToken mocks base method
func (m *MockUserServer) RefreshUserGroupToken(ctx context.Context, group *v1.UserGroup) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshUserGroupToken", ctx, group)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// RefreshUserGroupToken indicates an expected call of RefreshUserGroupToken
func (mr *MockUserServerMockRecorder) RefreshUserGroupToken(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshUserGroupToken", reflect.TypeOf((*MockUserServer)(nil).RefreshUserGroupToken), ctx, group)
}

// BatchAddUserToGroup mocks base method
func (m *MockUserServer) BatchAddUserToGroup(ctx context.Context, relation *v1.UserGroupRelation) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchAddUserToGroup", ctx, relation)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// BatchAddUserToGroup indicates an expected call of BatchAddUserToGroup
func (mr *MockUserServerMockRecorder) BatchAddUserToGroup(ctx, relation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchAddUserToGroup", reflect.TypeOf((*MockUserServer)(nil).BatchAddUserToGroup), ctx, relation)
}

// BatchRemoveUserFromGroup mocks base method
func (m *MockUserServer) BatchRemoveUserFromGroup(ctx context.Context, relation *v1.UserGroupRelation) *v1.BatchWriteResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchRemoveUserFromGroup", ctx, relation)
	ret0, _ := ret[0].(*v1.BatchWriteResponse)
	return ret0
}

// BatchRemoveUserFromGroup indicates an expected call of BatchRemoveUserFromGroup
func (mr *MockUserServerMockRecorder) BatchRemoveUserFromGroup(ctx, relation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchRemoveUserFromGroup", reflect.TypeOf((*MockUserServer)(nil).BatchRemoveUserFromGroup), ctx, relation)
}

// MockAuthStrategyServer is a mock of AuthStrategyServer interface
type MockAuthStrategyServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthStrategyServerMockRecorder
}

// MockAuthStrategyServerMockRecorder is the mock recorder for MockAuthStrategyServer
type MockAuthStrategyServerMockRecorder struct {
	mock *MockAuthStrategyServer
}

// NewMockAuthStrategyServer creates a new mock instance
func NewMockAuthStrategyServer(ctrl *gomock.Controller) *MockAuthStrategyServer {
	mock := &MockAuthStrategyServer{ctrl: ctrl}
	mock.recorder = &MockAuthStrategyServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthStrategyServer) EXPECT() *MockAuthStrategyServerMockRecorder {
	return m.recorder
}

// CreateStrategy mocks base method
func (m *MockAuthStrategyServer) CreateStrategy(ctx context.Context, strategy *v1.AuthStrategy) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStrategy", ctx, strategy)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// CreateStrategy indicates an expected call of CreateStrategy
func (mr *MockAuthStrategyServerMockRecorder) CreateStrategy(ctx, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStrategy", reflect.TypeOf((*MockAuthStrategyServer)(nil).CreateStrategy), ctx, strategy)
}

// UpdateStrategy mocks base method
func (m *MockAuthStrategyServer) UpdateStrategy(ctx context.Context, strategy *v1.ModifyAuthStrategy) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStrategy", ctx, strategy)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// UpdateStrategy indicates an expected call of UpdateStrategy
func (mr *MockAuthStrategyServerMockRecorder) UpdateStrategy(ctx, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStrategy", reflect.TypeOf((*MockAuthStrategyServer)(nil).UpdateStrategy), ctx, strategy)
}

// DeleteStrategy mocks base method
func (m *MockAuthStrategyServer) DeleteStrategy(ctx context.Context, strategy *v1.AuthStrategy) *v1.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStrategy", ctx, strategy)
	ret0, _ := ret[0].(*v1.Response)
	return ret0
}

// DeleteStrategy indicates an expected call of DeleteStrategy
func (mr *MockAuthStrategyServerMockRecorder) DeleteStrategy(ctx, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStrategy", reflect.TypeOf((*MockAuthStrategyServer)(nil).DeleteStrategy), ctx, strategy)
}

// ListStrategy mocks base method
func (m *MockAuthStrategyServer) ListStrategy(ctx context.Context, query map[string]string) *v1.BatchQueryResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStrategy", ctx, query)
	ret0, _ := ret[0].(*v1.BatchQueryResponse)
	return ret0
}

// ListStrategy indicates an expected call of ListStrategy
func (mr *MockAuthStrategyServerMockRecorder) ListStrategy(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStrategy", reflect.TypeOf((*MockAuthStrategyServer)(nil).ListStrategy), ctx, query)
}

// MockAuthority is a mock of Authority interface
type MockAuthority struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorityMockRecorder
}

// MockAuthorityMockRecorder is the mock recorder for MockAuthority
type MockAuthorityMockRecorder struct {
	mock *MockAuthority
}

// NewMockAuthority creates a new mock instance
func NewMockAuthority(ctrl *gomock.Controller) *MockAuthority {
	mock := &MockAuthority{ctrl: ctrl}
	mock.recorder = &MockAuthorityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthority) EXPECT() *MockAuthorityMockRecorder {
	return m.recorder
}

// VerifyToken mocks base method
func (m *MockAuthority) VerifyToken(actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyToken indicates an expected call of VerifyToken
func (mr *MockAuthorityMockRecorder) VerifyToken(actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockAuthority)(nil).VerifyToken), actualToken)
}

// VerifyNamespace mocks base method
func (m *MockAuthority) VerifyNamespace(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyNamespace", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyNamespace indicates an expected call of VerifyNamespace
func (mr *MockAuthorityMockRecorder) VerifyNamespace(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyNamespace", reflect.TypeOf((*MockAuthority)(nil).VerifyNamespace), expectToken, actualToken)
}

// VerifyService mocks base method
func (m *MockAuthority) VerifyService(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyService", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyService indicates an expected call of VerifyService
func (mr *MockAuthorityMockRecorder) VerifyService(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyService", reflect.TypeOf((*MockAuthority)(nil).VerifyService), expectToken, actualToken)
}

// VerifyInstance mocks base method
func (m *MockAuthority) VerifyInstance(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyInstance", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyInstance indicates an expected call of VerifyInstance
func (mr *MockAuthorityMockRecorder) VerifyInstance(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyInstance", reflect.TypeOf((*MockAuthority)(nil).VerifyInstance), expectToken, actualToken)
}

// VerifyRule mocks base method
func (m *MockAuthority) VerifyRule(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyRule", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyRule indicates an expected call of VerifyRule
func (mr *MockAuthorityMockRecorder) VerifyRule(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyRule", reflect.TypeOf((*MockAuthority)(nil).VerifyRule), expectToken, actualToken)
}

// VerifyPlatform mocks base method
func (m *MockAuthority) VerifyPlatform(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPlatform", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyPlatform indicates an expected call of VerifyPlatform
func (mr *MockAuthorityMockRecorder) VerifyPlatform(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPlatform", reflect.TypeOf((*MockAuthority)(nil).VerifyPlatform), expectToken, actualToken)
}

// VerifyMesh mocks base method
func (m *MockAuthority) VerifyMesh(expectToken, actualToken string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyMesh", expectToken, actualToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyMesh indicates an expected call of VerifyMesh
func (mr *MockAuthorityMockRecorder) VerifyMesh(expectToken, actualToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMesh", reflect.TypeOf((*MockAuthority)(nil).VerifyMesh), expectToken, actualToken)
}

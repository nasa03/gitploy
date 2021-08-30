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

// ActivateRepo mocks base method.
func (m *MockInteractor) ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateRepo", ctx, u, r, c)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateRepo indicates an expected call of ActivateRepo.
func (mr *MockInteractorMockRecorder) ActivateRepo(ctx, u, r, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateRepo", reflect.TypeOf((*MockInteractor)(nil).ActivateRepo), ctx, u, r, c)
}

// CompareCommits mocks base method.
func (m *MockInteractor) CompareCommits(ctx context.Context, u *ent.User, r *ent.Repo, base, head string, page, perPage int) ([]*vo.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareCommits", ctx, u, r, base, head, page, perPage)
	ret0, _ := ret[0].([]*vo.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompareCommits indicates an expected call of CompareCommits.
func (mr *MockInteractorMockRecorder) CompareCommits(ctx, u, r, base, head, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareCommits", reflect.TypeOf((*MockInteractor)(nil).CompareCommits), ctx, u, r, base, head, page, perPage)
}

// CreateApproval mocks base method.
func (m *MockInteractor) CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApproval", ctx, a)
	ret0, _ := ret[0].(*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApproval indicates an expected call of CreateApproval.
func (mr *MockInteractorMockRecorder) CreateApproval(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApproval", reflect.TypeOf((*MockInteractor)(nil).CreateApproval), ctx, a)
}

// CreateEvent mocks base method.
func (m *MockInteractor) CreateEvent(ctx context.Context, e *ent.Event) (*ent.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", ctx, e)
	ret0, _ := ret[0].(*ent.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockInteractorMockRecorder) CreateEvent(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockInteractor)(nil).CreateEvent), ctx, e)
}

// CreateRemoteDeployment mocks base method.
func (m *MockInteractor) CreateRemoteDeployment(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRemoteDeployment", ctx, u, re, d, env)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRemoteDeployment indicates an expected call of CreateRemoteDeployment.
func (mr *MockInteractorMockRecorder) CreateRemoteDeployment(ctx, u, re, d, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRemoteDeployment", reflect.TypeOf((*MockInteractor)(nil).CreateRemoteDeployment), ctx, u, re, d, env)
}

// DeactivateRepo mocks base method.
func (m *MockInteractor) DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeactivateRepo", ctx, u, r)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeactivateRepo indicates an expected call of DeactivateRepo.
func (mr *MockInteractorMockRecorder) DeactivateRepo(ctx, u, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateRepo", reflect.TypeOf((*MockInteractor)(nil).DeactivateRepo), ctx, u, r)
}

// DeleteApproval mocks base method.
func (m *MockInteractor) DeleteApproval(ctx context.Context, a *ent.Approval) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApproval", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApproval indicates an expected call of DeleteApproval.
func (mr *MockInteractorMockRecorder) DeleteApproval(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApproval", reflect.TypeOf((*MockInteractor)(nil).DeleteApproval), ctx, a)
}

// Deploy mocks base method.
func (m *MockInteractor) Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", ctx, u, re, d, env)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deploy indicates an expected call of Deploy.
func (mr *MockInteractorMockRecorder) Deploy(ctx, u, re, d, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockInteractor)(nil).Deploy), ctx, u, re, d, env)
}

// FindApprovalByID mocks base method.
func (m *MockInteractor) FindApprovalByID(ctx context.Context, id int) (*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindApprovalByID", ctx, id)
	ret0, _ := ret[0].(*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindApprovalByID indicates an expected call of FindApprovalByID.
func (mr *MockInteractorMockRecorder) FindApprovalByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindApprovalByID", reflect.TypeOf((*MockInteractor)(nil).FindApprovalByID), ctx, id)
}

// FindApprovalOfUser mocks base method.
func (m *MockInteractor) FindApprovalOfUser(ctx context.Context, d *ent.Deployment, u *ent.User) (*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindApprovalOfUser", ctx, d, u)
	ret0, _ := ret[0].(*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindApprovalOfUser indicates an expected call of FindApprovalOfUser.
func (mr *MockInteractorMockRecorder) FindApprovalOfUser(ctx, d, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindApprovalOfUser", reflect.TypeOf((*MockInteractor)(nil).FindApprovalOfUser), ctx, d, u)
}

// FindDeploymentByID mocks base method.
func (m *MockInteractor) FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentByID", ctx, id)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDeploymentByID indicates an expected call of FindDeploymentByID.
func (mr *MockInteractorMockRecorder) FindDeploymentByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentByID", reflect.TypeOf((*MockInteractor)(nil).FindDeploymentByID), ctx, id)
}

// FindDeploymentOfRepoByNumber mocks base method.
func (m *MockInteractor) FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentOfRepoByNumber", ctx, r, number)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDeploymentOfRepoByNumber indicates an expected call of FindDeploymentOfRepoByNumber.
func (mr *MockInteractorMockRecorder) FindDeploymentOfRepoByNumber(ctx, r, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentOfRepoByNumber", reflect.TypeOf((*MockInteractor)(nil).FindDeploymentOfRepoByNumber), ctx, r, number)
}

// FindPermOfRepo mocks base method.
func (m *MockInteractor) FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPermOfRepo", ctx, r, u)
	ret0, _ := ret[0].(*ent.Perm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPermOfRepo indicates an expected call of FindPermOfRepo.
func (mr *MockInteractorMockRecorder) FindPermOfRepo(ctx, r, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPermOfRepo", reflect.TypeOf((*MockInteractor)(nil).FindPermOfRepo), ctx, r, u)
}

// FindPrevSuccessDeployment mocks base method.
func (m *MockInteractor) FindPrevSuccessDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPrevSuccessDeployment", ctx, d)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPrevSuccessDeployment indicates an expected call of FindPrevSuccessDeployment.
func (mr *MockInteractorMockRecorder) FindPrevSuccessDeployment(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPrevSuccessDeployment", reflect.TypeOf((*MockInteractor)(nil).FindPrevSuccessDeployment), ctx, d)
}

// FindRepoOfUserByID mocks base method.
func (m *MockInteractor) FindRepoOfUserByID(ctx context.Context, u *ent.User, id string) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRepoOfUserByID", ctx, u, id)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRepoOfUserByID indicates an expected call of FindRepoOfUserByID.
func (mr *MockInteractorMockRecorder) FindRepoOfUserByID(ctx, u, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRepoOfUserByID", reflect.TypeOf((*MockInteractor)(nil).FindRepoOfUserByID), ctx, u, id)
}

// FindRepoOfUserByNamespaceName mocks base method.
func (m *MockInteractor) FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRepoOfUserByNamespaceName", ctx, u, namespace, name)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRepoOfUserByNamespaceName indicates an expected call of FindRepoOfUserByNamespaceName.
func (mr *MockInteractorMockRecorder) FindRepoOfUserByNamespaceName(ctx, u, namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRepoOfUserByNamespaceName", reflect.TypeOf((*MockInteractor)(nil).FindRepoOfUserByNamespaceName), ctx, u, namespace, name)
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

// GetBranch mocks base method.
func (m *MockInteractor) GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranch", ctx, u, r, branch)
	ret0, _ := ret[0].(*vo.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockInteractorMockRecorder) GetBranch(ctx, u, r, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockInteractor)(nil).GetBranch), ctx, u, r, branch)
}

// GetCommit mocks base method.
func (m *MockInteractor) GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommit", ctx, u, r, sha)
	ret0, _ := ret[0].(*vo.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommit indicates an expected call of GetCommit.
func (mr *MockInteractorMockRecorder) GetCommit(ctx, u, r, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockInteractor)(nil).GetCommit), ctx, u, r, sha)
}

// GetConfig mocks base method.
func (m *MockInteractor) GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", ctx, u, r)
	ret0, _ := ret[0].(*vo.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockInteractorMockRecorder) GetConfig(ctx, u, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockInteractor)(nil).GetConfig), ctx, u, r)
}

// GetNextDeploymentNumberOfRepo mocks base method.
func (m *MockInteractor) GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextDeploymentNumberOfRepo", ctx, r)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextDeploymentNumberOfRepo indicates an expected call of GetNextDeploymentNumberOfRepo.
func (mr *MockInteractorMockRecorder) GetNextDeploymentNumberOfRepo(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextDeploymentNumberOfRepo", reflect.TypeOf((*MockInteractor)(nil).GetNextDeploymentNumberOfRepo), ctx, r)
}

// GetTag mocks base method.
func (m *MockInteractor) GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTag", ctx, u, r, tag)
	ret0, _ := ret[0].(*vo.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTag indicates an expected call of GetTag.
func (mr *MockInteractorMockRecorder) GetTag(ctx, u, r, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTag", reflect.TypeOf((*MockInteractor)(nil).GetTag), ctx, u, r, tag)
}

// IsApproved mocks base method.
func (m *MockInteractor) IsApproved(ctx context.Context, d *ent.Deployment) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsApproved", ctx, d)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsApproved indicates an expected call of IsApproved.
func (mr *MockInteractorMockRecorder) IsApproved(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsApproved", reflect.TypeOf((*MockInteractor)(nil).IsApproved), ctx, d)
}

// ListApprovals mocks base method.
func (m *MockInteractor) ListApprovals(ctx context.Context, d *ent.Deployment) ([]*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApprovals", ctx, d)
	ret0, _ := ret[0].([]*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApprovals indicates an expected call of ListApprovals.
func (mr *MockInteractorMockRecorder) ListApprovals(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApprovals", reflect.TypeOf((*MockInteractor)(nil).ListApprovals), ctx, d)
}

// ListBranches mocks base method.
func (m *MockInteractor) ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBranches", ctx, u, r, page, perPage)
	ret0, _ := ret[0].([]*vo.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBranches indicates an expected call of ListBranches.
func (mr *MockInteractorMockRecorder) ListBranches(ctx, u, r, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranches", reflect.TypeOf((*MockInteractor)(nil).ListBranches), ctx, u, r, page, perPage)
}

// ListCommitStatuses mocks base method.
func (m *MockInteractor) ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*vo.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommitStatuses", ctx, u, r, sha)
	ret0, _ := ret[0].([]*vo.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommitStatuses indicates an expected call of ListCommitStatuses.
func (mr *MockInteractorMockRecorder) ListCommitStatuses(ctx, u, r, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommitStatuses", reflect.TypeOf((*MockInteractor)(nil).ListCommitStatuses), ctx, u, r, sha)
}

// ListCommits mocks base method.
func (m *MockInteractor) ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommits", ctx, u, r, branch, page, perPage)
	ret0, _ := ret[0].([]*vo.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommits indicates an expected call of ListCommits.
func (mr *MockInteractorMockRecorder) ListCommits(ctx, u, r, branch, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommits", reflect.TypeOf((*MockInteractor)(nil).ListCommits), ctx, u, r, branch, page, perPage)
}

// ListDeploymentsOfRepo mocks base method.
func (m *MockInteractor) ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, env, status string, page, perPage int) ([]*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeploymentsOfRepo", ctx, r, env, status, page, perPage)
	ret0, _ := ret[0].([]*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeploymentsOfRepo indicates an expected call of ListDeploymentsOfRepo.
func (mr *MockInteractorMockRecorder) ListDeploymentsOfRepo(ctx, r, env, status, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeploymentsOfRepo", reflect.TypeOf((*MockInteractor)(nil).ListDeploymentsOfRepo), ctx, r, env, status, page, perPage)
}

// ListPermsOfRepo mocks base method.
func (m *MockInteractor) ListPermsOfRepo(ctx context.Context, r *ent.Repo, q string, page, perPage int) ([]*ent.Perm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPermsOfRepo", ctx, r, q, page, perPage)
	ret0, _ := ret[0].([]*ent.Perm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPermsOfRepo indicates an expected call of ListPermsOfRepo.
func (mr *MockInteractorMockRecorder) ListPermsOfRepo(ctx, r, q, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPermsOfRepo", reflect.TypeOf((*MockInteractor)(nil).ListPermsOfRepo), ctx, r, q, page, perPage)
}

// ListReposOfUser mocks base method.
func (m *MockInteractor) ListReposOfUser(ctx context.Context, u *ent.User, sorted bool, q string, page, perPage int) ([]*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReposOfUser", ctx, u, sorted, q, page, perPage)
	ret0, _ := ret[0].([]*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReposOfUser indicates an expected call of ListReposOfUser.
func (mr *MockInteractorMockRecorder) ListReposOfUser(ctx, u, sorted, q, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReposOfUser", reflect.TypeOf((*MockInteractor)(nil).ListReposOfUser), ctx, u, sorted, q, page, perPage)
}

// ListTags mocks base method.
func (m *MockInteractor) ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags", ctx, u, r, page, perPage)
	ret0, _ := ret[0].([]*vo.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockInteractorMockRecorder) ListTags(ctx, u, r, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockInteractor)(nil).ListTags), ctx, u, r, page, perPage)
}

// Rollback mocks base method.
func (m *MockInteractor) Rollback(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", ctx, u, re, d, env)
	ret0, _ := ret[0].(*ent.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rollback indicates an expected call of Rollback.
func (mr *MockInteractorMockRecorder) Rollback(ctx, u, re, d, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockInteractor)(nil).Rollback), ctx, u, re, d, env)
}

// UpdateApproval mocks base method.
func (m *MockInteractor) UpdateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApproval", ctx, a)
	ret0, _ := ret[0].(*ent.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApproval indicates an expected call of UpdateApproval.
func (mr *MockInteractorMockRecorder) UpdateApproval(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApproval", reflect.TypeOf((*MockInteractor)(nil).UpdateApproval), ctx, a)
}

// UpdateRepo mocks base method.
func (m *MockInteractor) UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRepo", ctx, r)
	ret0, _ := ret[0].(*ent.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRepo indicates an expected call of UpdateRepo.
func (mr *MockInteractorMockRecorder) UpdateRepo(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRepo", reflect.TypeOf((*MockInteractor)(nil).UpdateRepo), ctx, r)
}

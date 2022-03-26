// Code generated by entc, DO NOT EDIT.

package deployment

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the deployment type in the database.
	Label = "deployment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldEnv holds the string denoting the env field in the database.
	FieldEnv = "env"
	// FieldRef holds the string denoting the ref field in the database.
	FieldRef = "ref"
	// FieldDynamicPayload holds the string denoting the dynamic_payload field in the database.
	FieldDynamicPayload = "dynamic_payload"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldSha holds the string denoting the sha field in the database.
	FieldSha = "sha"
	// FieldHTMLURL holds the string denoting the html_url field in the database.
	FieldHTMLURL = "html_url"
	// FieldProductionEnvironment holds the string denoting the production_environment field in the database.
	FieldProductionEnvironment = "production_environment"
	// FieldIsRollback holds the string denoting the is_rollback field in the database.
	FieldIsRollback = "is_rollback"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRepoID holds the string denoting the repo_id field in the database.
	FieldRepoID = "repo_id"
	// FieldIsApprovalEnabled holds the string denoting the is_approval_enabled field in the database.
	FieldIsApprovalEnabled = "is_approval_enabled"
	// FieldRequiredApprovalCount holds the string denoting the required_approval_count field in the database.
	FieldRequiredApprovalCount = "required_approval_count"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeRepo holds the string denoting the repo edge name in mutations.
	EdgeRepo = "repo"
	// EdgeReviews holds the string denoting the reviews edge name in mutations.
	EdgeReviews = "reviews"
	// EdgeDeploymentStatuses holds the string denoting the deployment_statuses edge name in mutations.
	EdgeDeploymentStatuses = "deployment_statuses"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the deployment in the database.
	Table = "deployments"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "deployments"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// RepoTable is the table that holds the repo relation/edge.
	RepoTable = "deployments"
	// RepoInverseTable is the table name for the Repo entity.
	// It exists in this package in order to avoid circular dependency with the "repo" package.
	RepoInverseTable = "repos"
	// RepoColumn is the table column denoting the repo relation/edge.
	RepoColumn = "repo_id"
	// ReviewsTable is the table that holds the reviews relation/edge.
	ReviewsTable = "reviews"
	// ReviewsInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewsInverseTable = "reviews"
	// ReviewsColumn is the table column denoting the reviews relation/edge.
	ReviewsColumn = "deployment_id"
	// DeploymentStatusesTable is the table that holds the deployment_statuses relation/edge.
	DeploymentStatusesTable = "deployment_status"
	// DeploymentStatusesInverseTable is the table name for the DeploymentStatus entity.
	// It exists in this package in order to avoid circular dependency with the "deploymentstatus" package.
	DeploymentStatusesInverseTable = "deployment_status"
	// DeploymentStatusesColumn is the table column denoting the deployment_statuses relation/edge.
	DeploymentStatusesColumn = "deployment_id"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "events"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "deployment_id"
)

// Columns holds all SQL columns for deployment fields.
var Columns = []string{
	FieldID,
	FieldNumber,
	FieldType,
	FieldEnv,
	FieldRef,
	FieldDynamicPayload,
	FieldStatus,
	FieldUID,
	FieldSha,
	FieldHTMLURL,
	FieldProductionEnvironment,
	FieldIsRollback,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
	FieldRepoID,
	FieldIsApprovalEnabled,
	FieldRequiredApprovalCount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// HTMLURLValidator is a validator for the "html_url" field. It is called by the builders before save.
	HTMLURLValidator func(string) error
	// DefaultProductionEnvironment holds the default value on creation for the "production_environment" field.
	DefaultProductionEnvironment bool
	// DefaultIsRollback holds the default value on creation for the "is_rollback" field.
	DefaultIsRollback bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// TypeCommit is the default value of the Type enum.
const DefaultType = TypeCommit

// Type values.
const (
	TypeCommit Type = "commit"
	TypeBranch Type = "branch"
	TypeTag    Type = "tag"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeCommit, TypeBranch, TypeTag:
		return nil
	default:
		return fmt.Errorf("deployment: invalid enum value for type field: %q", _type)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// StatusWaiting is the default value of the Status enum.
const DefaultStatus = StatusWaiting

// Status values.
const (
	StatusWaiting  Status = "waiting"
	StatusCreated  Status = "created"
	StatusQueued   Status = "queued"
	StatusRunning  Status = "running"
	StatusSuccess  Status = "success"
	StatusFailure  Status = "failure"
	StatusCanceled Status = "canceled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusWaiting, StatusCreated, StatusQueued, StatusRunning, StatusSuccess, StatusFailure, StatusCanceled:
		return nil
	default:
		return fmt.Errorf("deployment: invalid enum value for status field: %q", s)
	}
}

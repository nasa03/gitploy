// Code generated by entc, DO NOT EDIT.

package approval

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the approval type in the database.
	Label = "approval"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldDeploymentID holds the string denoting the deployment_id field in the database.
	FieldDeploymentID = "deployment_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeDeployment holds the string denoting the deployment edge name in mutations.
	EdgeDeployment = "deployment"
	// Table holds the table name of the approval in the database.
	Table = "approvals"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "approvals"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// DeploymentTable is the table the holds the deployment relation/edge.
	DeploymentTable = "approvals"
	// DeploymentInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentInverseTable = "deployments"
	// DeploymentColumn is the table column denoting the deployment relation/edge.
	DeploymentColumn = "deployment_id"
)

// Columns holds all SQL columns for approval fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
	FieldDeploymentID,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPending is the default value of the Status enum.
const DefaultStatus = StatusPending

// Status values.
const (
	StatusPending  Status = "pending"
	StatusDeclined Status = "declined"
	StatusApproved Status = "approved"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusDeclined, StatusApproved:
		return nil
	default:
		return fmt.Errorf("approval: invalid enum value for status field: %q", s)
	}
}

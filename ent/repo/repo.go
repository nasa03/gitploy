// Code generated by entc, DO NOT EDIT.

package repo

import (
	"time"
)

const (
	// Label holds the string label denoting the repo type in the database.
	Label = "repo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldConfigPath holds the string denoting the config_path field in the database.
	FieldConfigPath = "config_path"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldWebhookID holds the string denoting the webhook_id field in the database.
	FieldWebhookID = "webhook_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldLatestDeployedAt holds the string denoting the latest_deployed_at field in the database.
	FieldLatestDeployedAt = "latest_deployed_at"
	// EdgePerms holds the string denoting the perms edge name in mutations.
	EdgePerms = "perms"
	// EdgeDeployments holds the string denoting the deployments edge name in mutations.
	EdgeDeployments = "deployments"
	// EdgeCallback holds the string denoting the callback edge name in mutations.
	EdgeCallback = "callback"
	// EdgeLocks holds the string denoting the locks edge name in mutations.
	EdgeLocks = "locks"
	// Table holds the table name of the repo in the database.
	Table = "repos"
	// PermsTable is the table that holds the perms relation/edge.
	PermsTable = "perms"
	// PermsInverseTable is the table name for the Perm entity.
	// It exists in this package in order to avoid circular dependency with the "perm" package.
	PermsInverseTable = "perms"
	// PermsColumn is the table column denoting the perms relation/edge.
	PermsColumn = "repo_id"
	// DeploymentsTable is the table that holds the deployments relation/edge.
	DeploymentsTable = "deployments"
	// DeploymentsInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentsInverseTable = "deployments"
	// DeploymentsColumn is the table column denoting the deployments relation/edge.
	DeploymentsColumn = "repo_id"
	// CallbackTable is the table that holds the callback relation/edge.
	CallbackTable = "callbacks"
	// CallbackInverseTable is the table name for the Callback entity.
	// It exists in this package in order to avoid circular dependency with the "callback" package.
	CallbackInverseTable = "callbacks"
	// CallbackColumn is the table column denoting the callback relation/edge.
	CallbackColumn = "repo_id"
	// LocksTable is the table that holds the locks relation/edge.
	LocksTable = "locks"
	// LocksInverseTable is the table name for the Lock entity.
	// It exists in this package in order to avoid circular dependency with the "lock" package.
	LocksInverseTable = "locks"
	// LocksColumn is the table column denoting the locks relation/edge.
	LocksColumn = "repo_id"
)

// Columns holds all SQL columns for repo fields.
var Columns = []string{
	FieldID,
	FieldNamespace,
	FieldName,
	FieldDescription,
	FieldConfigPath,
	FieldActive,
	FieldWebhookID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLatestDeployedAt,
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
	// DefaultConfigPath holds the default value on creation for the "config_path" field.
	DefaultConfigPath string
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

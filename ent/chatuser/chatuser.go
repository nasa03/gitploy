// Code generated by entc, DO NOT EDIT.

package chatuser

import (
	"time"
)

const (
	// Label holds the string label denoting the chatuser type in the database.
	Label = "chat_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldRefresh holds the string denoting the refresh field in the database.
	FieldRefresh = "refresh"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// FieldBotToken holds the string denoting the bot_token field in the database.
	FieldBotToken = "bot_token"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the chatuser in the database.
	Table = "chat_users"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "chat_users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for chatuser fields.
var Columns = []string{
	FieldID,
	FieldToken,
	FieldRefresh,
	FieldExpiry,
	FieldBotToken,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
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

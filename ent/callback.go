// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gitploy-io/gitploy/ent/callback"
	"github.com/gitploy-io/gitploy/ent/repo"
)

// Callback is the model entity for the Callback schema.
type Callback struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"-"`
	// Type holds the value of the "type" field.
	Type callback.Type `json:"type"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// RepoID holds the value of the "repo_id" field.
	RepoID int64 `json:"repo_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CallbackQuery when eager-loading is set.
	Edges CallbackEdges `json:"edges"`
}

// CallbackEdges holds the relations/edges for other nodes in the graph.
type CallbackEdges struct {
	// Repo holds the value of the repo edge.
	Repo *Repo `json:"repo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RepoOrErr returns the Repo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CallbackEdges) RepoOrErr() (*Repo, error) {
	if e.loadedTypes[0] {
		if e.Repo == nil {
			// The edge repo was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: repo.Label}
		}
		return e.Repo, nil
	}
	return nil, &NotLoadedError{edge: "repo"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Callback) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case callback.FieldID, callback.FieldRepoID:
			values[i] = new(sql.NullInt64)
		case callback.FieldHash, callback.FieldType:
			values[i] = new(sql.NullString)
		case callback.FieldCreatedAt, callback.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Callback", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Callback fields.
func (c *Callback) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case callback.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case callback.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				c.Hash = value.String
			}
		case callback.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = callback.Type(value.String)
			}
		case callback.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case callback.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case callback.FieldRepoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field repo_id", values[i])
			} else if value.Valid {
				c.RepoID = value.Int64
			}
		}
	}
	return nil
}

// QueryRepo queries the "repo" edge of the Callback entity.
func (c *Callback) QueryRepo() *RepoQuery {
	return (&CallbackClient{config: c.config}).QueryRepo(c)
}

// Update returns a builder for updating this Callback.
// Note that you need to call Callback.Unwrap() before calling this method if this Callback
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Callback) Update() *CallbackUpdateOne {
	return (&CallbackClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Callback entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Callback) Unwrap() *Callback {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Callback is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Callback) String() string {
	var builder strings.Builder
	builder.WriteString("Callback(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", hash=<sensitive>")
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", repo_id=")
	builder.WriteString(fmt.Sprintf("%v", c.RepoID))
	builder.WriteByte(')')
	return builder.String()
}

// Callbacks is a parsable slice of Callback.
type Callbacks []*Callback

func (c Callbacks) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}

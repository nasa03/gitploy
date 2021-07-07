// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/repo"
	"github.com/hanjunlee/gitploy/ent/user"
)

// Deployment is the model entity for the Deployment schema.
type Deployment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Number holds the value of the "number" field.
	Number int `json:"number,omitempty"`
	// UID holds the value of the "uid" field.
	UID int64 `json:"uid,omitempty"`
	// Type holds the value of the "type" field.
	Type deployment.Type `json:"type,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref string `json:"ref,omitempty"`
	// Sha holds the value of the "sha" field.
	Sha string `json:"sha,omitempty"`
	// Env holds the value of the "env" field.
	Env string `json:"env,omitempty"`
	// Status holds the value of the "status" field.
	Status deployment.Status `json:"status,omitempty"`
	// RequiredApprovalCount holds the value of the "required_approval_count" field.
	RequiredApprovalCount int `json:"required_approval_count,omitempty"`
	// AutoDeploy holds the value of the "auto_deploy" field.
	AutoDeploy bool `json:"auto_deploy,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// RepoID holds the value of the "repo_id" field.
	RepoID string `json:"repo_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeploymentQuery when eager-loading is set.
	Edges DeploymentEdges `json:"edges"`
}

// DeploymentEdges holds the relations/edges for other nodes in the graph.
type DeploymentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Repo holds the value of the repo edge.
	Repo *Repo `json:"repo,omitempty"`
	// Approvals holds the value of the approvals edge.
	Approvals []*Approval `json:"approvals,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeploymentEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RepoOrErr returns the Repo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeploymentEdges) RepoOrErr() (*Repo, error) {
	if e.loadedTypes[1] {
		if e.Repo == nil {
			// The edge repo was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: repo.Label}
		}
		return e.Repo, nil
	}
	return nil, &NotLoadedError{edge: "repo"}
}

// ApprovalsOrErr returns the Approvals value or an error if the edge
// was not loaded in eager-loading.
func (e DeploymentEdges) ApprovalsOrErr() ([]*Approval, error) {
	if e.loadedTypes[2] {
		return e.Approvals, nil
	}
	return nil, &NotLoadedError{edge: "approvals"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e DeploymentEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[3] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Deployment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case deployment.FieldAutoDeploy:
			values[i] = new(sql.NullBool)
		case deployment.FieldID, deployment.FieldNumber, deployment.FieldUID, deployment.FieldRequiredApprovalCount:
			values[i] = new(sql.NullInt64)
		case deployment.FieldType, deployment.FieldRef, deployment.FieldSha, deployment.FieldEnv, deployment.FieldStatus, deployment.FieldUserID, deployment.FieldRepoID:
			values[i] = new(sql.NullString)
		case deployment.FieldCreatedAt, deployment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Deployment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Deployment fields.
func (d *Deployment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deployment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case deployment.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				d.Number = int(value.Int64)
			}
		case deployment.FieldUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				d.UID = value.Int64
			}
		case deployment.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				d.Type = deployment.Type(value.String)
			}
		case deployment.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				d.Ref = value.String
			}
		case deployment.FieldSha:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sha", values[i])
			} else if value.Valid {
				d.Sha = value.String
			}
		case deployment.FieldEnv:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field env", values[i])
			} else if value.Valid {
				d.Env = value.String
			}
		case deployment.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				d.Status = deployment.Status(value.String)
			}
		case deployment.FieldRequiredApprovalCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field required_approval_count", values[i])
			} else if value.Valid {
				d.RequiredApprovalCount = int(value.Int64)
			}
		case deployment.FieldAutoDeploy:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field auto_deploy", values[i])
			} else if value.Valid {
				d.AutoDeploy = value.Bool
			}
		case deployment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case deployment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case deployment.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				d.UserID = value.String
			}
		case deployment.FieldRepoID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repo_id", values[i])
			} else if value.Valid {
				d.RepoID = value.String
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Deployment entity.
func (d *Deployment) QueryUser() *UserQuery {
	return (&DeploymentClient{config: d.config}).QueryUser(d)
}

// QueryRepo queries the "repo" edge of the Deployment entity.
func (d *Deployment) QueryRepo() *RepoQuery {
	return (&DeploymentClient{config: d.config}).QueryRepo(d)
}

// QueryApprovals queries the "approvals" edge of the Deployment entity.
func (d *Deployment) QueryApprovals() *ApprovalQuery {
	return (&DeploymentClient{config: d.config}).QueryApprovals(d)
}

// QueryNotifications queries the "notifications" edge of the Deployment entity.
func (d *Deployment) QueryNotifications() *NotificationQuery {
	return (&DeploymentClient{config: d.config}).QueryNotifications(d)
}

// Update returns a builder for updating this Deployment.
// Note that you need to call Deployment.Unwrap() before calling this method if this Deployment
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Deployment) Update() *DeploymentUpdateOne {
	return (&DeploymentClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Deployment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Deployment) Unwrap() *Deployment {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Deployment is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Deployment) String() string {
	var builder strings.Builder
	builder.WriteString("Deployment(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", number=")
	builder.WriteString(fmt.Sprintf("%v", d.Number))
	builder.WriteString(", uid=")
	builder.WriteString(fmt.Sprintf("%v", d.UID))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", d.Type))
	builder.WriteString(", ref=")
	builder.WriteString(d.Ref)
	builder.WriteString(", sha=")
	builder.WriteString(d.Sha)
	builder.WriteString(", env=")
	builder.WriteString(d.Env)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", d.Status))
	builder.WriteString(", required_approval_count=")
	builder.WriteString(fmt.Sprintf("%v", d.RequiredApprovalCount))
	builder.WriteString(", auto_deploy=")
	builder.WriteString(fmt.Sprintf("%v", d.AutoDeploy))
	builder.WriteString(", created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(d.UserID)
	builder.WriteString(", repo_id=")
	builder.WriteString(d.RepoID)
	builder.WriteByte(')')
	return builder.String()
}

// Deployments is a parsable slice of Deployment.
type Deployments []*Deployment

func (d Deployments) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}

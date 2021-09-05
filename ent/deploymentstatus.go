// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/deploymentstatus"
)

// DeploymentStatus is the model entity for the DeploymentStatus schema.
type DeploymentStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitemtpy"`
	// LogURL holds the value of the "log_url" field.
	LogURL string `json:"log_url,omitemtpy"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// DeploymentID holds the value of the "deployment_id" field.
	DeploymentID int `json:"deployment_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeploymentStatusQuery when eager-loading is set.
	Edges DeploymentStatusEdges `json:"edges"`
}

// DeploymentStatusEdges holds the relations/edges for other nodes in the graph.
type DeploymentStatusEdges struct {
	// Deployment holds the value of the deployment edge.
	Deployment *Deployment `json:"deployment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DeploymentOrErr returns the Deployment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeploymentStatusEdges) DeploymentOrErr() (*Deployment, error) {
	if e.loadedTypes[0] {
		if e.Deployment == nil {
			// The edge deployment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: deployment.Label}
		}
		return e.Deployment, nil
	}
	return nil, &NotLoadedError{edge: "deployment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeploymentStatus) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case deploymentstatus.FieldID, deploymentstatus.FieldDeploymentID:
			values[i] = new(sql.NullInt64)
		case deploymentstatus.FieldStatus, deploymentstatus.FieldDescription, deploymentstatus.FieldLogURL:
			values[i] = new(sql.NullString)
		case deploymentstatus.FieldCreatedAt, deploymentstatus.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DeploymentStatus", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeploymentStatus fields.
func (ds *DeploymentStatus) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deploymentstatus.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ds.ID = int(value.Int64)
		case deploymentstatus.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ds.Status = value.String
			}
		case deploymentstatus.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ds.Description = value.String
			}
		case deploymentstatus.FieldLogURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field log_url", values[i])
			} else if value.Valid {
				ds.LogURL = value.String
			}
		case deploymentstatus.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ds.CreatedAt = value.Time
			}
		case deploymentstatus.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ds.UpdatedAt = value.Time
			}
		case deploymentstatus.FieldDeploymentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_id", values[i])
			} else if value.Valid {
				ds.DeploymentID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryDeployment queries the "deployment" edge of the DeploymentStatus entity.
func (ds *DeploymentStatus) QueryDeployment() *DeploymentQuery {
	return (&DeploymentStatusClient{config: ds.config}).QueryDeployment(ds)
}

// Update returns a builder for updating this DeploymentStatus.
// Note that you need to call DeploymentStatus.Unwrap() before calling this method if this DeploymentStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (ds *DeploymentStatus) Update() *DeploymentStatusUpdateOne {
	return (&DeploymentStatusClient{config: ds.config}).UpdateOne(ds)
}

// Unwrap unwraps the DeploymentStatus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ds *DeploymentStatus) Unwrap() *DeploymentStatus {
	tx, ok := ds.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeploymentStatus is not a transactional entity")
	}
	ds.config.driver = tx.drv
	return ds
}

// String implements the fmt.Stringer.
func (ds *DeploymentStatus) String() string {
	var builder strings.Builder
	builder.WriteString("DeploymentStatus(")
	builder.WriteString(fmt.Sprintf("id=%v", ds.ID))
	builder.WriteString(", status=")
	builder.WriteString(ds.Status)
	builder.WriteString(", description=")
	builder.WriteString(ds.Description)
	builder.WriteString(", log_url=")
	builder.WriteString(ds.LogURL)
	builder.WriteString(", created_at=")
	builder.WriteString(ds.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ds.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deployment_id=")
	builder.WriteString(fmt.Sprintf("%v", ds.DeploymentID))
	builder.WriteByte(')')
	return builder.String()
}

// DeploymentStatusSlice is a parsable slice of DeploymentStatus.
type DeploymentStatusSlice []*DeploymentStatus

func (ds DeploymentStatusSlice) config(cfg config) {
	for _i := range ds {
		ds[_i].config = cfg
	}
}

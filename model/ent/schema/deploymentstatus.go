package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DeploymentStatus holds the schema definition for the DeploymentStatus entity.
type DeploymentStatus struct {
	ent.Schema
}

// Fields of the DeploymentStatus.
func (DeploymentStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("status"),
		field.String("description").
			Optional(),
		field.String("log_url").
			Optional(),
		field.Time("created_at").
			Default(nowUTC),
		field.Time("updated_at").
			Default(nowUTC).
			UpdateDefault(nowUTC),

		// edges
		field.Int("deployment_id"),
	}
}

// Edges of the DeploymentStatus.
func (DeploymentStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployment", Deployment.Type).
			Ref("deployment_statuses").
			Field("deployment_id").
			Unique().
			Required(),
	}
}
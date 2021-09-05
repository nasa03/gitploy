// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/ent/deploymentstatus"
	"github.com/gitploy-io/gitploy/ent/predicate"
)

// DeploymentStatusDelete is the builder for deleting a DeploymentStatus entity.
type DeploymentStatusDelete struct {
	config
	hooks    []Hook
	mutation *DeploymentStatusMutation
}

// Where appends a list predicates to the DeploymentStatusDelete builder.
func (dsd *DeploymentStatusDelete) Where(ps ...predicate.DeploymentStatus) *DeploymentStatusDelete {
	dsd.mutation.Where(ps...)
	return dsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dsd *DeploymentStatusDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dsd.hooks) == 0 {
		affected, err = dsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dsd.mutation = mutation
			affected, err = dsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dsd.hooks) - 1; i >= 0; i-- {
			if dsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dsd *DeploymentStatusDelete) ExecX(ctx context.Context) int {
	n, err := dsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dsd *DeploymentStatusDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: deploymentstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deploymentstatus.FieldID,
			},
		},
	}
	if ps := dsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dsd.driver, _spec)
}

// DeploymentStatusDeleteOne is the builder for deleting a single DeploymentStatus entity.
type DeploymentStatusDeleteOne struct {
	dsd *DeploymentStatusDelete
}

// Exec executes the deletion query.
func (dsdo *DeploymentStatusDeleteOne) Exec(ctx context.Context) error {
	n, err := dsdo.dsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{deploymentstatus.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dsdo *DeploymentStatusDeleteOne) ExecX(ctx context.Context) {
	dsdo.dsd.ExecX(ctx)
}

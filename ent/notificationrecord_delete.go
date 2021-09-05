// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/ent/notificationrecord"
	"github.com/gitploy-io/gitploy/ent/predicate"
)

// NotificationRecordDelete is the builder for deleting a NotificationRecord entity.
type NotificationRecordDelete struct {
	config
	hooks    []Hook
	mutation *NotificationRecordMutation
}

// Where appends a list predicates to the NotificationRecordDelete builder.
func (nrd *NotificationRecordDelete) Where(ps ...predicate.NotificationRecord) *NotificationRecordDelete {
	nrd.mutation.Where(ps...)
	return nrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nrd *NotificationRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nrd.hooks) == 0 {
		affected, err = nrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nrd.mutation = mutation
			affected, err = nrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nrd.hooks) - 1; i >= 0; i-- {
			if nrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nrd *NotificationRecordDelete) ExecX(ctx context.Context) int {
	n, err := nrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nrd *NotificationRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: notificationrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: notificationrecord.FieldID,
			},
		},
	}
	if ps := nrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, nrd.driver, _spec)
}

// NotificationRecordDeleteOne is the builder for deleting a single NotificationRecord entity.
type NotificationRecordDeleteOne struct {
	nrd *NotificationRecordDelete
}

// Exec executes the deletion query.
func (nrdo *NotificationRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := nrdo.nrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{notificationrecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nrdo *NotificationRecordDeleteOne) ExecX(ctx context.Context) {
	nrdo.nrd.ExecX(ctx)
}

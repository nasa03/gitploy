// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/predicate"
)

// ChatCallbackDelete is the builder for deleting a ChatCallback entity.
type ChatCallbackDelete struct {
	config
	hooks    []Hook
	mutation *ChatCallbackMutation
}

// Where adds a new predicate to the ChatCallbackDelete builder.
func (ccd *ChatCallbackDelete) Where(ps ...predicate.ChatCallback) *ChatCallbackDelete {
	ccd.mutation.predicates = append(ccd.mutation.predicates, ps...)
	return ccd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ccd *ChatCallbackDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ccd.hooks) == 0 {
		affected, err = ccd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatCallbackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccd.mutation = mutation
			affected, err = ccd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccd.hooks) - 1; i >= 0; i-- {
			mut = ccd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccd *ChatCallbackDelete) ExecX(ctx context.Context) int {
	n, err := ccd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ccd *ChatCallbackDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: chatcallback.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: chatcallback.FieldID,
			},
		},
	}
	if ps := ccd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ccd.driver, _spec)
}

// ChatCallbackDeleteOne is the builder for deleting a single ChatCallback entity.
type ChatCallbackDeleteOne struct {
	ccd *ChatCallbackDelete
}

// Exec executes the deletion query.
func (ccdo *ChatCallbackDeleteOne) Exec(ctx context.Context) error {
	n, err := ccdo.ccd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{chatcallback.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ccdo *ChatCallbackDeleteOne) ExecX(ctx context.Context) {
	ccdo.ccd.ExecX(ctx)
}

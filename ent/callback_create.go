// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/ent/callback"
	"github.com/gitploy-io/gitploy/ent/repo"
)

// CallbackCreate is the builder for creating a Callback entity.
type CallbackCreate struct {
	config
	mutation *CallbackMutation
	hooks    []Hook
}

// SetHash sets the "hash" field.
func (cc *CallbackCreate) SetHash(s string) *CallbackCreate {
	cc.mutation.SetHash(s)
	return cc
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (cc *CallbackCreate) SetNillableHash(s *string) *CallbackCreate {
	if s != nil {
		cc.SetHash(*s)
	}
	return cc
}

// SetType sets the "type" field.
func (cc *CallbackCreate) SetType(c callback.Type) *CallbackCreate {
	cc.mutation.SetType(c)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CallbackCreate) SetCreatedAt(t time.Time) *CallbackCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CallbackCreate) SetNillableCreatedAt(t *time.Time) *CallbackCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CallbackCreate) SetUpdatedAt(t time.Time) *CallbackCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CallbackCreate) SetNillableUpdatedAt(t *time.Time) *CallbackCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetRepoID sets the "repo_id" field.
func (cc *CallbackCreate) SetRepoID(i int64) *CallbackCreate {
	cc.mutation.SetRepoID(i)
	return cc
}

// SetRepo sets the "repo" edge to the Repo entity.
func (cc *CallbackCreate) SetRepo(r *Repo) *CallbackCreate {
	return cc.SetRepoID(r.ID)
}

// Mutation returns the CallbackMutation object of the builder.
func (cc *CallbackCreate) Mutation() *CallbackMutation {
	return cc.mutation
}

// Save creates the Callback in the database.
func (cc *CallbackCreate) Save(ctx context.Context) (*Callback, error) {
	var (
		err  error
		node *Callback
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CallbackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CallbackCreate) SaveX(ctx context.Context) *Callback {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CallbackCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CallbackCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CallbackCreate) defaults() {
	if _, ok := cc.mutation.Hash(); !ok {
		v := callback.DefaultHash()
		cc.mutation.SetHash(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := callback.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := callback.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CallbackCreate) check() error {
	if _, ok := cc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "hash"`)}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := callback.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := cc.mutation.RepoID(); !ok {
		return &ValidationError{Name: "repo_id", err: errors.New(`ent: missing required field "repo_id"`)}
	}
	if _, ok := cc.mutation.RepoID(); !ok {
		return &ValidationError{Name: "repo", err: errors.New("ent: missing required edge \"repo\"")}
	}
	return nil
}

func (cc *CallbackCreate) sqlSave(ctx context.Context) (*Callback, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CallbackCreate) createSpec() (*Callback, *sqlgraph.CreateSpec) {
	var (
		_node = &Callback{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: callback.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: callback.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: callback.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: callback.FieldType,
		})
		_node.Type = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: callback.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: callback.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.RepoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   callback.RepoTable,
			Columns: []string{callback.RepoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: repo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RepoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CallbackCreateBulk is the builder for creating many Callback entities in bulk.
type CallbackCreateBulk struct {
	config
	builders []*CallbackCreate
}

// Save creates the Callback entities in the database.
func (ccb *CallbackCreateBulk) Save(ctx context.Context) ([]*Callback, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Callback, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CallbackMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CallbackCreateBulk) SaveX(ctx context.Context) []*Callback {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CallbackCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CallbackCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

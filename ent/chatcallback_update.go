// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/chatuser"
	"github.com/hanjunlee/gitploy/ent/predicate"
	"github.com/hanjunlee/gitploy/ent/repo"
)

// ChatCallbackUpdate is the builder for updating ChatCallback entities.
type ChatCallbackUpdate struct {
	config
	hooks    []Hook
	mutation *ChatCallbackMutation
}

// Where adds a new predicate for the ChatCallbackUpdate builder.
func (ccu *ChatCallbackUpdate) Where(ps ...predicate.ChatCallback) *ChatCallbackUpdate {
	ccu.mutation.predicates = append(ccu.mutation.predicates, ps...)
	return ccu
}

// SetState sets the "state" field.
func (ccu *ChatCallbackUpdate) SetState(s string) *ChatCallbackUpdate {
	ccu.mutation.SetState(s)
	return ccu
}

// SetType sets the "type" field.
func (ccu *ChatCallbackUpdate) SetType(c chatcallback.Type) *ChatCallbackUpdate {
	ccu.mutation.SetType(c)
	return ccu
}

// SetIsOpened sets the "is_opened" field.
func (ccu *ChatCallbackUpdate) SetIsOpened(b bool) *ChatCallbackUpdate {
	ccu.mutation.SetIsOpened(b)
	return ccu
}

// SetNillableIsOpened sets the "is_opened" field if the given value is not nil.
func (ccu *ChatCallbackUpdate) SetNillableIsOpened(b *bool) *ChatCallbackUpdate {
	if b != nil {
		ccu.SetIsOpened(*b)
	}
	return ccu
}

// SetCreatedAt sets the "created_at" field.
func (ccu *ChatCallbackUpdate) SetCreatedAt(t time.Time) *ChatCallbackUpdate {
	ccu.mutation.SetCreatedAt(t)
	return ccu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccu *ChatCallbackUpdate) SetNillableCreatedAt(t *time.Time) *ChatCallbackUpdate {
	if t != nil {
		ccu.SetCreatedAt(*t)
	}
	return ccu
}

// SetUpdatedAt sets the "updated_at" field.
func (ccu *ChatCallbackUpdate) SetUpdatedAt(t time.Time) *ChatCallbackUpdate {
	ccu.mutation.SetUpdatedAt(t)
	return ccu
}

// SetChatUserID sets the "chat_user_id" field.
func (ccu *ChatCallbackUpdate) SetChatUserID(s string) *ChatCallbackUpdate {
	ccu.mutation.SetChatUserID(s)
	return ccu
}

// SetRepoID sets the "repo_id" field.
func (ccu *ChatCallbackUpdate) SetRepoID(s string) *ChatCallbackUpdate {
	ccu.mutation.SetRepoID(s)
	return ccu
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (ccu *ChatCallbackUpdate) SetNillableRepoID(s *string) *ChatCallbackUpdate {
	if s != nil {
		ccu.SetRepoID(*s)
	}
	return ccu
}

// ClearRepoID clears the value of the "repo_id" field.
func (ccu *ChatCallbackUpdate) ClearRepoID() *ChatCallbackUpdate {
	ccu.mutation.ClearRepoID()
	return ccu
}

// SetChatUser sets the "chat_user" edge to the ChatUser entity.
func (ccu *ChatCallbackUpdate) SetChatUser(c *ChatUser) *ChatCallbackUpdate {
	return ccu.SetChatUserID(c.ID)
}

// SetRepo sets the "repo" edge to the Repo entity.
func (ccu *ChatCallbackUpdate) SetRepo(r *Repo) *ChatCallbackUpdate {
	return ccu.SetRepoID(r.ID)
}

// Mutation returns the ChatCallbackMutation object of the builder.
func (ccu *ChatCallbackUpdate) Mutation() *ChatCallbackMutation {
	return ccu.mutation
}

// ClearChatUser clears the "chat_user" edge to the ChatUser entity.
func (ccu *ChatCallbackUpdate) ClearChatUser() *ChatCallbackUpdate {
	ccu.mutation.ClearChatUser()
	return ccu
}

// ClearRepo clears the "repo" edge to the Repo entity.
func (ccu *ChatCallbackUpdate) ClearRepo() *ChatCallbackUpdate {
	ccu.mutation.ClearRepo()
	return ccu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *ChatCallbackUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ccu.defaults()
	if len(ccu.hooks) == 0 {
		if err = ccu.check(); err != nil {
			return 0, err
		}
		affected, err = ccu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatCallbackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccu.check(); err != nil {
				return 0, err
			}
			ccu.mutation = mutation
			affected, err = ccu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccu.hooks) - 1; i >= 0; i-- {
			mut = ccu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *ChatCallbackUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *ChatCallbackUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *ChatCallbackUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccu *ChatCallbackUpdate) defaults() {
	if _, ok := ccu.mutation.UpdatedAt(); !ok {
		v := chatcallback.UpdateDefaultUpdatedAt()
		ccu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccu *ChatCallbackUpdate) check() error {
	if v, ok := ccu.mutation.GetType(); ok {
		if err := chatcallback.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := ccu.mutation.ChatUserID(); ccu.mutation.ChatUserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"chat_user\"")
	}
	return nil
}

func (ccu *ChatCallbackUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chatcallback.Table,
			Columns: chatcallback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: chatcallback.FieldID,
			},
		},
	}
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatcallback.FieldState,
		})
	}
	if value, ok := ccu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: chatcallback.FieldType,
		})
	}
	if value, ok := ccu.mutation.IsOpened(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: chatcallback.FieldIsOpened,
		})
	}
	if value, ok := ccu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatcallback.FieldCreatedAt,
		})
	}
	if value, ok := ccu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatcallback.FieldUpdatedAt,
		})
	}
	if ccu.mutation.ChatUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatcallback.ChatUserTable,
			Columns: []string{chatcallback.ChatUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: chatuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.ChatUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatcallback.ChatUserTable,
			Columns: []string{chatcallback.ChatUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: chatuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ccu.mutation.RepoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatcallback.RepoTable,
			Columns: []string{chatcallback.RepoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccu.mutation.RepoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatcallback.RepoTable,
			Columns: []string{chatcallback.RepoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatcallback.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ChatCallbackUpdateOne is the builder for updating a single ChatCallback entity.
type ChatCallbackUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatCallbackMutation
}

// SetState sets the "state" field.
func (ccuo *ChatCallbackUpdateOne) SetState(s string) *ChatCallbackUpdateOne {
	ccuo.mutation.SetState(s)
	return ccuo
}

// SetType sets the "type" field.
func (ccuo *ChatCallbackUpdateOne) SetType(c chatcallback.Type) *ChatCallbackUpdateOne {
	ccuo.mutation.SetType(c)
	return ccuo
}

// SetIsOpened sets the "is_opened" field.
func (ccuo *ChatCallbackUpdateOne) SetIsOpened(b bool) *ChatCallbackUpdateOne {
	ccuo.mutation.SetIsOpened(b)
	return ccuo
}

// SetNillableIsOpened sets the "is_opened" field if the given value is not nil.
func (ccuo *ChatCallbackUpdateOne) SetNillableIsOpened(b *bool) *ChatCallbackUpdateOne {
	if b != nil {
		ccuo.SetIsOpened(*b)
	}
	return ccuo
}

// SetCreatedAt sets the "created_at" field.
func (ccuo *ChatCallbackUpdateOne) SetCreatedAt(t time.Time) *ChatCallbackUpdateOne {
	ccuo.mutation.SetCreatedAt(t)
	return ccuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccuo *ChatCallbackUpdateOne) SetNillableCreatedAt(t *time.Time) *ChatCallbackUpdateOne {
	if t != nil {
		ccuo.SetCreatedAt(*t)
	}
	return ccuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ccuo *ChatCallbackUpdateOne) SetUpdatedAt(t time.Time) *ChatCallbackUpdateOne {
	ccuo.mutation.SetUpdatedAt(t)
	return ccuo
}

// SetChatUserID sets the "chat_user_id" field.
func (ccuo *ChatCallbackUpdateOne) SetChatUserID(s string) *ChatCallbackUpdateOne {
	ccuo.mutation.SetChatUserID(s)
	return ccuo
}

// SetRepoID sets the "repo_id" field.
func (ccuo *ChatCallbackUpdateOne) SetRepoID(s string) *ChatCallbackUpdateOne {
	ccuo.mutation.SetRepoID(s)
	return ccuo
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (ccuo *ChatCallbackUpdateOne) SetNillableRepoID(s *string) *ChatCallbackUpdateOne {
	if s != nil {
		ccuo.SetRepoID(*s)
	}
	return ccuo
}

// ClearRepoID clears the value of the "repo_id" field.
func (ccuo *ChatCallbackUpdateOne) ClearRepoID() *ChatCallbackUpdateOne {
	ccuo.mutation.ClearRepoID()
	return ccuo
}

// SetChatUser sets the "chat_user" edge to the ChatUser entity.
func (ccuo *ChatCallbackUpdateOne) SetChatUser(c *ChatUser) *ChatCallbackUpdateOne {
	return ccuo.SetChatUserID(c.ID)
}

// SetRepo sets the "repo" edge to the Repo entity.
func (ccuo *ChatCallbackUpdateOne) SetRepo(r *Repo) *ChatCallbackUpdateOne {
	return ccuo.SetRepoID(r.ID)
}

// Mutation returns the ChatCallbackMutation object of the builder.
func (ccuo *ChatCallbackUpdateOne) Mutation() *ChatCallbackMutation {
	return ccuo.mutation
}

// ClearChatUser clears the "chat_user" edge to the ChatUser entity.
func (ccuo *ChatCallbackUpdateOne) ClearChatUser() *ChatCallbackUpdateOne {
	ccuo.mutation.ClearChatUser()
	return ccuo
}

// ClearRepo clears the "repo" edge to the Repo entity.
func (ccuo *ChatCallbackUpdateOne) ClearRepo() *ChatCallbackUpdateOne {
	ccuo.mutation.ClearRepo()
	return ccuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *ChatCallbackUpdateOne) Select(field string, fields ...string) *ChatCallbackUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated ChatCallback entity.
func (ccuo *ChatCallbackUpdateOne) Save(ctx context.Context) (*ChatCallback, error) {
	var (
		err  error
		node *ChatCallback
	)
	ccuo.defaults()
	if len(ccuo.hooks) == 0 {
		if err = ccuo.check(); err != nil {
			return nil, err
		}
		node, err = ccuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatCallbackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccuo.check(); err != nil {
				return nil, err
			}
			ccuo.mutation = mutation
			node, err = ccuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ccuo.hooks) - 1; i >= 0; i-- {
			mut = ccuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *ChatCallbackUpdateOne) SaveX(ctx context.Context) *ChatCallback {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *ChatCallbackUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *ChatCallbackUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccuo *ChatCallbackUpdateOne) defaults() {
	if _, ok := ccuo.mutation.UpdatedAt(); !ok {
		v := chatcallback.UpdateDefaultUpdatedAt()
		ccuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccuo *ChatCallbackUpdateOne) check() error {
	if v, ok := ccuo.mutation.GetType(); ok {
		if err := chatcallback.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := ccuo.mutation.ChatUserID(); ccuo.mutation.ChatUserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"chat_user\"")
	}
	return nil
}

func (ccuo *ChatCallbackUpdateOne) sqlSave(ctx context.Context) (_node *ChatCallback, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chatcallback.Table,
			Columns: chatcallback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: chatcallback.FieldID,
			},
		},
	}
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ChatCallback.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatcallback.FieldID)
		for _, f := range fields {
			if !chatcallback.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chatcallback.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatcallback.FieldState,
		})
	}
	if value, ok := ccuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: chatcallback.FieldType,
		})
	}
	if value, ok := ccuo.mutation.IsOpened(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: chatcallback.FieldIsOpened,
		})
	}
	if value, ok := ccuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatcallback.FieldCreatedAt,
		})
	}
	if value, ok := ccuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatcallback.FieldUpdatedAt,
		})
	}
	if ccuo.mutation.ChatUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatcallback.ChatUserTable,
			Columns: []string{chatcallback.ChatUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: chatuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.ChatUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatcallback.ChatUserTable,
			Columns: []string{chatcallback.ChatUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: chatuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ccuo.mutation.RepoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatcallback.RepoTable,
			Columns: []string{chatcallback.RepoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccuo.mutation.RepoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatcallback.RepoTable,
			Columns: []string{chatcallback.RepoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ChatCallback{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatcallback.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

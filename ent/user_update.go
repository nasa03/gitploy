// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hanjunlee/gitploy/ent/approval"
	"github.com/hanjunlee/gitploy/ent/chatuser"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/predicate"
	"github.com/hanjunlee/gitploy/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetLogin sets the "login" field.
func (uu *UserUpdate) SetLogin(s string) *UserUpdate {
	uu.mutation.SetLogin(s)
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetAdmin sets the "admin" field.
func (uu *UserUpdate) SetAdmin(b bool) *UserUpdate {
	uu.mutation.SetAdmin(b)
	return uu
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAdmin(b *bool) *UserUpdate {
	if b != nil {
		uu.SetAdmin(*b)
	}
	return uu
}

// SetToken sets the "token" field.
func (uu *UserUpdate) SetToken(s string) *UserUpdate {
	uu.mutation.SetToken(s)
	return uu
}

// SetRefresh sets the "refresh" field.
func (uu *UserUpdate) SetRefresh(s string) *UserUpdate {
	uu.mutation.SetRefresh(s)
	return uu
}

// SetExpiry sets the "expiry" field.
func (uu *UserUpdate) SetExpiry(t time.Time) *UserUpdate {
	uu.mutation.SetExpiry(t)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetChatUserID sets the "chat_user" edge to the ChatUser entity by ID.
func (uu *UserUpdate) SetChatUserID(id string) *UserUpdate {
	uu.mutation.SetChatUserID(id)
	return uu
}

// SetNillableChatUserID sets the "chat_user" edge to the ChatUser entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableChatUserID(id *string) *UserUpdate {
	if id != nil {
		uu = uu.SetChatUserID(*id)
	}
	return uu
}

// SetChatUser sets the "chat_user" edge to the ChatUser entity.
func (uu *UserUpdate) SetChatUser(c *ChatUser) *UserUpdate {
	return uu.SetChatUserID(c.ID)
}

// AddPermIDs adds the "perms" edge to the Perm entity by IDs.
func (uu *UserUpdate) AddPermIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPermIDs(ids...)
	return uu
}

// AddPerms adds the "perms" edges to the Perm entity.
func (uu *UserUpdate) AddPerms(p ...*Perm) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPermIDs(ids...)
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (uu *UserUpdate) AddDeploymentIDs(ids ...int) *UserUpdate {
	uu.mutation.AddDeploymentIDs(ids...)
	return uu
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (uu *UserUpdate) AddDeployments(d ...*Deployment) *UserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.AddDeploymentIDs(ids...)
}

// AddApprovalIDs adds the "approvals" edge to the Approval entity by IDs.
func (uu *UserUpdate) AddApprovalIDs(ids ...int) *UserUpdate {
	uu.mutation.AddApprovalIDs(ids...)
	return uu
}

// AddApprovals adds the "approvals" edges to the Approval entity.
func (uu *UserUpdate) AddApprovals(a ...*Approval) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddApprovalIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearChatUser clears the "chat_user" edge to the ChatUser entity.
func (uu *UserUpdate) ClearChatUser() *UserUpdate {
	uu.mutation.ClearChatUser()
	return uu
}

// ClearPerms clears all "perms" edges to the Perm entity.
func (uu *UserUpdate) ClearPerms() *UserUpdate {
	uu.mutation.ClearPerms()
	return uu
}

// RemovePermIDs removes the "perms" edge to Perm entities by IDs.
func (uu *UserUpdate) RemovePermIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePermIDs(ids...)
	return uu
}

// RemovePerms removes "perms" edges to Perm entities.
func (uu *UserUpdate) RemovePerms(p ...*Perm) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePermIDs(ids...)
}

// ClearDeployments clears all "deployments" edges to the Deployment entity.
func (uu *UserUpdate) ClearDeployments() *UserUpdate {
	uu.mutation.ClearDeployments()
	return uu
}

// RemoveDeploymentIDs removes the "deployments" edge to Deployment entities by IDs.
func (uu *UserUpdate) RemoveDeploymentIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveDeploymentIDs(ids...)
	return uu
}

// RemoveDeployments removes "deployments" edges to Deployment entities.
func (uu *UserUpdate) RemoveDeployments(d ...*Deployment) *UserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.RemoveDeploymentIDs(ids...)
}

// ClearApprovals clears all "approvals" edges to the Approval entity.
func (uu *UserUpdate) ClearApprovals() *UserUpdate {
	uu.mutation.ClearApprovals()
	return uu
}

// RemoveApprovalIDs removes the "approvals" edge to Approval entities by IDs.
func (uu *UserUpdate) RemoveApprovalIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveApprovalIDs(ids...)
	return uu
}

// RemoveApprovals removes "approvals" edges to Approval entities.
func (uu *UserUpdate) RemoveApprovals(a ...*Approval) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveApprovalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Login(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLogin,
		})
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatar,
		})
	}
	if uu.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatar,
		})
	}
	if value, ok := uu.mutation.Admin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldAdmin,
		})
	}
	if value, ok := uu.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldToken,
		})
	}
	if value, ok := uu.mutation.Refresh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRefresh,
		})
	}
	if value, ok := uu.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldExpiry,
		})
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uu.mutation.ChatUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ChatUserTable,
			Columns: []string{user.ChatUserColumn},
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
	if nodes := uu.mutation.ChatUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ChatUserTable,
			Columns: []string{user.ChatUserColumn},
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
	if uu.mutation.PermsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PermsTable,
			Columns: []string{user.PermsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: perm.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPermsIDs(); len(nodes) > 0 && !uu.mutation.PermsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PermsTable,
			Columns: []string{user.PermsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: perm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PermsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PermsTable,
			Columns: []string{user.PermsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: perm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.DeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeploymentsTable,
			Columns: []string{user.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deployment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedDeploymentsIDs(); len(nodes) > 0 && !uu.mutation.DeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeploymentsTable,
			Columns: []string{user.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeploymentsTable,
			Columns: []string{user.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ApprovalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ApprovalsTable,
			Columns: []string{user.ApprovalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: approval.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedApprovalsIDs(); len(nodes) > 0 && !uu.mutation.ApprovalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ApprovalsTable,
			Columns: []string{user.ApprovalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: approval.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ApprovalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ApprovalsTable,
			Columns: []string{user.ApprovalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: approval.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetLogin sets the "login" field.
func (uuo *UserUpdateOne) SetLogin(s string) *UserUpdateOne {
	uuo.mutation.SetLogin(s)
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetAdmin sets the "admin" field.
func (uuo *UserUpdateOne) SetAdmin(b bool) *UserUpdateOne {
	uuo.mutation.SetAdmin(b)
	return uuo
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAdmin(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetAdmin(*b)
	}
	return uuo
}

// SetToken sets the "token" field.
func (uuo *UserUpdateOne) SetToken(s string) *UserUpdateOne {
	uuo.mutation.SetToken(s)
	return uuo
}

// SetRefresh sets the "refresh" field.
func (uuo *UserUpdateOne) SetRefresh(s string) *UserUpdateOne {
	uuo.mutation.SetRefresh(s)
	return uuo
}

// SetExpiry sets the "expiry" field.
func (uuo *UserUpdateOne) SetExpiry(t time.Time) *UserUpdateOne {
	uuo.mutation.SetExpiry(t)
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetChatUserID sets the "chat_user" edge to the ChatUser entity by ID.
func (uuo *UserUpdateOne) SetChatUserID(id string) *UserUpdateOne {
	uuo.mutation.SetChatUserID(id)
	return uuo
}

// SetNillableChatUserID sets the "chat_user" edge to the ChatUser entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableChatUserID(id *string) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetChatUserID(*id)
	}
	return uuo
}

// SetChatUser sets the "chat_user" edge to the ChatUser entity.
func (uuo *UserUpdateOne) SetChatUser(c *ChatUser) *UserUpdateOne {
	return uuo.SetChatUserID(c.ID)
}

// AddPermIDs adds the "perms" edge to the Perm entity by IDs.
func (uuo *UserUpdateOne) AddPermIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPermIDs(ids...)
	return uuo
}

// AddPerms adds the "perms" edges to the Perm entity.
func (uuo *UserUpdateOne) AddPerms(p ...*Perm) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPermIDs(ids...)
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (uuo *UserUpdateOne) AddDeploymentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddDeploymentIDs(ids...)
	return uuo
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (uuo *UserUpdateOne) AddDeployments(d ...*Deployment) *UserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.AddDeploymentIDs(ids...)
}

// AddApprovalIDs adds the "approvals" edge to the Approval entity by IDs.
func (uuo *UserUpdateOne) AddApprovalIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddApprovalIDs(ids...)
	return uuo
}

// AddApprovals adds the "approvals" edges to the Approval entity.
func (uuo *UserUpdateOne) AddApprovals(a ...*Approval) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddApprovalIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearChatUser clears the "chat_user" edge to the ChatUser entity.
func (uuo *UserUpdateOne) ClearChatUser() *UserUpdateOne {
	uuo.mutation.ClearChatUser()
	return uuo
}

// ClearPerms clears all "perms" edges to the Perm entity.
func (uuo *UserUpdateOne) ClearPerms() *UserUpdateOne {
	uuo.mutation.ClearPerms()
	return uuo
}

// RemovePermIDs removes the "perms" edge to Perm entities by IDs.
func (uuo *UserUpdateOne) RemovePermIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePermIDs(ids...)
	return uuo
}

// RemovePerms removes "perms" edges to Perm entities.
func (uuo *UserUpdateOne) RemovePerms(p ...*Perm) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePermIDs(ids...)
}

// ClearDeployments clears all "deployments" edges to the Deployment entity.
func (uuo *UserUpdateOne) ClearDeployments() *UserUpdateOne {
	uuo.mutation.ClearDeployments()
	return uuo
}

// RemoveDeploymentIDs removes the "deployments" edge to Deployment entities by IDs.
func (uuo *UserUpdateOne) RemoveDeploymentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveDeploymentIDs(ids...)
	return uuo
}

// RemoveDeployments removes "deployments" edges to Deployment entities.
func (uuo *UserUpdateOne) RemoveDeployments(d ...*Deployment) *UserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.RemoveDeploymentIDs(ids...)
}

// ClearApprovals clears all "approvals" edges to the Approval entity.
func (uuo *UserUpdateOne) ClearApprovals() *UserUpdateOne {
	uuo.mutation.ClearApprovals()
	return uuo
}

// RemoveApprovalIDs removes the "approvals" edge to Approval entities by IDs.
func (uuo *UserUpdateOne) RemoveApprovalIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveApprovalIDs(ids...)
	return uuo
}

// RemoveApprovals removes "approvals" edges to Approval entities.
func (uuo *UserUpdateOne) RemoveApprovals(a ...*Approval) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveApprovalIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Login(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLogin,
		})
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatar,
		})
	}
	if uuo.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatar,
		})
	}
	if value, ok := uuo.mutation.Admin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldAdmin,
		})
	}
	if value, ok := uuo.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldToken,
		})
	}
	if value, ok := uuo.mutation.Refresh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRefresh,
		})
	}
	if value, ok := uuo.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldExpiry,
		})
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uuo.mutation.ChatUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ChatUserTable,
			Columns: []string{user.ChatUserColumn},
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
	if nodes := uuo.mutation.ChatUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ChatUserTable,
			Columns: []string{user.ChatUserColumn},
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
	if uuo.mutation.PermsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PermsTable,
			Columns: []string{user.PermsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: perm.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPermsIDs(); len(nodes) > 0 && !uuo.mutation.PermsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PermsTable,
			Columns: []string{user.PermsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: perm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PermsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PermsTable,
			Columns: []string{user.PermsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: perm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.DeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeploymentsTable,
			Columns: []string{user.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deployment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedDeploymentsIDs(); len(nodes) > 0 && !uuo.mutation.DeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeploymentsTable,
			Columns: []string{user.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeploymentsTable,
			Columns: []string{user.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ApprovalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ApprovalsTable,
			Columns: []string{user.ApprovalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: approval.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedApprovalsIDs(); len(nodes) > 0 && !uuo.mutation.ApprovalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ApprovalsTable,
			Columns: []string{user.ApprovalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: approval.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ApprovalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ApprovalsTable,
			Columns: []string{user.ApprovalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: approval.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

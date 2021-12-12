// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gitploy-io/gitploy/model/ent/chatuser"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
	"github.com/gitploy-io/gitploy/model/ent/user"
)

// ChatUserQuery is the builder for querying ChatUser entities.
type ChatUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ChatUser
	// eager-loading edges.
	withUser  *UserQuery
	modifiers []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChatUserQuery builder.
func (cuq *ChatUserQuery) Where(ps ...predicate.ChatUser) *ChatUserQuery {
	cuq.predicates = append(cuq.predicates, ps...)
	return cuq
}

// Limit adds a limit step to the query.
func (cuq *ChatUserQuery) Limit(limit int) *ChatUserQuery {
	cuq.limit = &limit
	return cuq
}

// Offset adds an offset step to the query.
func (cuq *ChatUserQuery) Offset(offset int) *ChatUserQuery {
	cuq.offset = &offset
	return cuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cuq *ChatUserQuery) Unique(unique bool) *ChatUserQuery {
	cuq.unique = &unique
	return cuq
}

// Order adds an order step to the query.
func (cuq *ChatUserQuery) Order(o ...OrderFunc) *ChatUserQuery {
	cuq.order = append(cuq.order, o...)
	return cuq
}

// QueryUser chains the current query on the "user" edge.
func (cuq *ChatUserQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: cuq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chatuser.Table, chatuser.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, chatuser.UserTable, chatuser.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ChatUser entity from the query.
// Returns a *NotFoundError when no ChatUser was found.
func (cuq *ChatUserQuery) First(ctx context.Context) (*ChatUser, error) {
	nodes, err := cuq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chatuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cuq *ChatUserQuery) FirstX(ctx context.Context) *ChatUser {
	node, err := cuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChatUser ID from the query.
// Returns a *NotFoundError when no ChatUser ID was found.
func (cuq *ChatUserQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cuq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chatuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cuq *ChatUserQuery) FirstIDX(ctx context.Context) string {
	id, err := cuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChatUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ChatUser entity is not found.
// Returns a *NotFoundError when no ChatUser entities are found.
func (cuq *ChatUserQuery) Only(ctx context.Context) (*ChatUser, error) {
	nodes, err := cuq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chatuser.Label}
	default:
		return nil, &NotSingularError{chatuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cuq *ChatUserQuery) OnlyX(ctx context.Context) *ChatUser {
	node, err := cuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChatUser ID in the query.
// Returns a *NotSingularError when exactly one ChatUser ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cuq *ChatUserQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cuq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chatuser.Label}
	default:
		err = &NotSingularError{chatuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cuq *ChatUserQuery) OnlyIDX(ctx context.Context) string {
	id, err := cuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChatUsers.
func (cuq *ChatUserQuery) All(ctx context.Context) ([]*ChatUser, error) {
	if err := cuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cuq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cuq *ChatUserQuery) AllX(ctx context.Context) []*ChatUser {
	nodes, err := cuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChatUser IDs.
func (cuq *ChatUserQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := cuq.Select(chatuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cuq *ChatUserQuery) IDsX(ctx context.Context) []string {
	ids, err := cuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cuq *ChatUserQuery) Count(ctx context.Context) (int, error) {
	if err := cuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cuq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cuq *ChatUserQuery) CountX(ctx context.Context) int {
	count, err := cuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cuq *ChatUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := cuq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cuq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cuq *ChatUserQuery) ExistX(ctx context.Context) bool {
	exist, err := cuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChatUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cuq *ChatUserQuery) Clone() *ChatUserQuery {
	if cuq == nil {
		return nil
	}
	return &ChatUserQuery{
		config:     cuq.config,
		limit:      cuq.limit,
		offset:     cuq.offset,
		order:      append([]OrderFunc{}, cuq.order...),
		predicates: append([]predicate.ChatUser{}, cuq.predicates...),
		withUser:   cuq.withUser.Clone(),
		// clone intermediate query.
		sql:  cuq.sql.Clone(),
		path: cuq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *ChatUserQuery) WithUser(opts ...func(*UserQuery)) *ChatUserQuery {
	query := &UserQuery{config: cuq.config}
	for _, opt := range opts {
		opt(query)
	}
	cuq.withUser = query
	return cuq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Token string `json:"token"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ChatUser.Query().
//		GroupBy(chatuser.FieldToken).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cuq *ChatUserQuery) GroupBy(field string, fields ...string) *ChatUserGroupBy {
	group := &ChatUserGroupBy{config: cuq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cuq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Token string `json:"token"`
//	}
//
//	client.ChatUser.Query().
//		Select(chatuser.FieldToken).
//		Scan(ctx, &v)
//
func (cuq *ChatUserQuery) Select(fields ...string) *ChatUserSelect {
	cuq.fields = append(cuq.fields, fields...)
	return &ChatUserSelect{ChatUserQuery: cuq}
}

func (cuq *ChatUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cuq.fields {
		if !chatuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cuq.path != nil {
		prev, err := cuq.path(ctx)
		if err != nil {
			return err
		}
		cuq.sql = prev
	}
	return nil
}

func (cuq *ChatUserQuery) sqlAll(ctx context.Context) ([]*ChatUser, error) {
	var (
		nodes       = []*ChatUser{}
		_spec       = cuq.querySpec()
		loadedTypes = [1]bool{
			cuq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ChatUser{config: cuq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cuq.modifiers) > 0 {
		_spec.Modifiers = cuq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, cuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cuq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*ChatUser)
		for i := range nodes {
			fk := nodes[i].UserID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (cuq *ChatUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cuq.querySpec()
	if len(cuq.modifiers) > 0 {
		_spec.Modifiers = cuq.modifiers
	}
	return sqlgraph.CountNodes(ctx, cuq.driver, _spec)
}

func (cuq *ChatUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cuq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cuq *ChatUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chatuser.Table,
			Columns: chatuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: chatuser.FieldID,
			},
		},
		From:   cuq.sql,
		Unique: true,
	}
	if unique := cuq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cuq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatuser.FieldID)
		for i := range fields {
			if fields[i] != chatuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cuq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cuq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cuq *ChatUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cuq.driver.Dialect())
	t1 := builder.Table(chatuser.Table)
	columns := cuq.fields
	if len(columns) == 0 {
		columns = chatuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cuq.sql != nil {
		selector = cuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, m := range cuq.modifiers {
		m(selector)
	}
	for _, p := range cuq.predicates {
		p(selector)
	}
	for _, p := range cuq.order {
		p(selector)
	}
	if offset := cuq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cuq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cuq *ChatUserQuery) ForUpdate(opts ...sql.LockOption) *ChatUserQuery {
	if cuq.driver.Dialect() == dialect.Postgres {
		cuq.Unique(false)
	}
	cuq.modifiers = append(cuq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cuq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cuq *ChatUserQuery) ForShare(opts ...sql.LockOption) *ChatUserQuery {
	if cuq.driver.Dialect() == dialect.Postgres {
		cuq.Unique(false)
	}
	cuq.modifiers = append(cuq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cuq
}

// ChatUserGroupBy is the group-by builder for ChatUser entities.
type ChatUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cugb *ChatUserGroupBy) Aggregate(fns ...AggregateFunc) *ChatUserGroupBy {
	cugb.fns = append(cugb.fns, fns...)
	return cugb
}

// Scan applies the group-by query and scans the result into the given value.
func (cugb *ChatUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cugb.path(ctx)
	if err != nil {
		return err
	}
	cugb.sql = query
	return cugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cugb *ChatUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cugb *ChatUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cugb.fields) > 1 {
		return nil, errors.New("ent: ChatUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cugb *ChatUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := cugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cugb *ChatUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatuser.Label}
	default:
		err = fmt.Errorf("ent: ChatUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cugb *ChatUserGroupBy) StringX(ctx context.Context) string {
	v, err := cugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cugb *ChatUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cugb.fields) > 1 {
		return nil, errors.New("ent: ChatUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cugb *ChatUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := cugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cugb *ChatUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatuser.Label}
	default:
		err = fmt.Errorf("ent: ChatUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cugb *ChatUserGroupBy) IntX(ctx context.Context) int {
	v, err := cugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cugb *ChatUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cugb.fields) > 1 {
		return nil, errors.New("ent: ChatUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cugb *ChatUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cugb *ChatUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatuser.Label}
	default:
		err = fmt.Errorf("ent: ChatUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cugb *ChatUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cugb *ChatUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cugb.fields) > 1 {
		return nil, errors.New("ent: ChatUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cugb *ChatUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cugb *ChatUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatuser.Label}
	default:
		err = fmt.Errorf("ent: ChatUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cugb *ChatUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := cugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cugb *ChatUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cugb.fields {
		if !chatuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cugb *ChatUserGroupBy) sqlQuery() *sql.Selector {
	selector := cugb.sql.Select()
	aggregation := make([]string, 0, len(cugb.fns))
	for _, fn := range cugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cugb.fields)+len(cugb.fns))
		for _, f := range cugb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cugb.fields...)...)
}

// ChatUserSelect is the builder for selecting fields of ChatUser entities.
type ChatUserSelect struct {
	*ChatUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cus *ChatUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cus.prepareQuery(ctx); err != nil {
		return err
	}
	cus.sql = cus.ChatUserQuery.sqlQuery(ctx)
	return cus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cus *ChatUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cus *ChatUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cus.fields) > 1 {
		return nil, errors.New("ent: ChatUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cus *ChatUserSelect) StringsX(ctx context.Context) []string {
	v, err := cus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cus *ChatUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatuser.Label}
	default:
		err = fmt.Errorf("ent: ChatUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cus *ChatUserSelect) StringX(ctx context.Context) string {
	v, err := cus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cus *ChatUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cus.fields) > 1 {
		return nil, errors.New("ent: ChatUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cus *ChatUserSelect) IntsX(ctx context.Context) []int {
	v, err := cus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cus *ChatUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatuser.Label}
	default:
		err = fmt.Errorf("ent: ChatUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cus *ChatUserSelect) IntX(ctx context.Context) int {
	v, err := cus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cus *ChatUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cus.fields) > 1 {
		return nil, errors.New("ent: ChatUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cus *ChatUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cus *ChatUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatuser.Label}
	default:
		err = fmt.Errorf("ent: ChatUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cus *ChatUserSelect) Float64X(ctx context.Context) float64 {
	v, err := cus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cus *ChatUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cus.fields) > 1 {
		return nil, errors.New("ent: ChatUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cus *ChatUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := cus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cus *ChatUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{chatuser.Label}
	default:
		err = fmt.Errorf("ent: ChatUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cus *ChatUserSelect) BoolX(ctx context.Context) bool {
	v, err := cus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cus *ChatUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cus.sql.Query()
	if err := cus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
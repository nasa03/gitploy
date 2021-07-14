// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/deploymentstatus"
	"github.com/hanjunlee/gitploy/ent/predicate"
)

// DeploymentStatusQuery is the builder for querying DeploymentStatus entities.
type DeploymentStatusQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DeploymentStatus
	// eager-loading edges.
	withDeployment *DeploymentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeploymentStatusQuery builder.
func (dsq *DeploymentStatusQuery) Where(ps ...predicate.DeploymentStatus) *DeploymentStatusQuery {
	dsq.predicates = append(dsq.predicates, ps...)
	return dsq
}

// Limit adds a limit step to the query.
func (dsq *DeploymentStatusQuery) Limit(limit int) *DeploymentStatusQuery {
	dsq.limit = &limit
	return dsq
}

// Offset adds an offset step to the query.
func (dsq *DeploymentStatusQuery) Offset(offset int) *DeploymentStatusQuery {
	dsq.offset = &offset
	return dsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dsq *DeploymentStatusQuery) Unique(unique bool) *DeploymentStatusQuery {
	dsq.unique = &unique
	return dsq
}

// Order adds an order step to the query.
func (dsq *DeploymentStatusQuery) Order(o ...OrderFunc) *DeploymentStatusQuery {
	dsq.order = append(dsq.order, o...)
	return dsq
}

// QueryDeployment chains the current query on the "deployment" edge.
func (dsq *DeploymentStatusQuery) QueryDeployment() *DeploymentQuery {
	query := &DeploymentQuery{config: dsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deploymentstatus.Table, deploymentstatus.FieldID, selector),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deploymentstatus.DeploymentTable, deploymentstatus.DeploymentColumn),
		)
		fromU = sqlgraph.SetNeighbors(dsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DeploymentStatus entity from the query.
// Returns a *NotFoundError when no DeploymentStatus was found.
func (dsq *DeploymentStatusQuery) First(ctx context.Context) (*DeploymentStatus, error) {
	nodes, err := dsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deploymentstatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dsq *DeploymentStatusQuery) FirstX(ctx context.Context) *DeploymentStatus {
	node, err := dsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeploymentStatus ID from the query.
// Returns a *NotFoundError when no DeploymentStatus ID was found.
func (dsq *DeploymentStatusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deploymentstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dsq *DeploymentStatusQuery) FirstIDX(ctx context.Context) int {
	id, err := dsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeploymentStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DeploymentStatus entity is not found.
// Returns a *NotFoundError when no DeploymentStatus entities are found.
func (dsq *DeploymentStatusQuery) Only(ctx context.Context) (*DeploymentStatus, error) {
	nodes, err := dsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deploymentstatus.Label}
	default:
		return nil, &NotSingularError{deploymentstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dsq *DeploymentStatusQuery) OnlyX(ctx context.Context) *DeploymentStatus {
	node, err := dsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeploymentStatus ID in the query.
// Returns a *NotSingularError when exactly one DeploymentStatus ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dsq *DeploymentStatusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deploymentstatus.Label}
	default:
		err = &NotSingularError{deploymentstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dsq *DeploymentStatusQuery) OnlyIDX(ctx context.Context) int {
	id, err := dsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeploymentStatusSlice.
func (dsq *DeploymentStatusQuery) All(ctx context.Context) ([]*DeploymentStatus, error) {
	if err := dsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dsq *DeploymentStatusQuery) AllX(ctx context.Context) []*DeploymentStatus {
	nodes, err := dsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeploymentStatus IDs.
func (dsq *DeploymentStatusQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dsq.Select(deploymentstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dsq *DeploymentStatusQuery) IDsX(ctx context.Context) []int {
	ids, err := dsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dsq *DeploymentStatusQuery) Count(ctx context.Context) (int, error) {
	if err := dsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dsq *DeploymentStatusQuery) CountX(ctx context.Context) int {
	count, err := dsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dsq *DeploymentStatusQuery) Exist(ctx context.Context) (bool, error) {
	if err := dsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dsq *DeploymentStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := dsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeploymentStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dsq *DeploymentStatusQuery) Clone() *DeploymentStatusQuery {
	if dsq == nil {
		return nil
	}
	return &DeploymentStatusQuery{
		config:         dsq.config,
		limit:          dsq.limit,
		offset:         dsq.offset,
		order:          append([]OrderFunc{}, dsq.order...),
		predicates:     append([]predicate.DeploymentStatus{}, dsq.predicates...),
		withDeployment: dsq.withDeployment.Clone(),
		// clone intermediate query.
		sql:  dsq.sql.Clone(),
		path: dsq.path,
	}
}

// WithDeployment tells the query-builder to eager-load the nodes that are connected to
// the "deployment" edge. The optional arguments are used to configure the query builder of the edge.
func (dsq *DeploymentStatusQuery) WithDeployment(opts ...func(*DeploymentQuery)) *DeploymentStatusQuery {
	query := &DeploymentQuery{config: dsq.config}
	for _, opt := range opts {
		opt(query)
	}
	dsq.withDeployment = query
	return dsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeploymentStatus.Query().
//		GroupBy(deploymentstatus.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dsq *DeploymentStatusQuery) GroupBy(field string, fields ...string) *DeploymentStatusGroupBy {
	group := &DeploymentStatusGroupBy{config: dsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status"`
//	}
//
//	client.DeploymentStatus.Query().
//		Select(deploymentstatus.FieldStatus).
//		Scan(ctx, &v)
//
func (dsq *DeploymentStatusQuery) Select(field string, fields ...string) *DeploymentStatusSelect {
	dsq.fields = append([]string{field}, fields...)
	return &DeploymentStatusSelect{DeploymentStatusQuery: dsq}
}

func (dsq *DeploymentStatusQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dsq.fields {
		if !deploymentstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dsq.path != nil {
		prev, err := dsq.path(ctx)
		if err != nil {
			return err
		}
		dsq.sql = prev
	}
	return nil
}

func (dsq *DeploymentStatusQuery) sqlAll(ctx context.Context) ([]*DeploymentStatus, error) {
	var (
		nodes       = []*DeploymentStatus{}
		_spec       = dsq.querySpec()
		loadedTypes = [1]bool{
			dsq.withDeployment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DeploymentStatus{config: dsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, dsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dsq.withDeployment; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DeploymentStatus)
		for i := range nodes {
			fk := nodes[i].DeploymentID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(deployment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "deployment_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Deployment = n
			}
		}
	}

	return nodes, nil
}

func (dsq *DeploymentStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dsq.querySpec()
	return sqlgraph.CountNodes(ctx, dsq.driver, _spec)
}

func (dsq *DeploymentStatusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dsq *DeploymentStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deploymentstatus.Table,
			Columns: deploymentstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deploymentstatus.FieldID,
			},
		},
		From:   dsq.sql,
		Unique: true,
	}
	if unique := dsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deploymentstatus.FieldID)
		for i := range fields {
			if fields[i] != deploymentstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dsq *DeploymentStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dsq.driver.Dialect())
	t1 := builder.Table(deploymentstatus.Table)
	selector := builder.Select(t1.Columns(deploymentstatus.Columns...)...).From(t1)
	if dsq.sql != nil {
		selector = dsq.sql
		selector.Select(selector.Columns(deploymentstatus.Columns...)...)
	}
	for _, p := range dsq.predicates {
		p(selector)
	}
	for _, p := range dsq.order {
		p(selector)
	}
	if offset := dsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeploymentStatusGroupBy is the group-by builder for DeploymentStatus entities.
type DeploymentStatusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dsgb *DeploymentStatusGroupBy) Aggregate(fns ...AggregateFunc) *DeploymentStatusGroupBy {
	dsgb.fns = append(dsgb.fns, fns...)
	return dsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dsgb *DeploymentStatusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dsgb.path(ctx)
	if err != nil {
		return err
	}
	dsgb.sql = query
	return dsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dsgb *DeploymentStatusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dsgb *DeploymentStatusGroupBy) StringsX(ctx context.Context) []string {
	v, err := dsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatus.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dsgb *DeploymentStatusGroupBy) StringX(ctx context.Context) string {
	v, err := dsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dsgb *DeploymentStatusGroupBy) IntsX(ctx context.Context) []int {
	v, err := dsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatus.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dsgb *DeploymentStatusGroupBy) IntX(ctx context.Context) int {
	v, err := dsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dsgb *DeploymentStatusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatus.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dsgb *DeploymentStatusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dsgb *DeploymentStatusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DeploymentStatusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatus.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dsgb *DeploymentStatusGroupBy) BoolX(ctx context.Context) bool {
	v, err := dsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dsgb *DeploymentStatusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dsgb.fields {
		if !deploymentstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dsgb *DeploymentStatusGroupBy) sqlQuery() *sql.Selector {
	selector := dsgb.sql
	columns := make([]string, 0, len(dsgb.fields)+len(dsgb.fns))
	columns = append(columns, dsgb.fields...)
	for _, fn := range dsgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dsgb.fields...)
}

// DeploymentStatusSelect is the builder for selecting fields of DeploymentStatus entities.
type DeploymentStatusSelect struct {
	*DeploymentStatusQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dss *DeploymentStatusSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dss.prepareQuery(ctx); err != nil {
		return err
	}
	dss.sql = dss.DeploymentStatusQuery.sqlQuery(ctx)
	return dss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dss *DeploymentStatusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dss *DeploymentStatusSelect) StringsX(ctx context.Context) []string {
	v, err := dss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatus.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dss *DeploymentStatusSelect) StringX(ctx context.Context) string {
	v, err := dss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dss *DeploymentStatusSelect) IntsX(ctx context.Context) []int {
	v, err := dss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatus.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dss *DeploymentStatusSelect) IntX(ctx context.Context) int {
	v, err := dss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dss *DeploymentStatusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatus.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dss *DeploymentStatusSelect) Float64X(ctx context.Context) float64 {
	v, err := dss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DeploymentStatusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dss *DeploymentStatusSelect) BoolsX(ctx context.Context) []bool {
	v, err := dss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dss *DeploymentStatusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deploymentstatus.Label}
	default:
		err = fmt.Errorf("ent: DeploymentStatusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dss *DeploymentStatusSelect) BoolX(ctx context.Context) bool {
	v, err := dss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dss *DeploymentStatusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dss.sqlQuery().Query()
	if err := dss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dss *DeploymentStatusSelect) sqlQuery() sql.Querier {
	selector := dss.sql
	selector.Select(selector.Columns(dss.fields...)...)
	return selector
}

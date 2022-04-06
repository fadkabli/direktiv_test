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
	"github.com/direktiv/direktiv/pkg/secrets/ent/namespacesecret"
	"github.com/direktiv/direktiv/pkg/secrets/ent/predicate"
)

// NamespaceSecretQuery is the builder for querying NamespaceSecret entities.
type NamespaceSecretQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NamespaceSecret
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NamespaceSecretQuery builder.
func (nsq *NamespaceSecretQuery) Where(ps ...predicate.NamespaceSecret) *NamespaceSecretQuery {
	nsq.predicates = append(nsq.predicates, ps...)
	return nsq
}

// Limit adds a limit step to the query.
func (nsq *NamespaceSecretQuery) Limit(limit int) *NamespaceSecretQuery {
	nsq.limit = &limit
	return nsq
}

// Offset adds an offset step to the query.
func (nsq *NamespaceSecretQuery) Offset(offset int) *NamespaceSecretQuery {
	nsq.offset = &offset
	return nsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nsq *NamespaceSecretQuery) Unique(unique bool) *NamespaceSecretQuery {
	nsq.unique = &unique
	return nsq
}

// Order adds an order step to the query.
func (nsq *NamespaceSecretQuery) Order(o ...OrderFunc) *NamespaceSecretQuery {
	nsq.order = append(nsq.order, o...)
	return nsq
}

// First returns the first NamespaceSecret entity from the query.
// Returns a *NotFoundError when no NamespaceSecret was found.
func (nsq *NamespaceSecretQuery) First(ctx context.Context) (*NamespaceSecret, error) {
	nodes, err := nsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{namespacesecret.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nsq *NamespaceSecretQuery) FirstX(ctx context.Context) *NamespaceSecret {
	node, err := nsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NamespaceSecret ID from the query.
// Returns a *NotFoundError when no NamespaceSecret ID was found.
func (nsq *NamespaceSecretQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{namespacesecret.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nsq *NamespaceSecretQuery) FirstIDX(ctx context.Context) int {
	id, err := nsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NamespaceSecret entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NamespaceSecret entity is found.
// Returns a *NotFoundError when no NamespaceSecret entities are found.
func (nsq *NamespaceSecretQuery) Only(ctx context.Context) (*NamespaceSecret, error) {
	nodes, err := nsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{namespacesecret.Label}
	default:
		return nil, &NotSingularError{namespacesecret.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nsq *NamespaceSecretQuery) OnlyX(ctx context.Context) *NamespaceSecret {
	node, err := nsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NamespaceSecret ID in the query.
// Returns a *NotSingularError when more than one NamespaceSecret ID is found.
// Returns a *NotFoundError when no entities are found.
func (nsq *NamespaceSecretQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{namespacesecret.Label}
	default:
		err = &NotSingularError{namespacesecret.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nsq *NamespaceSecretQuery) OnlyIDX(ctx context.Context) int {
	id, err := nsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NamespaceSecrets.
func (nsq *NamespaceSecretQuery) All(ctx context.Context) ([]*NamespaceSecret, error) {
	if err := nsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nsq *NamespaceSecretQuery) AllX(ctx context.Context) []*NamespaceSecret {
	nodes, err := nsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NamespaceSecret IDs.
func (nsq *NamespaceSecretQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := nsq.Select(namespacesecret.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nsq *NamespaceSecretQuery) IDsX(ctx context.Context) []int {
	ids, err := nsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nsq *NamespaceSecretQuery) Count(ctx context.Context) (int, error) {
	if err := nsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nsq *NamespaceSecretQuery) CountX(ctx context.Context) int {
	count, err := nsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nsq *NamespaceSecretQuery) Exist(ctx context.Context) (bool, error) {
	if err := nsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nsq *NamespaceSecretQuery) ExistX(ctx context.Context) bool {
	exist, err := nsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NamespaceSecretQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nsq *NamespaceSecretQuery) Clone() *NamespaceSecretQuery {
	if nsq == nil {
		return nil
	}
	return &NamespaceSecretQuery{
		config:     nsq.config,
		limit:      nsq.limit,
		offset:     nsq.offset,
		order:      append([]OrderFunc{}, nsq.order...),
		predicates: append([]predicate.NamespaceSecret{}, nsq.predicates...),
		// clone intermediate query.
		sql:    nsq.sql.Clone(),
		path:   nsq.path,
		unique: nsq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Ns string `json:"ns,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NamespaceSecret.Query().
//		GroupBy(namespacesecret.FieldNs).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nsq *NamespaceSecretQuery) GroupBy(field string, fields ...string) *NamespaceSecretGroupBy {
	group := &NamespaceSecretGroupBy{config: nsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Ns string `json:"ns,omitempty"`
//	}
//
//	client.NamespaceSecret.Query().
//		Select(namespacesecret.FieldNs).
//		Scan(ctx, &v)
//
func (nsq *NamespaceSecretQuery) Select(fields ...string) *NamespaceSecretSelect {
	nsq.fields = append(nsq.fields, fields...)
	return &NamespaceSecretSelect{NamespaceSecretQuery: nsq}
}

func (nsq *NamespaceSecretQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nsq.fields {
		if !namespacesecret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nsq.path != nil {
		prev, err := nsq.path(ctx)
		if err != nil {
			return err
		}
		nsq.sql = prev
	}
	return nil
}

func (nsq *NamespaceSecretQuery) sqlAll(ctx context.Context) ([]*NamespaceSecret, error) {
	var (
		nodes = []*NamespaceSecret{}
		_spec = nsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NamespaceSecret{config: nsq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, nsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (nsq *NamespaceSecretQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nsq.querySpec()
	_spec.Node.Columns = nsq.fields
	if len(nsq.fields) > 0 {
		_spec.Unique = nsq.unique != nil && *nsq.unique
	}
	return sqlgraph.CountNodes(ctx, nsq.driver, _spec)
}

func (nsq *NamespaceSecretQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nsq *NamespaceSecretQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   namespacesecret.Table,
			Columns: namespacesecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: namespacesecret.FieldID,
			},
		},
		From:   nsq.sql,
		Unique: true,
	}
	if unique := nsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, namespacesecret.FieldID)
		for i := range fields {
			if fields[i] != namespacesecret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nsq *NamespaceSecretQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nsq.driver.Dialect())
	t1 := builder.Table(namespacesecret.Table)
	columns := nsq.fields
	if len(columns) == 0 {
		columns = namespacesecret.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nsq.sql != nil {
		selector = nsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nsq.unique != nil && *nsq.unique {
		selector.Distinct()
	}
	for _, p := range nsq.predicates {
		p(selector)
	}
	for _, p := range nsq.order {
		p(selector)
	}
	if offset := nsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NamespaceSecretGroupBy is the group-by builder for NamespaceSecret entities.
type NamespaceSecretGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nsgb *NamespaceSecretGroupBy) Aggregate(fns ...AggregateFunc) *NamespaceSecretGroupBy {
	nsgb.fns = append(nsgb.fns, fns...)
	return nsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (nsgb *NamespaceSecretGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := nsgb.path(ctx)
	if err != nil {
		return err
	}
	nsgb.sql = query
	return nsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nsgb *NamespaceSecretGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := nsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (nsgb *NamespaceSecretGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(nsgb.fields) > 1 {
		return nil, errors.New("ent: NamespaceSecretGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := nsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nsgb *NamespaceSecretGroupBy) StringsX(ctx context.Context) []string {
	v, err := nsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nsgb *NamespaceSecretGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{namespacesecret.Label}
	default:
		err = fmt.Errorf("ent: NamespaceSecretGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nsgb *NamespaceSecretGroupBy) StringX(ctx context.Context) string {
	v, err := nsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (nsgb *NamespaceSecretGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(nsgb.fields) > 1 {
		return nil, errors.New("ent: NamespaceSecretGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := nsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nsgb *NamespaceSecretGroupBy) IntsX(ctx context.Context) []int {
	v, err := nsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nsgb *NamespaceSecretGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{namespacesecret.Label}
	default:
		err = fmt.Errorf("ent: NamespaceSecretGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nsgb *NamespaceSecretGroupBy) IntX(ctx context.Context) int {
	v, err := nsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (nsgb *NamespaceSecretGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(nsgb.fields) > 1 {
		return nil, errors.New("ent: NamespaceSecretGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := nsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nsgb *NamespaceSecretGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := nsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nsgb *NamespaceSecretGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{namespacesecret.Label}
	default:
		err = fmt.Errorf("ent: NamespaceSecretGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nsgb *NamespaceSecretGroupBy) Float64X(ctx context.Context) float64 {
	v, err := nsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (nsgb *NamespaceSecretGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(nsgb.fields) > 1 {
		return nil, errors.New("ent: NamespaceSecretGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := nsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nsgb *NamespaceSecretGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := nsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nsgb *NamespaceSecretGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{namespacesecret.Label}
	default:
		err = fmt.Errorf("ent: NamespaceSecretGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nsgb *NamespaceSecretGroupBy) BoolX(ctx context.Context) bool {
	v, err := nsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nsgb *NamespaceSecretGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range nsgb.fields {
		if !namespacesecret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := nsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (nsgb *NamespaceSecretGroupBy) sqlQuery() *sql.Selector {
	selector := nsgb.sql.Select()
	aggregation := make([]string, 0, len(nsgb.fns))
	for _, fn := range nsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(nsgb.fields)+len(nsgb.fns))
		for _, f := range nsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(nsgb.fields...)...)
}

// NamespaceSecretSelect is the builder for selecting fields of NamespaceSecret entities.
type NamespaceSecretSelect struct {
	*NamespaceSecretQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nss *NamespaceSecretSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nss.prepareQuery(ctx); err != nil {
		return err
	}
	nss.sql = nss.NamespaceSecretQuery.sqlQuery(ctx)
	return nss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nss *NamespaceSecretSelect) ScanX(ctx context.Context, v interface{}) {
	if err := nss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (nss *NamespaceSecretSelect) Strings(ctx context.Context) ([]string, error) {
	if len(nss.fields) > 1 {
		return nil, errors.New("ent: NamespaceSecretSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := nss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nss *NamespaceSecretSelect) StringsX(ctx context.Context) []string {
	v, err := nss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (nss *NamespaceSecretSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{namespacesecret.Label}
	default:
		err = fmt.Errorf("ent: NamespaceSecretSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nss *NamespaceSecretSelect) StringX(ctx context.Context) string {
	v, err := nss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (nss *NamespaceSecretSelect) Ints(ctx context.Context) ([]int, error) {
	if len(nss.fields) > 1 {
		return nil, errors.New("ent: NamespaceSecretSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := nss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nss *NamespaceSecretSelect) IntsX(ctx context.Context) []int {
	v, err := nss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (nss *NamespaceSecretSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{namespacesecret.Label}
	default:
		err = fmt.Errorf("ent: NamespaceSecretSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nss *NamespaceSecretSelect) IntX(ctx context.Context) int {
	v, err := nss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (nss *NamespaceSecretSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(nss.fields) > 1 {
		return nil, errors.New("ent: NamespaceSecretSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := nss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nss *NamespaceSecretSelect) Float64sX(ctx context.Context) []float64 {
	v, err := nss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (nss *NamespaceSecretSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{namespacesecret.Label}
	default:
		err = fmt.Errorf("ent: NamespaceSecretSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nss *NamespaceSecretSelect) Float64X(ctx context.Context) float64 {
	v, err := nss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (nss *NamespaceSecretSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(nss.fields) > 1 {
		return nil, errors.New("ent: NamespaceSecretSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := nss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nss *NamespaceSecretSelect) BoolsX(ctx context.Context) []bool {
	v, err := nss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (nss *NamespaceSecretSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{namespacesecret.Label}
	default:
		err = fmt.Errorf("ent: NamespaceSecretSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nss *NamespaceSecretSelect) BoolX(ctx context.Context) bool {
	v, err := nss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nss *NamespaceSecretSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nss.sql.Query()
	if err := nss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

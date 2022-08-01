// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirroractivity"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// LogMsgQuery is the builder for querying LogMsg entities.
type LogMsgQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LogMsg
	// eager-loading edges.
	withNamespace *NamespaceQuery
	withWorkflow  *WorkflowQuery
	withInstance  *InstanceQuery
	withActivity  *MirrorActivityQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LogMsgQuery builder.
func (lmq *LogMsgQuery) Where(ps ...predicate.LogMsg) *LogMsgQuery {
	lmq.predicates = append(lmq.predicates, ps...)
	return lmq
}

// Limit adds a limit step to the query.
func (lmq *LogMsgQuery) Limit(limit int) *LogMsgQuery {
	lmq.limit = &limit
	return lmq
}

// Offset adds an offset step to the query.
func (lmq *LogMsgQuery) Offset(offset int) *LogMsgQuery {
	lmq.offset = &offset
	return lmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lmq *LogMsgQuery) Unique(unique bool) *LogMsgQuery {
	lmq.unique = &unique
	return lmq
}

// Order adds an order step to the query.
func (lmq *LogMsgQuery) Order(o ...OrderFunc) *LogMsgQuery {
	lmq.order = append(lmq.order, o...)
	return lmq
}

// QueryNamespace chains the current query on the "namespace" edge.
func (lmq *LogMsgQuery) QueryNamespace() *NamespaceQuery {
	query := &NamespaceQuery{config: lmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(logmsg.Table, logmsg.FieldID, selector),
			sqlgraph.To(namespace.Table, namespace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, logmsg.NamespaceTable, logmsg.NamespaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(lmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkflow chains the current query on the "workflow" edge.
func (lmq *LogMsgQuery) QueryWorkflow() *WorkflowQuery {
	query := &WorkflowQuery{config: lmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(logmsg.Table, logmsg.FieldID, selector),
			sqlgraph.To(workflow.Table, workflow.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, logmsg.WorkflowTable, logmsg.WorkflowColumn),
		)
		fromU = sqlgraph.SetNeighbors(lmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInstance chains the current query on the "instance" edge.
func (lmq *LogMsgQuery) QueryInstance() *InstanceQuery {
	query := &InstanceQuery{config: lmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(logmsg.Table, logmsg.FieldID, selector),
			sqlgraph.To(instance.Table, instance.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, logmsg.InstanceTable, logmsg.InstanceColumn),
		)
		fromU = sqlgraph.SetNeighbors(lmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryActivity chains the current query on the "activity" edge.
func (lmq *LogMsgQuery) QueryActivity() *MirrorActivityQuery {
	query := &MirrorActivityQuery{config: lmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(logmsg.Table, logmsg.FieldID, selector),
			sqlgraph.To(mirroractivity.Table, mirroractivity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, logmsg.ActivityTable, logmsg.ActivityColumn),
		)
		fromU = sqlgraph.SetNeighbors(lmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LogMsg entity from the query.
// Returns a *NotFoundError when no LogMsg was found.
func (lmq *LogMsgQuery) First(ctx context.Context) (*LogMsg, error) {
	nodes, err := lmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{logmsg.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lmq *LogMsgQuery) FirstX(ctx context.Context) *LogMsg {
	node, err := lmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LogMsg ID from the query.
// Returns a *NotFoundError when no LogMsg ID was found.
func (lmq *LogMsgQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{logmsg.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lmq *LogMsgQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LogMsg entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LogMsg entity is found.
// Returns a *NotFoundError when no LogMsg entities are found.
func (lmq *LogMsgQuery) Only(ctx context.Context) (*LogMsg, error) {
	nodes, err := lmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{logmsg.Label}
	default:
		return nil, &NotSingularError{logmsg.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lmq *LogMsgQuery) OnlyX(ctx context.Context) *LogMsg {
	node, err := lmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LogMsg ID in the query.
// Returns a *NotSingularError when more than one LogMsg ID is found.
// Returns a *NotFoundError when no entities are found.
func (lmq *LogMsgQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{logmsg.Label}
	default:
		err = &NotSingularError{logmsg.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lmq *LogMsgQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LogMsgs.
func (lmq *LogMsgQuery) All(ctx context.Context) ([]*LogMsg, error) {
	if err := lmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lmq *LogMsgQuery) AllX(ctx context.Context) []*LogMsg {
	nodes, err := lmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LogMsg IDs.
func (lmq *LogMsgQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := lmq.Select(logmsg.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lmq *LogMsgQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lmq *LogMsgQuery) Count(ctx context.Context) (int, error) {
	if err := lmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lmq *LogMsgQuery) CountX(ctx context.Context) int {
	count, err := lmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lmq *LogMsgQuery) Exist(ctx context.Context) (bool, error) {
	if err := lmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lmq *LogMsgQuery) ExistX(ctx context.Context) bool {
	exist, err := lmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LogMsgQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lmq *LogMsgQuery) Clone() *LogMsgQuery {
	if lmq == nil {
		return nil
	}
	return &LogMsgQuery{
		config:        lmq.config,
		limit:         lmq.limit,
		offset:        lmq.offset,
		order:         append([]OrderFunc{}, lmq.order...),
		predicates:    append([]predicate.LogMsg{}, lmq.predicates...),
		withNamespace: lmq.withNamespace.Clone(),
		withWorkflow:  lmq.withWorkflow.Clone(),
		withInstance:  lmq.withInstance.Clone(),
		withActivity:  lmq.withActivity.Clone(),
		// clone intermediate query.
		sql:    lmq.sql.Clone(),
		path:   lmq.path,
		unique: lmq.unique,
	}
}

// WithNamespace tells the query-builder to eager-load the nodes that are connected to
// the "namespace" edge. The optional arguments are used to configure the query builder of the edge.
func (lmq *LogMsgQuery) WithNamespace(opts ...func(*NamespaceQuery)) *LogMsgQuery {
	query := &NamespaceQuery{config: lmq.config}
	for _, opt := range opts {
		opt(query)
	}
	lmq.withNamespace = query
	return lmq
}

// WithWorkflow tells the query-builder to eager-load the nodes that are connected to
// the "workflow" edge. The optional arguments are used to configure the query builder of the edge.
func (lmq *LogMsgQuery) WithWorkflow(opts ...func(*WorkflowQuery)) *LogMsgQuery {
	query := &WorkflowQuery{config: lmq.config}
	for _, opt := range opts {
		opt(query)
	}
	lmq.withWorkflow = query
	return lmq
}

// WithInstance tells the query-builder to eager-load the nodes that are connected to
// the "instance" edge. The optional arguments are used to configure the query builder of the edge.
func (lmq *LogMsgQuery) WithInstance(opts ...func(*InstanceQuery)) *LogMsgQuery {
	query := &InstanceQuery{config: lmq.config}
	for _, opt := range opts {
		opt(query)
	}
	lmq.withInstance = query
	return lmq
}

// WithActivity tells the query-builder to eager-load the nodes that are connected to
// the "activity" edge. The optional arguments are used to configure the query builder of the edge.
func (lmq *LogMsgQuery) WithActivity(opts ...func(*MirrorActivityQuery)) *LogMsgQuery {
	query := &MirrorActivityQuery{config: lmq.config}
	for _, opt := range opts {
		opt(query)
	}
	lmq.withActivity = query
	return lmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		T time.Time `json:"t,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LogMsg.Query().
//		GroupBy(logmsg.FieldT).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (lmq *LogMsgQuery) GroupBy(field string, fields ...string) *LogMsgGroupBy {
	grbuild := &LogMsgGroupBy{config: lmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lmq.sqlQuery(ctx), nil
	}
	grbuild.label = logmsg.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		T time.Time `json:"t,omitempty"`
//	}
//
//	client.LogMsg.Query().
//		Select(logmsg.FieldT).
//		Scan(ctx, &v)
//
func (lmq *LogMsgQuery) Select(fields ...string) *LogMsgSelect {
	lmq.fields = append(lmq.fields, fields...)
	selbuild := &LogMsgSelect{LogMsgQuery: lmq}
	selbuild.label = logmsg.Label
	selbuild.flds, selbuild.scan = &lmq.fields, selbuild.Scan
	return selbuild
}

func (lmq *LogMsgQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lmq.fields {
		if !logmsg.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lmq.path != nil {
		prev, err := lmq.path(ctx)
		if err != nil {
			return err
		}
		lmq.sql = prev
	}
	return nil
}

func (lmq *LogMsgQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LogMsg, error) {
	var (
		nodes       = []*LogMsg{}
		withFKs     = lmq.withFKs
		_spec       = lmq.querySpec()
		loadedTypes = [4]bool{
			lmq.withNamespace != nil,
			lmq.withWorkflow != nil,
			lmq.withInstance != nil,
			lmq.withActivity != nil,
		}
	)
	if lmq.withNamespace != nil || lmq.withWorkflow != nil || lmq.withInstance != nil || lmq.withActivity != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, logmsg.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*LogMsg).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &LogMsg{config: lmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := lmq.withNamespace; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*LogMsg)
		for i := range nodes {
			if nodes[i].namespace_logs == nil {
				continue
			}
			fk := *nodes[i].namespace_logs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(namespace.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "namespace_logs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Namespace = n
			}
		}
	}

	if query := lmq.withWorkflow; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*LogMsg)
		for i := range nodes {
			if nodes[i].workflow_logs == nil {
				continue
			}
			fk := *nodes[i].workflow_logs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(workflow.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workflow_logs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Workflow = n
			}
		}
	}

	if query := lmq.withInstance; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*LogMsg)
		for i := range nodes {
			if nodes[i].instance_logs == nil {
				continue
			}
			fk := *nodes[i].instance_logs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(instance.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "instance_logs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Instance = n
			}
		}
	}

	if query := lmq.withActivity; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*LogMsg)
		for i := range nodes {
			if nodes[i].mirror_activity_logs == nil {
				continue
			}
			fk := *nodes[i].mirror_activity_logs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(mirroractivity.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "mirror_activity_logs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Activity = n
			}
		}
	}

	return nodes, nil
}

func (lmq *LogMsgQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lmq.querySpec()
	_spec.Node.Columns = lmq.fields
	if len(lmq.fields) > 0 {
		_spec.Unique = lmq.unique != nil && *lmq.unique
	}
	return sqlgraph.CountNodes(ctx, lmq.driver, _spec)
}

func (lmq *LogMsgQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (lmq *LogMsgQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   logmsg.Table,
			Columns: logmsg.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: logmsg.FieldID,
			},
		},
		From:   lmq.sql,
		Unique: true,
	}
	if unique := lmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logmsg.FieldID)
		for i := range fields {
			if fields[i] != logmsg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lmq *LogMsgQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lmq.driver.Dialect())
	t1 := builder.Table(logmsg.Table)
	columns := lmq.fields
	if len(columns) == 0 {
		columns = logmsg.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lmq.sql != nil {
		selector = lmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lmq.unique != nil && *lmq.unique {
		selector.Distinct()
	}
	for _, p := range lmq.predicates {
		p(selector)
	}
	for _, p := range lmq.order {
		p(selector)
	}
	if offset := lmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LogMsgGroupBy is the group-by builder for LogMsg entities.
type LogMsgGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lmgb *LogMsgGroupBy) Aggregate(fns ...AggregateFunc) *LogMsgGroupBy {
	lmgb.fns = append(lmgb.fns, fns...)
	return lmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (lmgb *LogMsgGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lmgb.path(ctx)
	if err != nil {
		return err
	}
	lmgb.sql = query
	return lmgb.sqlScan(ctx, v)
}

func (lmgb *LogMsgGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lmgb.fields {
		if !logmsg.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lmgb *LogMsgGroupBy) sqlQuery() *sql.Selector {
	selector := lmgb.sql.Select()
	aggregation := make([]string, 0, len(lmgb.fns))
	for _, fn := range lmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lmgb.fields)+len(lmgb.fns))
		for _, f := range lmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lmgb.fields...)...)
}

// LogMsgSelect is the builder for selecting fields of LogMsg entities.
type LogMsgSelect struct {
	*LogMsgQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (lms *LogMsgSelect) Scan(ctx context.Context, v interface{}) error {
	if err := lms.prepareQuery(ctx); err != nil {
		return err
	}
	lms.sql = lms.LogMsgQuery.sqlQuery(ctx)
	return lms.sqlScan(ctx, v)
}

func (lms *LogMsgSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := lms.sql.Query()
	if err := lms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

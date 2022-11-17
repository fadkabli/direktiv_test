// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/services"
	"github.com/google/uuid"
)

// ServicesUpdate is the builder for updating Services entities.
type ServicesUpdate struct {
	config
	hooks    []Hook
	mutation *ServicesMutation
}

// Where appends a list predicates to the ServicesUpdate builder.
func (su *ServicesUpdate) Where(ps ...predicate.Services) *ServicesUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *ServicesUpdate) SetName(s string) *ServicesUpdate {
	su.mutation.SetName(s)
	return su
}

// SetData sets the "data" field.
func (su *ServicesUpdate) SetData(s string) *ServicesUpdate {
	su.mutation.SetData(s)
	return su
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (su *ServicesUpdate) SetNamespaceID(id uuid.UUID) *ServicesUpdate {
	su.mutation.SetNamespaceID(id)
	return su
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (su *ServicesUpdate) SetNamespace(n *Namespace) *ServicesUpdate {
	return su.SetNamespaceID(n.ID)
}

// Mutation returns the ServicesMutation object of the builder.
func (su *ServicesUpdate) Mutation() *ServicesMutation {
	return su.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (su *ServicesUpdate) ClearNamespace() *ServicesUpdate {
	su.mutation.ClearNamespace()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ServicesUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServicesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ServicesUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ServicesUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ServicesUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ServicesUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := services.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Services.name": %w`, err)}
		}
	}
	if _, ok := su.mutation.NamespaceID(); su.mutation.NamespaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Services.namespace"`)
	}
	return nil
}

func (su *ServicesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   services.Table,
			Columns: services.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: services.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(services.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Data(); ok {
		_spec.SetField(services.FieldData, field.TypeString, value)
	}
	if su.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   services.NamespaceTable,
			Columns: []string{services.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   services.NamespaceTable,
			Columns: []string{services.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{services.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ServicesUpdateOne is the builder for updating a single Services entity.
type ServicesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServicesMutation
}

// SetName sets the "name" field.
func (suo *ServicesUpdateOne) SetName(s string) *ServicesUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetData sets the "data" field.
func (suo *ServicesUpdateOne) SetData(s string) *ServicesUpdateOne {
	suo.mutation.SetData(s)
	return suo
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (suo *ServicesUpdateOne) SetNamespaceID(id uuid.UUID) *ServicesUpdateOne {
	suo.mutation.SetNamespaceID(id)
	return suo
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (suo *ServicesUpdateOne) SetNamespace(n *Namespace) *ServicesUpdateOne {
	return suo.SetNamespaceID(n.ID)
}

// Mutation returns the ServicesMutation object of the builder.
func (suo *ServicesUpdateOne) Mutation() *ServicesMutation {
	return suo.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (suo *ServicesUpdateOne) ClearNamespace() *ServicesUpdateOne {
	suo.mutation.ClearNamespace()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ServicesUpdateOne) Select(field string, fields ...string) *ServicesUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Services entity.
func (suo *ServicesUpdateOne) Save(ctx context.Context) (*Services, error) {
	var (
		err  error
		node *Services
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServicesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Services)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ServicesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ServicesUpdateOne) SaveX(ctx context.Context) *Services {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ServicesUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ServicesUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ServicesUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := services.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Services.name": %w`, err)}
		}
	}
	if _, ok := suo.mutation.NamespaceID(); suo.mutation.NamespaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Services.namespace"`)
	}
	return nil
}

func (suo *ServicesUpdateOne) sqlSave(ctx context.Context) (_node *Services, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   services.Table,
			Columns: services.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: services.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Services.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, services.FieldID)
		for _, f := range fields {
			if !services.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != services.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(services.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Data(); ok {
		_spec.SetField(services.FieldData, field.TypeString, value)
	}
	if suo.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   services.NamespaceTable,
			Columns: []string{services.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   services.NamespaceTable,
			Columns: []string{services.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Services{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{services.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

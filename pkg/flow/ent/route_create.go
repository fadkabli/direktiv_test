// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/ref"
	"github.com/direktiv/direktiv/pkg/flow/ent/route"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// RouteCreate is the builder for creating a Route entity.
type RouteCreate struct {
	config
	mutation *RouteMutation
	hooks    []Hook
}

// SetWeight sets the "weight" field.
func (rc *RouteCreate) SetWeight(i int) *RouteCreate {
	rc.mutation.SetWeight(i)
	return rc
}

// SetID sets the "id" field.
func (rc *RouteCreate) SetID(u uuid.UUID) *RouteCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RouteCreate) SetNillableID(u *uuid.UUID) *RouteCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (rc *RouteCreate) SetWorkflowID(id uuid.UUID) *RouteCreate {
	rc.mutation.SetWorkflowID(id)
	return rc
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (rc *RouteCreate) SetWorkflow(w *Workflow) *RouteCreate {
	return rc.SetWorkflowID(w.ID)
}

// SetRefID sets the "ref" edge to the Ref entity by ID.
func (rc *RouteCreate) SetRefID(id uuid.UUID) *RouteCreate {
	rc.mutation.SetRefID(id)
	return rc
}

// SetRef sets the "ref" edge to the Ref entity.
func (rc *RouteCreate) SetRef(r *Ref) *RouteCreate {
	return rc.SetRefID(r.ID)
}

// Mutation returns the RouteMutation object of the builder.
func (rc *RouteCreate) Mutation() *RouteMutation {
	return rc.mutation
}

// Save creates the Route in the database.
func (rc *RouteCreate) Save(ctx context.Context) (*Route, error) {
	var (
		err  error
		node *Route
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RouteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Route)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RouteMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RouteCreate) SaveX(ctx context.Context) *Route {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RouteCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RouteCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RouteCreate) defaults() {
	if _, ok := rc.mutation.ID(); !ok {
		v := route.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RouteCreate) check() error {
	if _, ok := rc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Route.weight"`)}
	}
	if _, ok := rc.mutation.WorkflowID(); !ok {
		return &ValidationError{Name: "workflow", err: errors.New(`ent: missing required edge "Route.workflow"`)}
	}
	if _, ok := rc.mutation.RefID(); !ok {
		return &ValidationError{Name: "ref", err: errors.New(`ent: missing required edge "Route.ref"`)}
	}
	return nil
}

func (rc *RouteCreate) sqlSave(ctx context.Context) (*Route, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (rc *RouteCreate) createSpec() (*Route, *sqlgraph.CreateSpec) {
	var (
		_node = &Route{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: route.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: route.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.Weight(); ok {
		_spec.SetField(route.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if nodes := rc.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.WorkflowTable,
			Columns: []string{route.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workflow_routes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.RefTable,
			Columns: []string{route.RefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ref.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ref_routes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RouteCreateBulk is the builder for creating many Route entities in bulk.
type RouteCreateBulk struct {
	config
	builders []*RouteCreate
}

// Save creates the Route entities in the database.
func (rcb *RouteCreateBulk) Save(ctx context.Context) ([]*Route, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Route, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RouteMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RouteCreateBulk) SaveX(ctx context.Context) []*Route {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RouteCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RouteCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

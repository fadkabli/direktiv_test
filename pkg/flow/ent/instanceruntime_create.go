// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/instanceruntime"
)

// InstanceRuntimeCreate is the builder for creating a InstanceRuntime entity.
type InstanceRuntimeCreate struct {
	config
	mutation *InstanceRuntimeMutation
	hooks    []Hook
}

// SetInput sets the "input" field.
func (irc *InstanceRuntimeCreate) SetInput(b []byte) *InstanceRuntimeCreate {
	irc.mutation.SetInput(b)
	return irc
}

// SetData sets the "data" field.
func (irc *InstanceRuntimeCreate) SetData(s string) *InstanceRuntimeCreate {
	irc.mutation.SetData(s)
	return irc
}

// SetController sets the "controller" field.
func (irc *InstanceRuntimeCreate) SetController(s string) *InstanceRuntimeCreate {
	irc.mutation.SetController(s)
	return irc
}

// SetNillableController sets the "controller" field if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableController(s *string) *InstanceRuntimeCreate {
	if s != nil {
		irc.SetController(*s)
	}
	return irc
}

// SetMemory sets the "memory" field.
func (irc *InstanceRuntimeCreate) SetMemory(s string) *InstanceRuntimeCreate {
	irc.mutation.SetMemory(s)
	return irc
}

// SetNillableMemory sets the "memory" field if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableMemory(s *string) *InstanceRuntimeCreate {
	if s != nil {
		irc.SetMemory(*s)
	}
	return irc
}

// SetFlow sets the "flow" field.
func (irc *InstanceRuntimeCreate) SetFlow(s []string) *InstanceRuntimeCreate {
	irc.mutation.SetFlow(s)
	return irc
}

// SetOutput sets the "output" field.
func (irc *InstanceRuntimeCreate) SetOutput(s string) *InstanceRuntimeCreate {
	irc.mutation.SetOutput(s)
	return irc
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableOutput(s *string) *InstanceRuntimeCreate {
	if s != nil {
		irc.SetOutput(*s)
	}
	return irc
}

// SetStateBeginTime sets the "stateBeginTime" field.
func (irc *InstanceRuntimeCreate) SetStateBeginTime(t time.Time) *InstanceRuntimeCreate {
	irc.mutation.SetStateBeginTime(t)
	return irc
}

// SetNillableStateBeginTime sets the "stateBeginTime" field if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableStateBeginTime(t *time.Time) *InstanceRuntimeCreate {
	if t != nil {
		irc.SetStateBeginTime(*t)
	}
	return irc
}

// SetDeadline sets the "deadline" field.
func (irc *InstanceRuntimeCreate) SetDeadline(t time.Time) *InstanceRuntimeCreate {
	irc.mutation.SetDeadline(t)
	return irc
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableDeadline(t *time.Time) *InstanceRuntimeCreate {
	if t != nil {
		irc.SetDeadline(*t)
	}
	return irc
}

// SetAttempts sets the "attempts" field.
func (irc *InstanceRuntimeCreate) SetAttempts(i int) *InstanceRuntimeCreate {
	irc.mutation.SetAttempts(i)
	return irc
}

// SetNillableAttempts sets the "attempts" field if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableAttempts(i *int) *InstanceRuntimeCreate {
	if i != nil {
		irc.SetAttempts(*i)
	}
	return irc
}

// SetCallerData sets the "caller_data" field.
func (irc *InstanceRuntimeCreate) SetCallerData(s string) *InstanceRuntimeCreate {
	irc.mutation.SetCallerData(s)
	return irc
}

// SetNillableCallerData sets the "caller_data" field if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableCallerData(s *string) *InstanceRuntimeCreate {
	if s != nil {
		irc.SetCallerData(*s)
	}
	return irc
}

// SetInstanceContext sets the "instanceContext" field.
func (irc *InstanceRuntimeCreate) SetInstanceContext(s string) *InstanceRuntimeCreate {
	irc.mutation.SetInstanceContext(s)
	return irc
}

// SetNillableInstanceContext sets the "instanceContext" field if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableInstanceContext(s *string) *InstanceRuntimeCreate {
	if s != nil {
		irc.SetInstanceContext(*s)
	}
	return irc
}

// SetStateContext sets the "stateContext" field.
func (irc *InstanceRuntimeCreate) SetStateContext(s string) *InstanceRuntimeCreate {
	irc.mutation.SetStateContext(s)
	return irc
}

// SetNillableStateContext sets the "stateContext" field if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableStateContext(s *string) *InstanceRuntimeCreate {
	if s != nil {
		irc.SetStateContext(*s)
	}
	return irc
}

// SetID sets the "id" field.
func (irc *InstanceRuntimeCreate) SetID(u uuid.UUID) *InstanceRuntimeCreate {
	irc.mutation.SetID(u)
	return irc
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (irc *InstanceRuntimeCreate) SetInstanceID(id uuid.UUID) *InstanceRuntimeCreate {
	irc.mutation.SetInstanceID(id)
	return irc
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableInstanceID(id *uuid.UUID) *InstanceRuntimeCreate {
	if id != nil {
		irc = irc.SetInstanceID(*id)
	}
	return irc
}

// SetInstance sets the "instance" edge to the Instance entity.
func (irc *InstanceRuntimeCreate) SetInstance(i *Instance) *InstanceRuntimeCreate {
	return irc.SetInstanceID(i.ID)
}

// SetCallerID sets the "caller" edge to the Instance entity by ID.
func (irc *InstanceRuntimeCreate) SetCallerID(id uuid.UUID) *InstanceRuntimeCreate {
	irc.mutation.SetCallerID(id)
	return irc
}

// SetNillableCallerID sets the "caller" edge to the Instance entity by ID if the given value is not nil.
func (irc *InstanceRuntimeCreate) SetNillableCallerID(id *uuid.UUID) *InstanceRuntimeCreate {
	if id != nil {
		irc = irc.SetCallerID(*id)
	}
	return irc
}

// SetCaller sets the "caller" edge to the Instance entity.
func (irc *InstanceRuntimeCreate) SetCaller(i *Instance) *InstanceRuntimeCreate {
	return irc.SetCallerID(i.ID)
}

// Mutation returns the InstanceRuntimeMutation object of the builder.
func (irc *InstanceRuntimeCreate) Mutation() *InstanceRuntimeMutation {
	return irc.mutation
}

// Save creates the InstanceRuntime in the database.
func (irc *InstanceRuntimeCreate) Save(ctx context.Context) (*InstanceRuntime, error) {
	var (
		err  error
		node *InstanceRuntime
	)
	irc.defaults()
	if len(irc.hooks) == 0 {
		if err = irc.check(); err != nil {
			return nil, err
		}
		node, err = irc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceRuntimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = irc.check(); err != nil {
				return nil, err
			}
			irc.mutation = mutation
			if node, err = irc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(irc.hooks) - 1; i >= 0; i-- {
			if irc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = irc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, irc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (irc *InstanceRuntimeCreate) SaveX(ctx context.Context) *InstanceRuntime {
	v, err := irc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (irc *InstanceRuntimeCreate) Exec(ctx context.Context) error {
	_, err := irc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (irc *InstanceRuntimeCreate) ExecX(ctx context.Context) {
	if err := irc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (irc *InstanceRuntimeCreate) defaults() {
	if _, ok := irc.mutation.ID(); !ok {
		v := instanceruntime.DefaultID()
		irc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (irc *InstanceRuntimeCreate) check() error {
	if _, ok := irc.mutation.Input(); !ok {
		return &ValidationError{Name: "input", err: errors.New(`ent: missing required field "input"`)}
	}
	if _, ok := irc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "data"`)}
	}
	return nil
}

func (irc *InstanceRuntimeCreate) sqlSave(ctx context.Context) (*InstanceRuntime, error) {
	_node, _spec := irc.createSpec()
	if err := sqlgraph.CreateNode(ctx, irc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (irc *InstanceRuntimeCreate) createSpec() (*InstanceRuntime, *sqlgraph.CreateSpec) {
	var (
		_node = &InstanceRuntime{config: irc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: instanceruntime.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: instanceruntime.FieldID,
			},
		}
	)
	if id, ok := irc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := irc.mutation.Input(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: instanceruntime.FieldInput,
		})
		_node.Input = value
	}
	if value, ok := irc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldData,
		})
		_node.Data = value
	}
	if value, ok := irc.mutation.Controller(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldController,
		})
		_node.Controller = value
	}
	if value, ok := irc.mutation.Memory(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldMemory,
		})
		_node.Memory = value
	}
	if value, ok := irc.mutation.Flow(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: instanceruntime.FieldFlow,
		})
		_node.Flow = value
	}
	if value, ok := irc.mutation.Output(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldOutput,
		})
		_node.Output = value
	}
	if value, ok := irc.mutation.StateBeginTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instanceruntime.FieldStateBeginTime,
		})
		_node.StateBeginTime = value
	}
	if value, ok := irc.mutation.Deadline(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instanceruntime.FieldDeadline,
		})
		_node.Deadline = value
	}
	if value, ok := irc.mutation.Attempts(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instanceruntime.FieldAttempts,
		})
		_node.Attempts = value
	}
	if value, ok := irc.mutation.CallerData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldCallerData,
		})
		_node.CallerData = value
	}
	if value, ok := irc.mutation.InstanceContext(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldInstanceContext,
		})
		_node.InstanceContext = value
	}
	if value, ok := irc.mutation.StateContext(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldStateContext,
		})
		_node.StateContext = value
	}
	if nodes := irc.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   instanceruntime.InstanceTable,
			Columns: []string{instanceruntime.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.instance_runtime = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := irc.mutation.CallerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instanceruntime.CallerTable,
			Columns: []string{instanceruntime.CallerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.instance_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InstanceRuntimeCreateBulk is the builder for creating many InstanceRuntime entities in bulk.
type InstanceRuntimeCreateBulk struct {
	config
	builders []*InstanceRuntimeCreate
}

// Save creates the InstanceRuntime entities in the database.
func (ircb *InstanceRuntimeCreateBulk) Save(ctx context.Context) ([]*InstanceRuntime, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ircb.builders))
	nodes := make([]*InstanceRuntime, len(ircb.builders))
	mutators := make([]Mutator, len(ircb.builders))
	for i := range ircb.builders {
		func(i int, root context.Context) {
			builder := ircb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InstanceRuntimeMutation)
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
					_, err = mutators[i+1].Mutate(root, ircb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ircb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ircb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ircb *InstanceRuntimeCreateBulk) SaveX(ctx context.Context) []*InstanceRuntime {
	v, err := ircb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ircb *InstanceRuntimeCreateBulk) Exec(ctx context.Context) error {
	_, err := ircb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ircb *InstanceRuntimeCreateBulk) ExecX(ctx context.Context) {
	if err := ircb.Exec(ctx); err != nil {
		panic(err)
	}
}

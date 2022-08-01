// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/eventswait"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// EventsWaitUpdate is the builder for updating EventsWait entities.
type EventsWaitUpdate struct {
	config
	hooks    []Hook
	mutation *EventsWaitMutation
}

// Where appends a list predicates to the EventsWaitUpdate builder.
func (ewu *EventsWaitUpdate) Where(ps ...predicate.EventsWait) *EventsWaitUpdate {
	ewu.mutation.Where(ps...)
	return ewu
}

// SetEvents sets the "events" field.
func (ewu *EventsWaitUpdate) SetEvents(m map[string]interface{}) *EventsWaitUpdate {
	ewu.mutation.SetEvents(m)
	return ewu
}

// SetWorkfloweventID sets the "workflowevent" edge to the Events entity by ID.
func (ewu *EventsWaitUpdate) SetWorkfloweventID(id uuid.UUID) *EventsWaitUpdate {
	ewu.mutation.SetWorkfloweventID(id)
	return ewu
}

// SetWorkflowevent sets the "workflowevent" edge to the Events entity.
func (ewu *EventsWaitUpdate) SetWorkflowevent(e *Events) *EventsWaitUpdate {
	return ewu.SetWorkfloweventID(e.ID)
}

// Mutation returns the EventsWaitMutation object of the builder.
func (ewu *EventsWaitUpdate) Mutation() *EventsWaitMutation {
	return ewu.mutation
}

// ClearWorkflowevent clears the "workflowevent" edge to the Events entity.
func (ewu *EventsWaitUpdate) ClearWorkflowevent() *EventsWaitUpdate {
	ewu.mutation.ClearWorkflowevent()
	return ewu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ewu *EventsWaitUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ewu.hooks) == 0 {
		if err = ewu.check(); err != nil {
			return 0, err
		}
		affected, err = ewu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventsWaitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ewu.check(); err != nil {
				return 0, err
			}
			ewu.mutation = mutation
			affected, err = ewu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ewu.hooks) - 1; i >= 0; i-- {
			if ewu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ewu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ewu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ewu *EventsWaitUpdate) SaveX(ctx context.Context) int {
	affected, err := ewu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ewu *EventsWaitUpdate) Exec(ctx context.Context) error {
	_, err := ewu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ewu *EventsWaitUpdate) ExecX(ctx context.Context) {
	if err := ewu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ewu *EventsWaitUpdate) check() error {
	if _, ok := ewu.mutation.WorkfloweventID(); ewu.mutation.WorkfloweventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EventsWait.workflowevent"`)
	}
	return nil
}

func (ewu *EventsWaitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventswait.Table,
			Columns: eventswait.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: eventswait.FieldID,
			},
		},
	}
	if ps := ewu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ewu.mutation.Events(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: eventswait.FieldEvents,
		})
	}
	if ewu.mutation.WorkfloweventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventswait.WorkfloweventTable,
			Columns: []string{eventswait.WorkfloweventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: events.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ewu.mutation.WorkfloweventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventswait.WorkfloweventTable,
			Columns: []string{eventswait.WorkfloweventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: events.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ewu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventswait.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EventsWaitUpdateOne is the builder for updating a single EventsWait entity.
type EventsWaitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventsWaitMutation
}

// SetEvents sets the "events" field.
func (ewuo *EventsWaitUpdateOne) SetEvents(m map[string]interface{}) *EventsWaitUpdateOne {
	ewuo.mutation.SetEvents(m)
	return ewuo
}

// SetWorkfloweventID sets the "workflowevent" edge to the Events entity by ID.
func (ewuo *EventsWaitUpdateOne) SetWorkfloweventID(id uuid.UUID) *EventsWaitUpdateOne {
	ewuo.mutation.SetWorkfloweventID(id)
	return ewuo
}

// SetWorkflowevent sets the "workflowevent" edge to the Events entity.
func (ewuo *EventsWaitUpdateOne) SetWorkflowevent(e *Events) *EventsWaitUpdateOne {
	return ewuo.SetWorkfloweventID(e.ID)
}

// Mutation returns the EventsWaitMutation object of the builder.
func (ewuo *EventsWaitUpdateOne) Mutation() *EventsWaitMutation {
	return ewuo.mutation
}

// ClearWorkflowevent clears the "workflowevent" edge to the Events entity.
func (ewuo *EventsWaitUpdateOne) ClearWorkflowevent() *EventsWaitUpdateOne {
	ewuo.mutation.ClearWorkflowevent()
	return ewuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ewuo *EventsWaitUpdateOne) Select(field string, fields ...string) *EventsWaitUpdateOne {
	ewuo.fields = append([]string{field}, fields...)
	return ewuo
}

// Save executes the query and returns the updated EventsWait entity.
func (ewuo *EventsWaitUpdateOne) Save(ctx context.Context) (*EventsWait, error) {
	var (
		err  error
		node *EventsWait
	)
	if len(ewuo.hooks) == 0 {
		if err = ewuo.check(); err != nil {
			return nil, err
		}
		node, err = ewuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventsWaitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ewuo.check(); err != nil {
				return nil, err
			}
			ewuo.mutation = mutation
			node, err = ewuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ewuo.hooks) - 1; i >= 0; i-- {
			if ewuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ewuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ewuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EventsWait)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EventsWaitMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ewuo *EventsWaitUpdateOne) SaveX(ctx context.Context) *EventsWait {
	node, err := ewuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ewuo *EventsWaitUpdateOne) Exec(ctx context.Context) error {
	_, err := ewuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ewuo *EventsWaitUpdateOne) ExecX(ctx context.Context) {
	if err := ewuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ewuo *EventsWaitUpdateOne) check() error {
	if _, ok := ewuo.mutation.WorkfloweventID(); ewuo.mutation.WorkfloweventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EventsWait.workflowevent"`)
	}
	return nil
}

func (ewuo *EventsWaitUpdateOne) sqlSave(ctx context.Context) (_node *EventsWait, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventswait.Table,
			Columns: eventswait.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: eventswait.FieldID,
			},
		},
	}
	id, ok := ewuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EventsWait.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ewuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventswait.FieldID)
		for _, f := range fields {
			if !eventswait.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != eventswait.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ewuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ewuo.mutation.Events(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: eventswait.FieldEvents,
		})
	}
	if ewuo.mutation.WorkfloweventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventswait.WorkfloweventTable,
			Columns: []string{eventswait.WorkfloweventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: events.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ewuo.mutation.WorkfloweventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventswait.WorkfloweventTable,
			Columns: []string{eventswait.WorkfloweventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: events.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EventsWait{config: ewuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ewuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventswait.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

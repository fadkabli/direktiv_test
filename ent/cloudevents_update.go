// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/direktiv/direktiv/ent/cloudevents"
	"github.com/direktiv/direktiv/ent/predicate"
)

// CloudEventsUpdate is the builder for updating CloudEvents entities.
type CloudEventsUpdate struct {
	config
	hooks    []Hook
	mutation *CloudEventsMutation
}

// Where appends a list predicates to the CloudEventsUpdate builder.
func (ceu *CloudEventsUpdate) Where(ps ...predicate.CloudEvents) *CloudEventsUpdate {
	ceu.mutation.Where(ps...)
	return ceu
}

// SetEvent sets the "event" field.
func (ceu *CloudEventsUpdate) SetEvent(e event.Event) *CloudEventsUpdate {
	ceu.mutation.SetEvent(e)
	return ceu
}

// SetProcessed sets the "processed" field.
func (ceu *CloudEventsUpdate) SetProcessed(b bool) *CloudEventsUpdate {
	ceu.mutation.SetProcessed(b)
	return ceu
}

// Mutation returns the CloudEventsMutation object of the builder.
func (ceu *CloudEventsUpdate) Mutation() *CloudEventsMutation {
	return ceu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ceu *CloudEventsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ceu.hooks) == 0 {
		affected, err = ceu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CloudEventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ceu.mutation = mutation
			affected, err = ceu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ceu.hooks) - 1; i >= 0; i-- {
			if ceu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ceu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceu *CloudEventsUpdate) SaveX(ctx context.Context) int {
	affected, err := ceu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ceu *CloudEventsUpdate) Exec(ctx context.Context) error {
	_, err := ceu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceu *CloudEventsUpdate) ExecX(ctx context.Context) {
	if err := ceu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ceu *CloudEventsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cloudevents.Table,
			Columns: cloudevents.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: cloudevents.FieldID,
			},
		},
	}
	if ps := ceu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceu.mutation.Event(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: cloudevents.FieldEvent,
		})
	}
	if value, ok := ceu.mutation.Processed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cloudevents.FieldProcessed,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ceu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cloudevents.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CloudEventsUpdateOne is the builder for updating a single CloudEvents entity.
type CloudEventsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CloudEventsMutation
}

// SetEvent sets the "event" field.
func (ceuo *CloudEventsUpdateOne) SetEvent(e event.Event) *CloudEventsUpdateOne {
	ceuo.mutation.SetEvent(e)
	return ceuo
}

// SetProcessed sets the "processed" field.
func (ceuo *CloudEventsUpdateOne) SetProcessed(b bool) *CloudEventsUpdateOne {
	ceuo.mutation.SetProcessed(b)
	return ceuo
}

// Mutation returns the CloudEventsMutation object of the builder.
func (ceuo *CloudEventsUpdateOne) Mutation() *CloudEventsMutation {
	return ceuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ceuo *CloudEventsUpdateOne) Select(field string, fields ...string) *CloudEventsUpdateOne {
	ceuo.fields = append([]string{field}, fields...)
	return ceuo
}

// Save executes the query and returns the updated CloudEvents entity.
func (ceuo *CloudEventsUpdateOne) Save(ctx context.Context) (*CloudEvents, error) {
	var (
		err  error
		node *CloudEvents
	)
	if len(ceuo.hooks) == 0 {
		node, err = ceuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CloudEventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ceuo.mutation = mutation
			node, err = ceuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ceuo.hooks) - 1; i >= 0; i-- {
			if ceuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ceuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceuo *CloudEventsUpdateOne) SaveX(ctx context.Context) *CloudEvents {
	node, err := ceuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ceuo *CloudEventsUpdateOne) Exec(ctx context.Context) error {
	_, err := ceuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceuo *CloudEventsUpdateOne) ExecX(ctx context.Context) {
	if err := ceuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ceuo *CloudEventsUpdateOne) sqlSave(ctx context.Context) (_node *CloudEvents, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cloudevents.Table,
			Columns: cloudevents.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: cloudevents.FieldID,
			},
		},
	}
	id, ok := ceuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CloudEvents.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ceuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cloudevents.FieldID)
		for _, f := range fields {
			if !cloudevents.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cloudevents.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ceuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceuo.mutation.Event(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: cloudevents.FieldEvent,
		})
	}
	if value, ok := ceuo.mutation.Processed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cloudevents.FieldProcessed,
		})
	}
	_node = &CloudEvents{config: ceuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ceuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cloudevents.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

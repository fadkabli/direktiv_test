// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudeventfilters"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
)

// CloudEventFiltersDelete is the builder for deleting a CloudEventFilters entity.
type CloudEventFiltersDelete struct {
	config
	hooks    []Hook
	mutation *CloudEventFiltersMutation
}

// Where appends a list predicates to the CloudEventFiltersDelete builder.
func (cefd *CloudEventFiltersDelete) Where(ps ...predicate.CloudEventFilters) *CloudEventFiltersDelete {
	cefd.mutation.Where(ps...)
	return cefd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cefd *CloudEventFiltersDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cefd.hooks) == 0 {
		affected, err = cefd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CloudEventFiltersMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cefd.mutation = mutation
			affected, err = cefd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cefd.hooks) - 1; i >= 0; i-- {
			if cefd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cefd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cefd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cefd *CloudEventFiltersDelete) ExecX(ctx context.Context) int {
	n, err := cefd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cefd *CloudEventFiltersDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: cloudeventfilters.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cloudeventfilters.FieldID,
			},
		},
	}
	if ps := cefd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cefd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CloudEventFiltersDeleteOne is the builder for deleting a single CloudEventFilters entity.
type CloudEventFiltersDeleteOne struct {
	cefd *CloudEventFiltersDelete
}

// Exec executes the deletion query.
func (cefdo *CloudEventFiltersDeleteOne) Exec(ctx context.Context) error {
	n, err := cefdo.cefd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cloudeventfilters.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cefdo *CloudEventFiltersDeleteOne) ExecX(ctx context.Context) {
	cefdo.cefd.ExecX(ctx)
}

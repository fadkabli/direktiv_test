// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/metrics/ent/metrics"
)

// MetricsCreate is the builder for creating a Metrics entity.
type MetricsCreate struct {
	config
	mutation *MetricsMutation
	hooks    []Hook
}

// SetNamespace sets the "namespace" field.
func (mc *MetricsCreate) SetNamespace(s string) *MetricsCreate {
	mc.mutation.SetNamespace(s)
	return mc
}

// SetWorkflow sets the "workflow" field.
func (mc *MetricsCreate) SetWorkflow(s string) *MetricsCreate {
	mc.mutation.SetWorkflow(s)
	return mc
}

// SetRevision sets the "revision" field.
func (mc *MetricsCreate) SetRevision(s string) *MetricsCreate {
	mc.mutation.SetRevision(s)
	return mc
}

// SetInstance sets the "instance" field.
func (mc *MetricsCreate) SetInstance(s string) *MetricsCreate {
	mc.mutation.SetInstance(s)
	return mc
}

// SetState sets the "state" field.
func (mc *MetricsCreate) SetState(s string) *MetricsCreate {
	mc.mutation.SetState(s)
	return mc
}

// SetTimestamp sets the "timestamp" field.
func (mc *MetricsCreate) SetTimestamp(t time.Time) *MetricsCreate {
	mc.mutation.SetTimestamp(t)
	return mc
}

// SetWorkflowMs sets the "workflow_ms" field.
func (mc *MetricsCreate) SetWorkflowMs(i int64) *MetricsCreate {
	mc.mutation.SetWorkflowMs(i)
	return mc
}

// SetIsolateMs sets the "isolate_ms" field.
func (mc *MetricsCreate) SetIsolateMs(i int64) *MetricsCreate {
	mc.mutation.SetIsolateMs(i)
	return mc
}

// SetErrorCode sets the "error_code" field.
func (mc *MetricsCreate) SetErrorCode(s string) *MetricsCreate {
	mc.mutation.SetErrorCode(s)
	return mc
}

// SetNillableErrorCode sets the "error_code" field if the given value is not nil.
func (mc *MetricsCreate) SetNillableErrorCode(s *string) *MetricsCreate {
	if s != nil {
		mc.SetErrorCode(*s)
	}
	return mc
}

// SetInvoker sets the "invoker" field.
func (mc *MetricsCreate) SetInvoker(s string) *MetricsCreate {
	mc.mutation.SetInvoker(s)
	return mc
}

// SetNext sets the "next" field.
func (mc *MetricsCreate) SetNext(i int8) *MetricsCreate {
	mc.mutation.SetNext(i)
	return mc
}

// SetTransition sets the "transition" field.
func (mc *MetricsCreate) SetTransition(s string) *MetricsCreate {
	mc.mutation.SetTransition(s)
	return mc
}

// SetNillableTransition sets the "transition" field if the given value is not nil.
func (mc *MetricsCreate) SetNillableTransition(s *string) *MetricsCreate {
	if s != nil {
		mc.SetTransition(*s)
	}
	return mc
}

// Mutation returns the MetricsMutation object of the builder.
func (mc *MetricsCreate) Mutation() *MetricsMutation {
	return mc.mutation
}

// Save creates the Metrics in the database.
func (mc *MetricsCreate) Save(ctx context.Context) (*Metrics, error) {
	var (
		err  error
		node *Metrics
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Metrics)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MetricsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetricsCreate) SaveX(ctx context.Context) *Metrics {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MetricsCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MetricsCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetricsCreate) check() error {
	if _, ok := mc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required field "Metrics.namespace"`)}
	}
	if v, ok := mc.mutation.Namespace(); ok {
		if err := metrics.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`ent: validator failed for field "Metrics.namespace": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Workflow(); !ok {
		return &ValidationError{Name: "workflow", err: errors.New(`ent: missing required field "Metrics.workflow"`)}
	}
	if v, ok := mc.mutation.Workflow(); ok {
		if err := metrics.WorkflowValidator(v); err != nil {
			return &ValidationError{Name: "workflow", err: fmt.Errorf(`ent: validator failed for field "Metrics.workflow": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Revision(); !ok {
		return &ValidationError{Name: "revision", err: errors.New(`ent: missing required field "Metrics.revision"`)}
	}
	if _, ok := mc.mutation.Instance(); !ok {
		return &ValidationError{Name: "instance", err: errors.New(`ent: missing required field "Metrics.instance"`)}
	}
	if v, ok := mc.mutation.Instance(); ok {
		if err := metrics.InstanceValidator(v); err != nil {
			return &ValidationError{Name: "instance", err: fmt.Errorf(`ent: validator failed for field "Metrics.instance": %w`, err)}
		}
	}
	if _, ok := mc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "Metrics.state"`)}
	}
	if v, ok := mc.mutation.State(); ok {
		if err := metrics.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Metrics.state": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Metrics.timestamp"`)}
	}
	if _, ok := mc.mutation.WorkflowMs(); !ok {
		return &ValidationError{Name: "workflow_ms", err: errors.New(`ent: missing required field "Metrics.workflow_ms"`)}
	}
	if v, ok := mc.mutation.WorkflowMs(); ok {
		if err := metrics.WorkflowMsValidator(v); err != nil {
			return &ValidationError{Name: "workflow_ms", err: fmt.Errorf(`ent: validator failed for field "Metrics.workflow_ms": %w`, err)}
		}
	}
	if _, ok := mc.mutation.IsolateMs(); !ok {
		return &ValidationError{Name: "isolate_ms", err: errors.New(`ent: missing required field "Metrics.isolate_ms"`)}
	}
	if v, ok := mc.mutation.IsolateMs(); ok {
		if err := metrics.IsolateMsValidator(v); err != nil {
			return &ValidationError{Name: "isolate_ms", err: fmt.Errorf(`ent: validator failed for field "Metrics.isolate_ms": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Invoker(); !ok {
		return &ValidationError{Name: "invoker", err: errors.New(`ent: missing required field "Metrics.invoker"`)}
	}
	if _, ok := mc.mutation.Next(); !ok {
		return &ValidationError{Name: "next", err: errors.New(`ent: missing required field "Metrics.next"`)}
	}
	if v, ok := mc.mutation.Next(); ok {
		if err := metrics.NextValidator(v); err != nil {
			return &ValidationError{Name: "next", err: fmt.Errorf(`ent: validator failed for field "Metrics.next": %w`, err)}
		}
	}
	return nil
}

func (mc *MetricsCreate) sqlSave(ctx context.Context) (*Metrics, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MetricsCreate) createSpec() (*Metrics, *sqlgraph.CreateSpec) {
	var (
		_node = &Metrics{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: metrics.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metrics.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Namespace(); ok {
		_spec.SetField(metrics.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := mc.mutation.Workflow(); ok {
		_spec.SetField(metrics.FieldWorkflow, field.TypeString, value)
		_node.Workflow = value
	}
	if value, ok := mc.mutation.Revision(); ok {
		_spec.SetField(metrics.FieldRevision, field.TypeString, value)
		_node.Revision = value
	}
	if value, ok := mc.mutation.Instance(); ok {
		_spec.SetField(metrics.FieldInstance, field.TypeString, value)
		_node.Instance = value
	}
	if value, ok := mc.mutation.State(); ok {
		_spec.SetField(metrics.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := mc.mutation.Timestamp(); ok {
		_spec.SetField(metrics.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if value, ok := mc.mutation.WorkflowMs(); ok {
		_spec.SetField(metrics.FieldWorkflowMs, field.TypeInt64, value)
		_node.WorkflowMs = value
	}
	if value, ok := mc.mutation.IsolateMs(); ok {
		_spec.SetField(metrics.FieldIsolateMs, field.TypeInt64, value)
		_node.IsolateMs = value
	}
	if value, ok := mc.mutation.ErrorCode(); ok {
		_spec.SetField(metrics.FieldErrorCode, field.TypeString, value)
		_node.ErrorCode = value
	}
	if value, ok := mc.mutation.Invoker(); ok {
		_spec.SetField(metrics.FieldInvoker, field.TypeString, value)
		_node.Invoker = value
	}
	if value, ok := mc.mutation.Next(); ok {
		_spec.SetField(metrics.FieldNext, field.TypeInt8, value)
		_node.Next = value
	}
	if value, ok := mc.mutation.Transition(); ok {
		_spec.SetField(metrics.FieldTransition, field.TypeString, value)
		_node.Transition = value
	}
	return _node, _spec
}

// MetricsCreateBulk is the builder for creating many Metrics entities in bulk.
type MetricsCreateBulk struct {
	config
	builders []*MetricsCreate
}

// Save creates the Metrics entities in the database.
func (mcb *MetricsCreateBulk) Save(ctx context.Context) ([]*Metrics, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Metrics, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetricsMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MetricsCreateBulk) SaveX(ctx context.Context) []*Metrics {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MetricsCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MetricsCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

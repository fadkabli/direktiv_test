// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/logtag"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirroractivity"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// LogMsgUpdate is the builder for updating LogMsg entities.
type LogMsgUpdate struct {
	config
	hooks     []Hook
	mutation  *LogMsgMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LogMsgUpdate builder.
func (lmu *LogMsgUpdate) Where(ps ...predicate.LogMsg) *LogMsgUpdate {
	lmu.mutation.Where(ps...)
	return lmu
}

// SetT sets the "t" field.
func (lmu *LogMsgUpdate) SetT(t time.Time) *LogMsgUpdate {
	lmu.mutation.SetT(t)
	return lmu
}

// SetMsg sets the "msg" field.
func (lmu *LogMsgUpdate) SetMsg(s string) *LogMsgUpdate {
	lmu.mutation.SetMsg(s)
	return lmu
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (lmu *LogMsgUpdate) SetNamespaceID(id uuid.UUID) *LogMsgUpdate {
	lmu.mutation.SetNamespaceID(id)
	return lmu
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableNamespaceID(id *uuid.UUID) *LogMsgUpdate {
	if id != nil {
		lmu = lmu.SetNamespaceID(*id)
	}
	return lmu
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (lmu *LogMsgUpdate) SetNamespace(n *Namespace) *LogMsgUpdate {
	return lmu.SetNamespaceID(n.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (lmu *LogMsgUpdate) SetWorkflowID(id uuid.UUID) *LogMsgUpdate {
	lmu.mutation.SetWorkflowID(id)
	return lmu
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableWorkflowID(id *uuid.UUID) *LogMsgUpdate {
	if id != nil {
		lmu = lmu.SetWorkflowID(*id)
	}
	return lmu
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (lmu *LogMsgUpdate) SetWorkflow(w *Workflow) *LogMsgUpdate {
	return lmu.SetWorkflowID(w.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (lmu *LogMsgUpdate) SetInstanceID(id uuid.UUID) *LogMsgUpdate {
	lmu.mutation.SetInstanceID(id)
	return lmu
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableInstanceID(id *uuid.UUID) *LogMsgUpdate {
	if id != nil {
		lmu = lmu.SetInstanceID(*id)
	}
	return lmu
}

// SetInstance sets the "instance" edge to the Instance entity.
func (lmu *LogMsgUpdate) SetInstance(i *Instance) *LogMsgUpdate {
	return lmu.SetInstanceID(i.ID)
}

// SetActivityID sets the "activity" edge to the MirrorActivity entity by ID.
func (lmu *LogMsgUpdate) SetActivityID(id uuid.UUID) *LogMsgUpdate {
	lmu.mutation.SetActivityID(id)
	return lmu
}

// SetNillableActivityID sets the "activity" edge to the MirrorActivity entity by ID if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableActivityID(id *uuid.UUID) *LogMsgUpdate {
	if id != nil {
		lmu = lmu.SetActivityID(*id)
	}
	return lmu
}

// SetActivity sets the "activity" edge to the MirrorActivity entity.
func (lmu *LogMsgUpdate) SetActivity(m *MirrorActivity) *LogMsgUpdate {
	return lmu.SetActivityID(m.ID)
}

// AddLogtagIDs adds the "logtag" edge to the LogTag entity by IDs.
func (lmu *LogMsgUpdate) AddLogtagIDs(ids ...uuid.UUID) *LogMsgUpdate {
	lmu.mutation.AddLogtagIDs(ids...)
	return lmu
}

// AddLogtag adds the "logtag" edges to the LogTag entity.
func (lmu *LogMsgUpdate) AddLogtag(l ...*LogTag) *LogMsgUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lmu.AddLogtagIDs(ids...)
}

// Mutation returns the LogMsgMutation object of the builder.
func (lmu *LogMsgUpdate) Mutation() *LogMsgMutation {
	return lmu.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (lmu *LogMsgUpdate) ClearNamespace() *LogMsgUpdate {
	lmu.mutation.ClearNamespace()
	return lmu
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (lmu *LogMsgUpdate) ClearWorkflow() *LogMsgUpdate {
	lmu.mutation.ClearWorkflow()
	return lmu
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (lmu *LogMsgUpdate) ClearInstance() *LogMsgUpdate {
	lmu.mutation.ClearInstance()
	return lmu
}

// ClearActivity clears the "activity" edge to the MirrorActivity entity.
func (lmu *LogMsgUpdate) ClearActivity() *LogMsgUpdate {
	lmu.mutation.ClearActivity()
	return lmu
}

// ClearLogtag clears all "logtag" edges to the LogTag entity.
func (lmu *LogMsgUpdate) ClearLogtag() *LogMsgUpdate {
	lmu.mutation.ClearLogtag()
	return lmu
}

// RemoveLogtagIDs removes the "logtag" edge to LogTag entities by IDs.
func (lmu *LogMsgUpdate) RemoveLogtagIDs(ids ...uuid.UUID) *LogMsgUpdate {
	lmu.mutation.RemoveLogtagIDs(ids...)
	return lmu
}

// RemoveLogtag removes "logtag" edges to LogTag entities.
func (lmu *LogMsgUpdate) RemoveLogtag(l ...*LogTag) *LogMsgUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lmu.RemoveLogtagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lmu *LogMsgUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lmu.hooks) == 0 {
		affected, err = lmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LogMsgMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lmu.mutation = mutation
			affected, err = lmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lmu.hooks) - 1; i >= 0; i-- {
			if lmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lmu *LogMsgUpdate) SaveX(ctx context.Context) int {
	affected, err := lmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lmu *LogMsgUpdate) Exec(ctx context.Context) error {
	_, err := lmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmu *LogMsgUpdate) ExecX(ctx context.Context) {
	if err := lmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lmu *LogMsgUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LogMsgUpdate {
	lmu.modifiers = append(lmu.modifiers, modifiers...)
	return lmu
}

func (lmu *LogMsgUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   logmsg.Table,
			Columns: logmsg.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: logmsg.FieldID,
			},
		},
	}
	if ps := lmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmu.mutation.T(); ok {
		_spec.SetField(logmsg.FieldT, field.TypeTime, value)
	}
	if value, ok := lmu.mutation.Msg(); ok {
		_spec.SetField(logmsg.FieldMsg, field.TypeString, value)
	}
	if lmu.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
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
	if nodes := lmu.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
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
	if lmu.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.WorkflowTable,
			Columns: []string{logmsg.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.WorkflowTable,
			Columns: []string{logmsg.WorkflowColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmu.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmu.mutation.ActivityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.ActivityTable,
			Columns: []string{logmsg.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirroractivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.ActivityTable,
			Columns: []string{logmsg.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirroractivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmu.mutation.LogtagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   logmsg.LogtagTable,
			Columns: []string{logmsg.LogtagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logtag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.RemovedLogtagIDs(); len(nodes) > 0 && !lmu.mutation.LogtagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   logmsg.LogtagTable,
			Columns: []string{logmsg.LogtagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logtag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.LogtagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   logmsg.LogtagTable,
			Columns: []string{logmsg.LogtagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logtag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(lmu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, lmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logmsg.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LogMsgUpdateOne is the builder for updating a single LogMsg entity.
type LogMsgUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LogMsgMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetT sets the "t" field.
func (lmuo *LogMsgUpdateOne) SetT(t time.Time) *LogMsgUpdateOne {
	lmuo.mutation.SetT(t)
	return lmuo
}

// SetMsg sets the "msg" field.
func (lmuo *LogMsgUpdateOne) SetMsg(s string) *LogMsgUpdateOne {
	lmuo.mutation.SetMsg(s)
	return lmuo
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (lmuo *LogMsgUpdateOne) SetNamespaceID(id uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.SetNamespaceID(id)
	return lmuo
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableNamespaceID(id *uuid.UUID) *LogMsgUpdateOne {
	if id != nil {
		lmuo = lmuo.SetNamespaceID(*id)
	}
	return lmuo
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (lmuo *LogMsgUpdateOne) SetNamespace(n *Namespace) *LogMsgUpdateOne {
	return lmuo.SetNamespaceID(n.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (lmuo *LogMsgUpdateOne) SetWorkflowID(id uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.SetWorkflowID(id)
	return lmuo
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableWorkflowID(id *uuid.UUID) *LogMsgUpdateOne {
	if id != nil {
		lmuo = lmuo.SetWorkflowID(*id)
	}
	return lmuo
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (lmuo *LogMsgUpdateOne) SetWorkflow(w *Workflow) *LogMsgUpdateOne {
	return lmuo.SetWorkflowID(w.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (lmuo *LogMsgUpdateOne) SetInstanceID(id uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.SetInstanceID(id)
	return lmuo
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableInstanceID(id *uuid.UUID) *LogMsgUpdateOne {
	if id != nil {
		lmuo = lmuo.SetInstanceID(*id)
	}
	return lmuo
}

// SetInstance sets the "instance" edge to the Instance entity.
func (lmuo *LogMsgUpdateOne) SetInstance(i *Instance) *LogMsgUpdateOne {
	return lmuo.SetInstanceID(i.ID)
}

// SetActivityID sets the "activity" edge to the MirrorActivity entity by ID.
func (lmuo *LogMsgUpdateOne) SetActivityID(id uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.SetActivityID(id)
	return lmuo
}

// SetNillableActivityID sets the "activity" edge to the MirrorActivity entity by ID if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableActivityID(id *uuid.UUID) *LogMsgUpdateOne {
	if id != nil {
		lmuo = lmuo.SetActivityID(*id)
	}
	return lmuo
}

// SetActivity sets the "activity" edge to the MirrorActivity entity.
func (lmuo *LogMsgUpdateOne) SetActivity(m *MirrorActivity) *LogMsgUpdateOne {
	return lmuo.SetActivityID(m.ID)
}

// AddLogtagIDs adds the "logtag" edge to the LogTag entity by IDs.
func (lmuo *LogMsgUpdateOne) AddLogtagIDs(ids ...uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.AddLogtagIDs(ids...)
	return lmuo
}

// AddLogtag adds the "logtag" edges to the LogTag entity.
func (lmuo *LogMsgUpdateOne) AddLogtag(l ...*LogTag) *LogMsgUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lmuo.AddLogtagIDs(ids...)
}

// Mutation returns the LogMsgMutation object of the builder.
func (lmuo *LogMsgUpdateOne) Mutation() *LogMsgMutation {
	return lmuo.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (lmuo *LogMsgUpdateOne) ClearNamespace() *LogMsgUpdateOne {
	lmuo.mutation.ClearNamespace()
	return lmuo
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (lmuo *LogMsgUpdateOne) ClearWorkflow() *LogMsgUpdateOne {
	lmuo.mutation.ClearWorkflow()
	return lmuo
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (lmuo *LogMsgUpdateOne) ClearInstance() *LogMsgUpdateOne {
	lmuo.mutation.ClearInstance()
	return lmuo
}

// ClearActivity clears the "activity" edge to the MirrorActivity entity.
func (lmuo *LogMsgUpdateOne) ClearActivity() *LogMsgUpdateOne {
	lmuo.mutation.ClearActivity()
	return lmuo
}

// ClearLogtag clears all "logtag" edges to the LogTag entity.
func (lmuo *LogMsgUpdateOne) ClearLogtag() *LogMsgUpdateOne {
	lmuo.mutation.ClearLogtag()
	return lmuo
}

// RemoveLogtagIDs removes the "logtag" edge to LogTag entities by IDs.
func (lmuo *LogMsgUpdateOne) RemoveLogtagIDs(ids ...uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.RemoveLogtagIDs(ids...)
	return lmuo
}

// RemoveLogtag removes "logtag" edges to LogTag entities.
func (lmuo *LogMsgUpdateOne) RemoveLogtag(l ...*LogTag) *LogMsgUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lmuo.RemoveLogtagIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lmuo *LogMsgUpdateOne) Select(field string, fields ...string) *LogMsgUpdateOne {
	lmuo.fields = append([]string{field}, fields...)
	return lmuo
}

// Save executes the query and returns the updated LogMsg entity.
func (lmuo *LogMsgUpdateOne) Save(ctx context.Context) (*LogMsg, error) {
	var (
		err  error
		node *LogMsg
	)
	if len(lmuo.hooks) == 0 {
		node, err = lmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LogMsgMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lmuo.mutation = mutation
			node, err = lmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lmuo.hooks) - 1; i >= 0; i-- {
			if lmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LogMsg)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LogMsgMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lmuo *LogMsgUpdateOne) SaveX(ctx context.Context) *LogMsg {
	node, err := lmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lmuo *LogMsgUpdateOne) Exec(ctx context.Context) error {
	_, err := lmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmuo *LogMsgUpdateOne) ExecX(ctx context.Context) {
	if err := lmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lmuo *LogMsgUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LogMsgUpdateOne {
	lmuo.modifiers = append(lmuo.modifiers, modifiers...)
	return lmuo
}

func (lmuo *LogMsgUpdateOne) sqlSave(ctx context.Context) (_node *LogMsg, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   logmsg.Table,
			Columns: logmsg.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: logmsg.FieldID,
			},
		},
	}
	id, ok := lmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LogMsg.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logmsg.FieldID)
		for _, f := range fields {
			if !logmsg.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != logmsg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmuo.mutation.T(); ok {
		_spec.SetField(logmsg.FieldT, field.TypeTime, value)
	}
	if value, ok := lmuo.mutation.Msg(); ok {
		_spec.SetField(logmsg.FieldMsg, field.TypeString, value)
	}
	if lmuo.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
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
	if nodes := lmuo.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
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
	if lmuo.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.WorkflowTable,
			Columns: []string{logmsg.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.WorkflowTable,
			Columns: []string{logmsg.WorkflowColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmuo.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmuo.mutation.ActivityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.ActivityTable,
			Columns: []string{logmsg.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirroractivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.ActivityTable,
			Columns: []string{logmsg.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirroractivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmuo.mutation.LogtagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   logmsg.LogtagTable,
			Columns: []string{logmsg.LogtagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logtag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.RemovedLogtagIDs(); len(nodes) > 0 && !lmuo.mutation.LogtagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   logmsg.LogtagTable,
			Columns: []string{logmsg.LogtagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logtag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.LogtagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   logmsg.LogtagTable,
			Columns: []string{logmsg.LogtagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logtag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(lmuo.modifiers...)
	_node = &LogMsg{config: lmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logmsg.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/ref"
	"github.com/direktiv/direktiv/pkg/flow/ent/revision"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// RevisionUpdate is the builder for updating Revision entities.
type RevisionUpdate struct {
	config
	hooks     []Hook
	mutation  *RevisionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RevisionUpdate builder.
func (ru *RevisionUpdate) Where(ps ...predicate.Revision) *RevisionUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetMetadata sets the "metadata" field.
func (ru *RevisionUpdate) SetMetadata(m map[string]interface{}) *RevisionUpdate {
	ru.mutation.SetMetadata(m)
	return ru
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ru *RevisionUpdate) SetWorkflowID(id uuid.UUID) *RevisionUpdate {
	ru.mutation.SetWorkflowID(id)
	return ru
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ru *RevisionUpdate) SetWorkflow(w *Workflow) *RevisionUpdate {
	return ru.SetWorkflowID(w.ID)
}

// AddRefIDs adds the "refs" edge to the Ref entity by IDs.
func (ru *RevisionUpdate) AddRefIDs(ids ...uuid.UUID) *RevisionUpdate {
	ru.mutation.AddRefIDs(ids...)
	return ru
}

// AddRefs adds the "refs" edges to the Ref entity.
func (ru *RevisionUpdate) AddRefs(r ...*Ref) *RevisionUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRefIDs(ids...)
}

// AddInstanceIDs adds the "instances" edge to the Instance entity by IDs.
func (ru *RevisionUpdate) AddInstanceIDs(ids ...uuid.UUID) *RevisionUpdate {
	ru.mutation.AddInstanceIDs(ids...)
	return ru
}

// AddInstances adds the "instances" edges to the Instance entity.
func (ru *RevisionUpdate) AddInstances(i ...*Instance) *RevisionUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.AddInstanceIDs(ids...)
}

// Mutation returns the RevisionMutation object of the builder.
func (ru *RevisionUpdate) Mutation() *RevisionMutation {
	return ru.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (ru *RevisionUpdate) ClearWorkflow() *RevisionUpdate {
	ru.mutation.ClearWorkflow()
	return ru
}

// ClearRefs clears all "refs" edges to the Ref entity.
func (ru *RevisionUpdate) ClearRefs() *RevisionUpdate {
	ru.mutation.ClearRefs()
	return ru
}

// RemoveRefIDs removes the "refs" edge to Ref entities by IDs.
func (ru *RevisionUpdate) RemoveRefIDs(ids ...uuid.UUID) *RevisionUpdate {
	ru.mutation.RemoveRefIDs(ids...)
	return ru
}

// RemoveRefs removes "refs" edges to Ref entities.
func (ru *RevisionUpdate) RemoveRefs(r ...*Ref) *RevisionUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRefIDs(ids...)
}

// ClearInstances clears all "instances" edges to the Instance entity.
func (ru *RevisionUpdate) ClearInstances() *RevisionUpdate {
	ru.mutation.ClearInstances()
	return ru
}

// RemoveInstanceIDs removes the "instances" edge to Instance entities by IDs.
func (ru *RevisionUpdate) RemoveInstanceIDs(ids ...uuid.UUID) *RevisionUpdate {
	ru.mutation.RemoveInstanceIDs(ids...)
	return ru
}

// RemoveInstances removes "instances" edges to Instance entities.
func (ru *RevisionUpdate) RemoveInstances(i ...*Instance) *RevisionUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.RemoveInstanceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RevisionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RevisionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RevisionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RevisionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RevisionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RevisionUpdate) check() error {
	if _, ok := ru.mutation.WorkflowID(); ru.mutation.WorkflowCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Revision.workflow"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *RevisionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RevisionUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *RevisionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   revision.Table,
			Columns: revision.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: revision.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Metadata(); ok {
		_spec.SetField(revision.FieldMetadata, field.TypeJSON, value)
	}
	if ru.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   revision.WorkflowTable,
			Columns: []string{revision.WorkflowColumn},
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
	if nodes := ru.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   revision.WorkflowTable,
			Columns: []string{revision.WorkflowColumn},
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
	if ru.mutation.RefsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.RefsTable,
			Columns: []string{revision.RefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ref.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRefsIDs(); len(nodes) > 0 && !ru.mutation.RefsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.RefsTable,
			Columns: []string{revision.RefsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RefsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.RefsTable,
			Columns: []string{revision.RefsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.InstancesTable,
			Columns: []string{revision.InstancesColumn},
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
	if nodes := ru.mutation.RemovedInstancesIDs(); len(nodes) > 0 && !ru.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.InstancesTable,
			Columns: []string{revision.InstancesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.InstancesTable,
			Columns: []string{revision.InstancesColumn},
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
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{revision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RevisionUpdateOne is the builder for updating a single Revision entity.
type RevisionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RevisionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetMetadata sets the "metadata" field.
func (ruo *RevisionUpdateOne) SetMetadata(m map[string]interface{}) *RevisionUpdateOne {
	ruo.mutation.SetMetadata(m)
	return ruo
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ruo *RevisionUpdateOne) SetWorkflowID(id uuid.UUID) *RevisionUpdateOne {
	ruo.mutation.SetWorkflowID(id)
	return ruo
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ruo *RevisionUpdateOne) SetWorkflow(w *Workflow) *RevisionUpdateOne {
	return ruo.SetWorkflowID(w.ID)
}

// AddRefIDs adds the "refs" edge to the Ref entity by IDs.
func (ruo *RevisionUpdateOne) AddRefIDs(ids ...uuid.UUID) *RevisionUpdateOne {
	ruo.mutation.AddRefIDs(ids...)
	return ruo
}

// AddRefs adds the "refs" edges to the Ref entity.
func (ruo *RevisionUpdateOne) AddRefs(r ...*Ref) *RevisionUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRefIDs(ids...)
}

// AddInstanceIDs adds the "instances" edge to the Instance entity by IDs.
func (ruo *RevisionUpdateOne) AddInstanceIDs(ids ...uuid.UUID) *RevisionUpdateOne {
	ruo.mutation.AddInstanceIDs(ids...)
	return ruo
}

// AddInstances adds the "instances" edges to the Instance entity.
func (ruo *RevisionUpdateOne) AddInstances(i ...*Instance) *RevisionUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.AddInstanceIDs(ids...)
}

// Mutation returns the RevisionMutation object of the builder.
func (ruo *RevisionUpdateOne) Mutation() *RevisionMutation {
	return ruo.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (ruo *RevisionUpdateOne) ClearWorkflow() *RevisionUpdateOne {
	ruo.mutation.ClearWorkflow()
	return ruo
}

// ClearRefs clears all "refs" edges to the Ref entity.
func (ruo *RevisionUpdateOne) ClearRefs() *RevisionUpdateOne {
	ruo.mutation.ClearRefs()
	return ruo
}

// RemoveRefIDs removes the "refs" edge to Ref entities by IDs.
func (ruo *RevisionUpdateOne) RemoveRefIDs(ids ...uuid.UUID) *RevisionUpdateOne {
	ruo.mutation.RemoveRefIDs(ids...)
	return ruo
}

// RemoveRefs removes "refs" edges to Ref entities.
func (ruo *RevisionUpdateOne) RemoveRefs(r ...*Ref) *RevisionUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRefIDs(ids...)
}

// ClearInstances clears all "instances" edges to the Instance entity.
func (ruo *RevisionUpdateOne) ClearInstances() *RevisionUpdateOne {
	ruo.mutation.ClearInstances()
	return ruo
}

// RemoveInstanceIDs removes the "instances" edge to Instance entities by IDs.
func (ruo *RevisionUpdateOne) RemoveInstanceIDs(ids ...uuid.UUID) *RevisionUpdateOne {
	ruo.mutation.RemoveInstanceIDs(ids...)
	return ruo
}

// RemoveInstances removes "instances" edges to Instance entities.
func (ruo *RevisionUpdateOne) RemoveInstances(i ...*Instance) *RevisionUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.RemoveInstanceIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RevisionUpdateOne) Select(field string, fields ...string) *RevisionUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Revision entity.
func (ruo *RevisionUpdateOne) Save(ctx context.Context) (*Revision, error) {
	var (
		err  error
		node *Revision
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RevisionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Revision)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RevisionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RevisionUpdateOne) SaveX(ctx context.Context) *Revision {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RevisionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RevisionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RevisionUpdateOne) check() error {
	if _, ok := ruo.mutation.WorkflowID(); ruo.mutation.WorkflowCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Revision.workflow"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *RevisionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RevisionUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *RevisionUpdateOne) sqlSave(ctx context.Context) (_node *Revision, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   revision.Table,
			Columns: revision.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: revision.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Revision.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, revision.FieldID)
		for _, f := range fields {
			if !revision.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != revision.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Metadata(); ok {
		_spec.SetField(revision.FieldMetadata, field.TypeJSON, value)
	}
	if ruo.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   revision.WorkflowTable,
			Columns: []string{revision.WorkflowColumn},
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
	if nodes := ruo.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   revision.WorkflowTable,
			Columns: []string{revision.WorkflowColumn},
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
	if ruo.mutation.RefsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.RefsTable,
			Columns: []string{revision.RefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ref.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRefsIDs(); len(nodes) > 0 && !ruo.mutation.RefsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.RefsTable,
			Columns: []string{revision.RefsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RefsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.RefsTable,
			Columns: []string{revision.RefsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.InstancesTable,
			Columns: []string{revision.InstancesColumn},
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
	if nodes := ruo.mutation.RemovedInstancesIDs(); len(nodes) > 0 && !ruo.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.InstancesTable,
			Columns: []string{revision.InstancesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.InstancesTable,
			Columns: []string{revision.InstancesColumn},
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
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Revision{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{revision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

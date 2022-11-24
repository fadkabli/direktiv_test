// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/annotation"
	"github.com/direktiv/direktiv/pkg/flow/ent/inode"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirror"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// InodeCreate is the builder for creating a Inode entity.
type InodeCreate struct {
	config
	mutation *InodeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ic *InodeCreate) SetCreatedAt(t time.Time) *InodeCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *InodeCreate) SetNillableCreatedAt(t *time.Time) *InodeCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *InodeCreate) SetUpdatedAt(t time.Time) *InodeCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *InodeCreate) SetNillableUpdatedAt(t *time.Time) *InodeCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetName sets the "name" field.
func (ic *InodeCreate) SetName(s string) *InodeCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ic *InodeCreate) SetNillableName(s *string) *InodeCreate {
	if s != nil {
		ic.SetName(*s)
	}
	return ic
}

// SetType sets the "type" field.
func (ic *InodeCreate) SetType(s string) *InodeCreate {
	ic.mutation.SetType(s)
	return ic
}

// SetAttributes sets the "attributes" field.
func (ic *InodeCreate) SetAttributes(s []string) *InodeCreate {
	ic.mutation.SetAttributes(s)
	return ic
}

// SetExtendedType sets the "extended_type" field.
func (ic *InodeCreate) SetExtendedType(s string) *InodeCreate {
	ic.mutation.SetExtendedType(s)
	return ic
}

// SetNillableExtendedType sets the "extended_type" field if the given value is not nil.
func (ic *InodeCreate) SetNillableExtendedType(s *string) *InodeCreate {
	if s != nil {
		ic.SetExtendedType(*s)
	}
	return ic
}

// SetReadOnly sets the "readOnly" field.
func (ic *InodeCreate) SetReadOnly(b bool) *InodeCreate {
	ic.mutation.SetReadOnly(b)
	return ic
}

// SetNillableReadOnly sets the "readOnly" field if the given value is not nil.
func (ic *InodeCreate) SetNillableReadOnly(b *bool) *InodeCreate {
	if b != nil {
		ic.SetReadOnly(*b)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *InodeCreate) SetID(u uuid.UUID) *InodeCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *InodeCreate) SetNillableID(u *uuid.UUID) *InodeCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (ic *InodeCreate) SetNamespaceID(id uuid.UUID) *InodeCreate {
	ic.mutation.SetNamespaceID(id)
	return ic
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (ic *InodeCreate) SetNamespace(n *Namespace) *InodeCreate {
	return ic.SetNamespaceID(n.ID)
}

// AddChildIDs adds the "children" edge to the Inode entity by IDs.
func (ic *InodeCreate) AddChildIDs(ids ...uuid.UUID) *InodeCreate {
	ic.mutation.AddChildIDs(ids...)
	return ic
}

// AddChildren adds the "children" edges to the Inode entity.
func (ic *InodeCreate) AddChildren(i ...*Inode) *InodeCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ic.AddChildIDs(ids...)
}

// SetParentID sets the "parent" edge to the Inode entity by ID.
func (ic *InodeCreate) SetParentID(id uuid.UUID) *InodeCreate {
	ic.mutation.SetParentID(id)
	return ic
}

// SetNillableParentID sets the "parent" edge to the Inode entity by ID if the given value is not nil.
func (ic *InodeCreate) SetNillableParentID(id *uuid.UUID) *InodeCreate {
	if id != nil {
		ic = ic.SetParentID(*id)
	}
	return ic
}

// SetParent sets the "parent" edge to the Inode entity.
func (ic *InodeCreate) SetParent(i *Inode) *InodeCreate {
	return ic.SetParentID(i.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ic *InodeCreate) SetWorkflowID(id uuid.UUID) *InodeCreate {
	ic.mutation.SetWorkflowID(id)
	return ic
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (ic *InodeCreate) SetNillableWorkflowID(id *uuid.UUID) *InodeCreate {
	if id != nil {
		ic = ic.SetWorkflowID(*id)
	}
	return ic
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ic *InodeCreate) SetWorkflow(w *Workflow) *InodeCreate {
	return ic.SetWorkflowID(w.ID)
}

// SetMirrorID sets the "mirror" edge to the Mirror entity by ID.
func (ic *InodeCreate) SetMirrorID(id uuid.UUID) *InodeCreate {
	ic.mutation.SetMirrorID(id)
	return ic
}

// SetNillableMirrorID sets the "mirror" edge to the Mirror entity by ID if the given value is not nil.
func (ic *InodeCreate) SetNillableMirrorID(id *uuid.UUID) *InodeCreate {
	if id != nil {
		ic = ic.SetMirrorID(*id)
	}
	return ic
}

// SetMirror sets the "mirror" edge to the Mirror entity.
func (ic *InodeCreate) SetMirror(m *Mirror) *InodeCreate {
	return ic.SetMirrorID(m.ID)
}

// AddAnnotationIDs adds the "annotations" edge to the Annotation entity by IDs.
func (ic *InodeCreate) AddAnnotationIDs(ids ...uuid.UUID) *InodeCreate {
	ic.mutation.AddAnnotationIDs(ids...)
	return ic
}

// AddAnnotations adds the "annotations" edges to the Annotation entity.
func (ic *InodeCreate) AddAnnotations(a ...*Annotation) *InodeCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ic.AddAnnotationIDs(ids...)
}

// Mutation returns the InodeMutation object of the builder.
func (ic *InodeCreate) Mutation() *InodeMutation {
	return ic.mutation
}

// Save creates the Inode in the database.
func (ic *InodeCreate) Save(ctx context.Context) (*Inode, error) {
	var (
		err  error
		node *Inode
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Inode)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from InodeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InodeCreate) SaveX(ctx context.Context) *Inode {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InodeCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InodeCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InodeCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := inode.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := inode.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.ReadOnly(); !ok {
		v := inode.DefaultReadOnly
		ic.mutation.SetReadOnly(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := inode.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InodeCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Inode.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Inode.updated_at"`)}
	}
	if v, ok := ic.mutation.Name(); ok {
		if err := inode.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Inode.name": %w`, err)}
		}
	}
	if _, ok := ic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Inode.type"`)}
	}
	if _, ok := ic.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required edge "Inode.namespace"`)}
	}
	return nil
}

func (ic *InodeCreate) sqlSave(ctx context.Context) (*Inode, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
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

func (ic *InodeCreate) createSpec() (*Inode, *sqlgraph.CreateSpec) {
	var (
		_node = &Inode{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: inode.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: inode.FieldID,
			},
		}
	)
	_spec.OnConflict = ic.conflict
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(inode.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(inode.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(inode.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.GetType(); ok {
		_spec.SetField(inode.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := ic.mutation.Attributes(); ok {
		_spec.SetField(inode.FieldAttributes, field.TypeJSON, value)
		_node.Attributes = value
	}
	if value, ok := ic.mutation.ExtendedType(); ok {
		_spec.SetField(inode.FieldExtendedType, field.TypeString, value)
		_node.ExtendedType = value
	}
	if value, ok := ic.mutation.ReadOnly(); ok {
		_spec.SetField(inode.FieldReadOnly, field.TypeBool, value)
		_node.ReadOnly = value
	}
	if nodes := ic.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inode.NamespaceTable,
			Columns: []string{inode.NamespaceColumn},
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
		_node.namespace_inodes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inode.ChildrenTable,
			Columns: []string{inode.ChildrenColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inode.ParentTable,
			Columns: []string{inode.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.inode_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   inode.WorkflowTable,
			Columns: []string{inode.WorkflowColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.MirrorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   inode.MirrorTable,
			Columns: []string{inode.MirrorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirror.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.AnnotationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inode.AnnotationsTable,
			Columns: []string{inode.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: annotation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Inode.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InodeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ic *InodeCreate) OnConflict(opts ...sql.ConflictOption) *InodeUpsertOne {
	ic.conflict = opts
	return &InodeUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Inode.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ic *InodeCreate) OnConflictColumns(columns ...string) *InodeUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &InodeUpsertOne{
		create: ic,
	}
}

type (
	// InodeUpsertOne is the builder for "upsert"-ing
	//  one Inode node.
	InodeUpsertOne struct {
		create *InodeCreate
	}

	// InodeUpsert is the "OnConflict" setter.
	InodeUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *InodeUpsert) SetUpdatedAt(v time.Time) *InodeUpsert {
	u.Set(inode.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InodeUpsert) UpdateUpdatedAt() *InodeUpsert {
	u.SetExcluded(inode.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *InodeUpsert) SetName(v string) *InodeUpsert {
	u.Set(inode.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *InodeUpsert) UpdateName() *InodeUpsert {
	u.SetExcluded(inode.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *InodeUpsert) ClearName() *InodeUpsert {
	u.SetNull(inode.FieldName)
	return u
}

// SetAttributes sets the "attributes" field.
func (u *InodeUpsert) SetAttributes(v []string) *InodeUpsert {
	u.Set(inode.FieldAttributes, v)
	return u
}

// UpdateAttributes sets the "attributes" field to the value that was provided on create.
func (u *InodeUpsert) UpdateAttributes() *InodeUpsert {
	u.SetExcluded(inode.FieldAttributes)
	return u
}

// ClearAttributes clears the value of the "attributes" field.
func (u *InodeUpsert) ClearAttributes() *InodeUpsert {
	u.SetNull(inode.FieldAttributes)
	return u
}

// SetExtendedType sets the "extended_type" field.
func (u *InodeUpsert) SetExtendedType(v string) *InodeUpsert {
	u.Set(inode.FieldExtendedType, v)
	return u
}

// UpdateExtendedType sets the "extended_type" field to the value that was provided on create.
func (u *InodeUpsert) UpdateExtendedType() *InodeUpsert {
	u.SetExcluded(inode.FieldExtendedType)
	return u
}

// ClearExtendedType clears the value of the "extended_type" field.
func (u *InodeUpsert) ClearExtendedType() *InodeUpsert {
	u.SetNull(inode.FieldExtendedType)
	return u
}

// SetReadOnly sets the "readOnly" field.
func (u *InodeUpsert) SetReadOnly(v bool) *InodeUpsert {
	u.Set(inode.FieldReadOnly, v)
	return u
}

// UpdateReadOnly sets the "readOnly" field to the value that was provided on create.
func (u *InodeUpsert) UpdateReadOnly() *InodeUpsert {
	u.SetExcluded(inode.FieldReadOnly)
	return u
}

// ClearReadOnly clears the value of the "readOnly" field.
func (u *InodeUpsert) ClearReadOnly() *InodeUpsert {
	u.SetNull(inode.FieldReadOnly)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Inode.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(inode.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *InodeUpsertOne) UpdateNewValues() *InodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(inode.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(inode.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.GetType(); exists {
			s.SetIgnore(inode.FieldType)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Inode.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *InodeUpsertOne) Ignore() *InodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InodeUpsertOne) DoNothing() *InodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InodeCreate.OnConflict
// documentation for more info.
func (u *InodeUpsertOne) Update(set func(*InodeUpsert)) *InodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InodeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InodeUpsertOne) SetUpdatedAt(v time.Time) *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InodeUpsertOne) UpdateUpdatedAt() *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *InodeUpsertOne) SetName(v string) *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *InodeUpsertOne) UpdateName() *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *InodeUpsertOne) ClearName() *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.ClearName()
	})
}

// SetAttributes sets the "attributes" field.
func (u *InodeUpsertOne) SetAttributes(v []string) *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.SetAttributes(v)
	})
}

// UpdateAttributes sets the "attributes" field to the value that was provided on create.
func (u *InodeUpsertOne) UpdateAttributes() *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateAttributes()
	})
}

// ClearAttributes clears the value of the "attributes" field.
func (u *InodeUpsertOne) ClearAttributes() *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.ClearAttributes()
	})
}

// SetExtendedType sets the "extended_type" field.
func (u *InodeUpsertOne) SetExtendedType(v string) *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.SetExtendedType(v)
	})
}

// UpdateExtendedType sets the "extended_type" field to the value that was provided on create.
func (u *InodeUpsertOne) UpdateExtendedType() *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateExtendedType()
	})
}

// ClearExtendedType clears the value of the "extended_type" field.
func (u *InodeUpsertOne) ClearExtendedType() *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.ClearExtendedType()
	})
}

// SetReadOnly sets the "readOnly" field.
func (u *InodeUpsertOne) SetReadOnly(v bool) *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.SetReadOnly(v)
	})
}

// UpdateReadOnly sets the "readOnly" field to the value that was provided on create.
func (u *InodeUpsertOne) UpdateReadOnly() *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateReadOnly()
	})
}

// ClearReadOnly clears the value of the "readOnly" field.
func (u *InodeUpsertOne) ClearReadOnly() *InodeUpsertOne {
	return u.Update(func(s *InodeUpsert) {
		s.ClearReadOnly()
	})
}

// Exec executes the query.
func (u *InodeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for InodeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InodeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *InodeUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: InodeUpsertOne.ID is not supported by MySQL driver. Use InodeUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *InodeUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// InodeCreateBulk is the builder for creating many Inode entities in bulk.
type InodeCreateBulk struct {
	config
	builders []*InodeCreate
	conflict []sql.ConflictOption
}

// Save creates the Inode entities in the database.
func (icb *InodeCreateBulk) Save(ctx context.Context) ([]*Inode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Inode, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InodeMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InodeCreateBulk) SaveX(ctx context.Context) []*Inode {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InodeCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InodeCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Inode.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InodeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (icb *InodeCreateBulk) OnConflict(opts ...sql.ConflictOption) *InodeUpsertBulk {
	icb.conflict = opts
	return &InodeUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Inode.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (icb *InodeCreateBulk) OnConflictColumns(columns ...string) *InodeUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &InodeUpsertBulk{
		create: icb,
	}
}

// InodeUpsertBulk is the builder for "upsert"-ing
// a bulk of Inode nodes.
type InodeUpsertBulk struct {
	create *InodeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Inode.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(inode.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *InodeUpsertBulk) UpdateNewValues() *InodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(inode.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(inode.FieldCreatedAt)
			}
			if _, exists := b.mutation.GetType(); exists {
				s.SetIgnore(inode.FieldType)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Inode.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *InodeUpsertBulk) Ignore() *InodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InodeUpsertBulk) DoNothing() *InodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InodeCreateBulk.OnConflict
// documentation for more info.
func (u *InodeUpsertBulk) Update(set func(*InodeUpsert)) *InodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InodeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InodeUpsertBulk) SetUpdatedAt(v time.Time) *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InodeUpsertBulk) UpdateUpdatedAt() *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *InodeUpsertBulk) SetName(v string) *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *InodeUpsertBulk) UpdateName() *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *InodeUpsertBulk) ClearName() *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.ClearName()
	})
}

// SetAttributes sets the "attributes" field.
func (u *InodeUpsertBulk) SetAttributes(v []string) *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.SetAttributes(v)
	})
}

// UpdateAttributes sets the "attributes" field to the value that was provided on create.
func (u *InodeUpsertBulk) UpdateAttributes() *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateAttributes()
	})
}

// ClearAttributes clears the value of the "attributes" field.
func (u *InodeUpsertBulk) ClearAttributes() *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.ClearAttributes()
	})
}

// SetExtendedType sets the "extended_type" field.
func (u *InodeUpsertBulk) SetExtendedType(v string) *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.SetExtendedType(v)
	})
}

// UpdateExtendedType sets the "extended_type" field to the value that was provided on create.
func (u *InodeUpsertBulk) UpdateExtendedType() *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateExtendedType()
	})
}

// ClearExtendedType clears the value of the "extended_type" field.
func (u *InodeUpsertBulk) ClearExtendedType() *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.ClearExtendedType()
	})
}

// SetReadOnly sets the "readOnly" field.
func (u *InodeUpsertBulk) SetReadOnly(v bool) *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.SetReadOnly(v)
	})
}

// UpdateReadOnly sets the "readOnly" field to the value that was provided on create.
func (u *InodeUpsertBulk) UpdateReadOnly() *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.UpdateReadOnly()
	})
}

// ClearReadOnly clears the value of the "readOnly" field.
func (u *InodeUpsertBulk) ClearReadOnly() *InodeUpsertBulk {
	return u.Update(func(s *InodeUpsert) {
		s.ClearReadOnly()
	})
}

// Exec executes the query.
func (u *InodeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the InodeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for InodeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InodeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

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
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/ref"
	"github.com/direktiv/direktiv/pkg/flow/ent/revision"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// RevisionCreate is the builder for creating a Revision entity.
type RevisionCreate struct {
	config
	mutation *RevisionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (rc *RevisionCreate) SetCreatedAt(t time.Time) *RevisionCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RevisionCreate) SetNillableCreatedAt(t *time.Time) *RevisionCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetHash sets the "hash" field.
func (rc *RevisionCreate) SetHash(s string) *RevisionCreate {
	rc.mutation.SetHash(s)
	return rc
}

// SetSource sets the "source" field.
func (rc *RevisionCreate) SetSource(b []byte) *RevisionCreate {
	rc.mutation.SetSource(b)
	return rc
}

// SetMetadata sets the "metadata" field.
func (rc *RevisionCreate) SetMetadata(m map[string]interface{}) *RevisionCreate {
	rc.mutation.SetMetadata(m)
	return rc
}

// SetID sets the "id" field.
func (rc *RevisionCreate) SetID(u uuid.UUID) *RevisionCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RevisionCreate) SetNillableID(u *uuid.UUID) *RevisionCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (rc *RevisionCreate) SetWorkflowID(id uuid.UUID) *RevisionCreate {
	rc.mutation.SetWorkflowID(id)
	return rc
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (rc *RevisionCreate) SetWorkflow(w *Workflow) *RevisionCreate {
	return rc.SetWorkflowID(w.ID)
}

// AddRefIDs adds the "refs" edge to the Ref entity by IDs.
func (rc *RevisionCreate) AddRefIDs(ids ...uuid.UUID) *RevisionCreate {
	rc.mutation.AddRefIDs(ids...)
	return rc
}

// AddRefs adds the "refs" edges to the Ref entity.
func (rc *RevisionCreate) AddRefs(r ...*Ref) *RevisionCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddRefIDs(ids...)
}

// AddInstanceIDs adds the "instances" edge to the Instance entity by IDs.
func (rc *RevisionCreate) AddInstanceIDs(ids ...uuid.UUID) *RevisionCreate {
	rc.mutation.AddInstanceIDs(ids...)
	return rc
}

// AddInstances adds the "instances" edges to the Instance entity.
func (rc *RevisionCreate) AddInstances(i ...*Instance) *RevisionCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return rc.AddInstanceIDs(ids...)
}

// Mutation returns the RevisionMutation object of the builder.
func (rc *RevisionCreate) Mutation() *RevisionMutation {
	return rc.mutation
}

// Save creates the Revision in the database.
func (rc *RevisionCreate) Save(ctx context.Context) (*Revision, error) {
	var (
		err  error
		node *Revision
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RevisionMutation)
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
		nv, ok := v.(*Revision)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RevisionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RevisionCreate) SaveX(ctx context.Context) *Revision {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RevisionCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RevisionCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RevisionCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := revision.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := revision.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RevisionCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Revision.created_at"`)}
	}
	if _, ok := rc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Revision.hash"`)}
	}
	if _, ok := rc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required field "Revision.source"`)}
	}
	if _, ok := rc.mutation.Metadata(); !ok {
		return &ValidationError{Name: "metadata", err: errors.New(`ent: missing required field "Revision.metadata"`)}
	}
	if _, ok := rc.mutation.WorkflowID(); !ok {
		return &ValidationError{Name: "workflow", err: errors.New(`ent: missing required edge "Revision.workflow"`)}
	}
	return nil
}

func (rc *RevisionCreate) sqlSave(ctx context.Context) (*Revision, error) {
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

func (rc *RevisionCreate) createSpec() (*Revision, *sqlgraph.CreateSpec) {
	var (
		_node = &Revision{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: revision.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: revision.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(revision.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.Hash(); ok {
		_spec.SetField(revision.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if value, ok := rc.mutation.Source(); ok {
		_spec.SetField(revision.FieldSource, field.TypeBytes, value)
		_node.Source = value
	}
	if value, ok := rc.mutation.Metadata(); ok {
		_spec.SetField(revision.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if nodes := rc.mutation.WorkflowIDs(); len(nodes) > 0 {
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
		_node.workflow_revisions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RefsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.InstancesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Revision.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RevisionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rc *RevisionCreate) OnConflict(opts ...sql.ConflictOption) *RevisionUpsertOne {
	rc.conflict = opts
	return &RevisionUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Revision.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RevisionCreate) OnConflictColumns(columns ...string) *RevisionUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RevisionUpsertOne{
		create: rc,
	}
}

type (
	// RevisionUpsertOne is the builder for "upsert"-ing
	//  one Revision node.
	RevisionUpsertOne struct {
		create *RevisionCreate
	}

	// RevisionUpsert is the "OnConflict" setter.
	RevisionUpsert struct {
		*sql.UpdateSet
	}
)

// SetMetadata sets the "metadata" field.
func (u *RevisionUpsert) SetMetadata(v map[string]interface{}) *RevisionUpsert {
	u.Set(revision.FieldMetadata, v)
	return u
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *RevisionUpsert) UpdateMetadata() *RevisionUpsert {
	u.SetExcluded(revision.FieldMetadata)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Revision.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(revision.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RevisionUpsertOne) UpdateNewValues() *RevisionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(revision.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(revision.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.Hash(); exists {
			s.SetIgnore(revision.FieldHash)
		}
		if _, exists := u.create.mutation.Source(); exists {
			s.SetIgnore(revision.FieldSource)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Revision.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RevisionUpsertOne) Ignore() *RevisionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RevisionUpsertOne) DoNothing() *RevisionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RevisionCreate.OnConflict
// documentation for more info.
func (u *RevisionUpsertOne) Update(set func(*RevisionUpsert)) *RevisionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RevisionUpsert{UpdateSet: update})
	}))
	return u
}

// SetMetadata sets the "metadata" field.
func (u *RevisionUpsertOne) SetMetadata(v map[string]interface{}) *RevisionUpsertOne {
	return u.Update(func(s *RevisionUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *RevisionUpsertOne) UpdateMetadata() *RevisionUpsertOne {
	return u.Update(func(s *RevisionUpsert) {
		s.UpdateMetadata()
	})
}

// Exec executes the query.
func (u *RevisionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RevisionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RevisionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RevisionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: RevisionUpsertOne.ID is not supported by MySQL driver. Use RevisionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RevisionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RevisionCreateBulk is the builder for creating many Revision entities in bulk.
type RevisionCreateBulk struct {
	config
	builders []*RevisionCreate
	conflict []sql.ConflictOption
}

// Save creates the Revision entities in the database.
func (rcb *RevisionCreateBulk) Save(ctx context.Context) ([]*Revision, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Revision, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RevisionMutation)
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
					spec.OnConflict = rcb.conflict
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
func (rcb *RevisionCreateBulk) SaveX(ctx context.Context) []*Revision {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RevisionCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RevisionCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Revision.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RevisionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rcb *RevisionCreateBulk) OnConflict(opts ...sql.ConflictOption) *RevisionUpsertBulk {
	rcb.conflict = opts
	return &RevisionUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Revision.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RevisionCreateBulk) OnConflictColumns(columns ...string) *RevisionUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RevisionUpsertBulk{
		create: rcb,
	}
}

// RevisionUpsertBulk is the builder for "upsert"-ing
// a bulk of Revision nodes.
type RevisionUpsertBulk struct {
	create *RevisionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Revision.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(revision.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RevisionUpsertBulk) UpdateNewValues() *RevisionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(revision.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(revision.FieldCreatedAt)
			}
			if _, exists := b.mutation.Hash(); exists {
				s.SetIgnore(revision.FieldHash)
			}
			if _, exists := b.mutation.Source(); exists {
				s.SetIgnore(revision.FieldSource)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Revision.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RevisionUpsertBulk) Ignore() *RevisionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RevisionUpsertBulk) DoNothing() *RevisionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RevisionCreateBulk.OnConflict
// documentation for more info.
func (u *RevisionUpsertBulk) Update(set func(*RevisionUpsert)) *RevisionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RevisionUpsert{UpdateSet: update})
	}))
	return u
}

// SetMetadata sets the "metadata" field.
func (u *RevisionUpsertBulk) SetMetadata(v map[string]interface{}) *RevisionUpsertBulk {
	return u.Update(func(s *RevisionUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *RevisionUpsertBulk) UpdateMetadata() *RevisionUpsertBulk {
	return u.Update(func(s *RevisionUpsert) {
		s.UpdateMetadata()
	})
}

// Exec executes the query.
func (u *RevisionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RevisionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RevisionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RevisionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

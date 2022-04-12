// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/inode"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirror"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirroractivity"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// MirrorUpdate is the builder for updating Mirror entities.
type MirrorUpdate struct {
	config
	hooks    []Hook
	mutation *MirrorMutation
}

// Where appends a list predicates to the MirrorUpdate builder.
func (mu *MirrorUpdate) Where(ps ...predicate.Mirror) *MirrorUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetURL sets the "url" field.
func (mu *MirrorUpdate) SetURL(s string) *MirrorUpdate {
	mu.mutation.SetURL(s)
	return mu
}

// SetRef sets the "ref" field.
func (mu *MirrorUpdate) SetRef(s string) *MirrorUpdate {
	mu.mutation.SetRef(s)
	return mu
}

// SetCron sets the "cron" field.
func (mu *MirrorUpdate) SetCron(s string) *MirrorUpdate {
	mu.mutation.SetCron(s)
	return mu
}

// SetPublicKey sets the "public_key" field.
func (mu *MirrorUpdate) SetPublicKey(s string) *MirrorUpdate {
	mu.mutation.SetPublicKey(s)
	return mu
}

// SetPrivateKey sets the "private_key" field.
func (mu *MirrorUpdate) SetPrivateKey(s string) *MirrorUpdate {
	mu.mutation.SetPrivateKey(s)
	return mu
}

// SetPassphrase sets the "passphrase" field.
func (mu *MirrorUpdate) SetPassphrase(s string) *MirrorUpdate {
	mu.mutation.SetPassphrase(s)
	return mu
}

// SetCommit sets the "commit" field.
func (mu *MirrorUpdate) SetCommit(s string) *MirrorUpdate {
	mu.mutation.SetCommit(s)
	return mu
}

// SetLocked sets the "locked" field.
func (mu *MirrorUpdate) SetLocked(b bool) *MirrorUpdate {
	mu.mutation.SetLocked(b)
	return mu
}

// SetLastSync sets the "last_sync" field.
func (mu *MirrorUpdate) SetLastSync(t time.Time) *MirrorUpdate {
	mu.mutation.SetLastSync(t)
	return mu
}

// SetNillableLastSync sets the "last_sync" field if the given value is not nil.
func (mu *MirrorUpdate) SetNillableLastSync(t *time.Time) *MirrorUpdate {
	if t != nil {
		mu.SetLastSync(*t)
	}
	return mu
}

// ClearLastSync clears the value of the "last_sync" field.
func (mu *MirrorUpdate) ClearLastSync() *MirrorUpdate {
	mu.mutation.ClearLastSync()
	return mu
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (mu *MirrorUpdate) SetNamespaceID(id uuid.UUID) *MirrorUpdate {
	mu.mutation.SetNamespaceID(id)
	return mu
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (mu *MirrorUpdate) SetNamespace(n *Namespace) *MirrorUpdate {
	return mu.SetNamespaceID(n.ID)
}

// SetInodeID sets the "inode" edge to the Inode entity by ID.
func (mu *MirrorUpdate) SetInodeID(id uuid.UUID) *MirrorUpdate {
	mu.mutation.SetInodeID(id)
	return mu
}

// SetNillableInodeID sets the "inode" edge to the Inode entity by ID if the given value is not nil.
func (mu *MirrorUpdate) SetNillableInodeID(id *uuid.UUID) *MirrorUpdate {
	if id != nil {
		mu = mu.SetInodeID(*id)
	}
	return mu
}

// SetInode sets the "inode" edge to the Inode entity.
func (mu *MirrorUpdate) SetInode(i *Inode) *MirrorUpdate {
	return mu.SetInodeID(i.ID)
}

// AddActivityIDs adds the "activities" edge to the MirrorActivity entity by IDs.
func (mu *MirrorUpdate) AddActivityIDs(ids ...uuid.UUID) *MirrorUpdate {
	mu.mutation.AddActivityIDs(ids...)
	return mu
}

// AddActivities adds the "activities" edges to the MirrorActivity entity.
func (mu *MirrorUpdate) AddActivities(m ...*MirrorActivity) *MirrorUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddActivityIDs(ids...)
}

// Mutation returns the MirrorMutation object of the builder.
func (mu *MirrorUpdate) Mutation() *MirrorMutation {
	return mu.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (mu *MirrorUpdate) ClearNamespace() *MirrorUpdate {
	mu.mutation.ClearNamespace()
	return mu
}

// ClearInode clears the "inode" edge to the Inode entity.
func (mu *MirrorUpdate) ClearInode() *MirrorUpdate {
	mu.mutation.ClearInode()
	return mu
}

// ClearActivities clears all "activities" edges to the MirrorActivity entity.
func (mu *MirrorUpdate) ClearActivities() *MirrorUpdate {
	mu.mutation.ClearActivities()
	return mu
}

// RemoveActivityIDs removes the "activities" edge to MirrorActivity entities by IDs.
func (mu *MirrorUpdate) RemoveActivityIDs(ids ...uuid.UUID) *MirrorUpdate {
	mu.mutation.RemoveActivityIDs(ids...)
	return mu
}

// RemoveActivities removes "activities" edges to MirrorActivity entities.
func (mu *MirrorUpdate) RemoveActivities(m ...*MirrorActivity) *MirrorUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveActivityIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MirrorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MirrorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MirrorUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MirrorUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MirrorUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MirrorUpdate) check() error {
	if _, ok := mu.mutation.NamespaceID(); mu.mutation.NamespaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mirror.namespace"`)
	}
	return nil
}

func (mu *MirrorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mirror.Table,
			Columns: mirror.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mirror.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldURL,
		})
	}
	if value, ok := mu.mutation.Ref(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldRef,
		})
	}
	if value, ok := mu.mutation.Cron(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldCron,
		})
	}
	if value, ok := mu.mutation.PublicKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldPublicKey,
		})
	}
	if value, ok := mu.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldPrivateKey,
		})
	}
	if value, ok := mu.mutation.Passphrase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldPassphrase,
		})
	}
	if value, ok := mu.mutation.Commit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldCommit,
		})
	}
	if value, ok := mu.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: mirror.FieldLocked,
		})
	}
	if value, ok := mu.mutation.LastSync(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mirror.FieldLastSync,
		})
	}
	if mu.mutation.LastSyncCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: mirror.FieldLastSync,
		})
	}
	if mu.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirror.NamespaceTable,
			Columns: []string{mirror.NamespaceColumn},
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
	if nodes := mu.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirror.NamespaceTable,
			Columns: []string{mirror.NamespaceColumn},
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
	if mu.mutation.InodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   mirror.InodeTable,
			Columns: []string{mirror.InodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.InodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   mirror.InodeTable,
			Columns: []string{mirror.InodeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirror.ActivitiesTable,
			Columns: []string{mirror.ActivitiesColumn},
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
	if nodes := mu.mutation.RemovedActivitiesIDs(); len(nodes) > 0 && !mu.mutation.ActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirror.ActivitiesTable,
			Columns: []string{mirror.ActivitiesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirror.ActivitiesTable,
			Columns: []string{mirror.ActivitiesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mirror.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MirrorUpdateOne is the builder for updating a single Mirror entity.
type MirrorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MirrorMutation
}

// SetURL sets the "url" field.
func (muo *MirrorUpdateOne) SetURL(s string) *MirrorUpdateOne {
	muo.mutation.SetURL(s)
	return muo
}

// SetRef sets the "ref" field.
func (muo *MirrorUpdateOne) SetRef(s string) *MirrorUpdateOne {
	muo.mutation.SetRef(s)
	return muo
}

// SetCron sets the "cron" field.
func (muo *MirrorUpdateOne) SetCron(s string) *MirrorUpdateOne {
	muo.mutation.SetCron(s)
	return muo
}

// SetPublicKey sets the "public_key" field.
func (muo *MirrorUpdateOne) SetPublicKey(s string) *MirrorUpdateOne {
	muo.mutation.SetPublicKey(s)
	return muo
}

// SetPrivateKey sets the "private_key" field.
func (muo *MirrorUpdateOne) SetPrivateKey(s string) *MirrorUpdateOne {
	muo.mutation.SetPrivateKey(s)
	return muo
}

// SetPassphrase sets the "passphrase" field.
func (muo *MirrorUpdateOne) SetPassphrase(s string) *MirrorUpdateOne {
	muo.mutation.SetPassphrase(s)
	return muo
}

// SetCommit sets the "commit" field.
func (muo *MirrorUpdateOne) SetCommit(s string) *MirrorUpdateOne {
	muo.mutation.SetCommit(s)
	return muo
}

// SetLocked sets the "locked" field.
func (muo *MirrorUpdateOne) SetLocked(b bool) *MirrorUpdateOne {
	muo.mutation.SetLocked(b)
	return muo
}

// SetLastSync sets the "last_sync" field.
func (muo *MirrorUpdateOne) SetLastSync(t time.Time) *MirrorUpdateOne {
	muo.mutation.SetLastSync(t)
	return muo
}

// SetNillableLastSync sets the "last_sync" field if the given value is not nil.
func (muo *MirrorUpdateOne) SetNillableLastSync(t *time.Time) *MirrorUpdateOne {
	if t != nil {
		muo.SetLastSync(*t)
	}
	return muo
}

// ClearLastSync clears the value of the "last_sync" field.
func (muo *MirrorUpdateOne) ClearLastSync() *MirrorUpdateOne {
	muo.mutation.ClearLastSync()
	return muo
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (muo *MirrorUpdateOne) SetNamespaceID(id uuid.UUID) *MirrorUpdateOne {
	muo.mutation.SetNamespaceID(id)
	return muo
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (muo *MirrorUpdateOne) SetNamespace(n *Namespace) *MirrorUpdateOne {
	return muo.SetNamespaceID(n.ID)
}

// SetInodeID sets the "inode" edge to the Inode entity by ID.
func (muo *MirrorUpdateOne) SetInodeID(id uuid.UUID) *MirrorUpdateOne {
	muo.mutation.SetInodeID(id)
	return muo
}

// SetNillableInodeID sets the "inode" edge to the Inode entity by ID if the given value is not nil.
func (muo *MirrorUpdateOne) SetNillableInodeID(id *uuid.UUID) *MirrorUpdateOne {
	if id != nil {
		muo = muo.SetInodeID(*id)
	}
	return muo
}

// SetInode sets the "inode" edge to the Inode entity.
func (muo *MirrorUpdateOne) SetInode(i *Inode) *MirrorUpdateOne {
	return muo.SetInodeID(i.ID)
}

// AddActivityIDs adds the "activities" edge to the MirrorActivity entity by IDs.
func (muo *MirrorUpdateOne) AddActivityIDs(ids ...uuid.UUID) *MirrorUpdateOne {
	muo.mutation.AddActivityIDs(ids...)
	return muo
}

// AddActivities adds the "activities" edges to the MirrorActivity entity.
func (muo *MirrorUpdateOne) AddActivities(m ...*MirrorActivity) *MirrorUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddActivityIDs(ids...)
}

// Mutation returns the MirrorMutation object of the builder.
func (muo *MirrorUpdateOne) Mutation() *MirrorMutation {
	return muo.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (muo *MirrorUpdateOne) ClearNamespace() *MirrorUpdateOne {
	muo.mutation.ClearNamespace()
	return muo
}

// ClearInode clears the "inode" edge to the Inode entity.
func (muo *MirrorUpdateOne) ClearInode() *MirrorUpdateOne {
	muo.mutation.ClearInode()
	return muo
}

// ClearActivities clears all "activities" edges to the MirrorActivity entity.
func (muo *MirrorUpdateOne) ClearActivities() *MirrorUpdateOne {
	muo.mutation.ClearActivities()
	return muo
}

// RemoveActivityIDs removes the "activities" edge to MirrorActivity entities by IDs.
func (muo *MirrorUpdateOne) RemoveActivityIDs(ids ...uuid.UUID) *MirrorUpdateOne {
	muo.mutation.RemoveActivityIDs(ids...)
	return muo
}

// RemoveActivities removes "activities" edges to MirrorActivity entities.
func (muo *MirrorUpdateOne) RemoveActivities(m ...*MirrorActivity) *MirrorUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveActivityIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MirrorUpdateOne) Select(field string, fields ...string) *MirrorUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mirror entity.
func (muo *MirrorUpdateOne) Save(ctx context.Context) (*Mirror, error) {
	var (
		err  error
		node *Mirror
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MirrorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MirrorUpdateOne) SaveX(ctx context.Context) *Mirror {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MirrorUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MirrorUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MirrorUpdateOne) check() error {
	if _, ok := muo.mutation.NamespaceID(); muo.mutation.NamespaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mirror.namespace"`)
	}
	return nil
}

func (muo *MirrorUpdateOne) sqlSave(ctx context.Context) (_node *Mirror, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mirror.Table,
			Columns: mirror.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mirror.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mirror.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mirror.FieldID)
		for _, f := range fields {
			if !mirror.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mirror.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldURL,
		})
	}
	if value, ok := muo.mutation.Ref(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldRef,
		})
	}
	if value, ok := muo.mutation.Cron(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldCron,
		})
	}
	if value, ok := muo.mutation.PublicKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldPublicKey,
		})
	}
	if value, ok := muo.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldPrivateKey,
		})
	}
	if value, ok := muo.mutation.Passphrase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldPassphrase,
		})
	}
	if value, ok := muo.mutation.Commit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldCommit,
		})
	}
	if value, ok := muo.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: mirror.FieldLocked,
		})
	}
	if value, ok := muo.mutation.LastSync(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mirror.FieldLastSync,
		})
	}
	if muo.mutation.LastSyncCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: mirror.FieldLastSync,
		})
	}
	if muo.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirror.NamespaceTable,
			Columns: []string{mirror.NamespaceColumn},
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
	if nodes := muo.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirror.NamespaceTable,
			Columns: []string{mirror.NamespaceColumn},
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
	if muo.mutation.InodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   mirror.InodeTable,
			Columns: []string{mirror.InodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.InodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   mirror.InodeTable,
			Columns: []string{mirror.InodeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirror.ActivitiesTable,
			Columns: []string{mirror.ActivitiesColumn},
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
	if nodes := muo.mutation.RemovedActivitiesIDs(); len(nodes) > 0 && !muo.mutation.ActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirror.ActivitiesTable,
			Columns: []string{mirror.ActivitiesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirror.ActivitiesTable,
			Columns: []string{mirror.ActivitiesColumn},
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
	_node = &Mirror{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mirror.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

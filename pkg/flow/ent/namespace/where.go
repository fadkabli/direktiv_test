// Code generated by ent, DO NOT EDIT.

package namespace

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Config applies equality check predicate on the "config" field. It's identical to ConfigEQ.
func Config(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfig), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Namespace {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Namespace {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Namespace {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Namespace {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// ConfigEQ applies the EQ predicate on the "config" field.
func ConfigEQ(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfig), v))
	})
}

// ConfigNEQ applies the NEQ predicate on the "config" field.
func ConfigNEQ(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfig), v))
	})
}

// ConfigIn applies the In predicate on the "config" field.
func ConfigIn(vs ...string) predicate.Namespace {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldConfig), v...))
	})
}

// ConfigNotIn applies the NotIn predicate on the "config" field.
func ConfigNotIn(vs ...string) predicate.Namespace {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldConfig), v...))
	})
}

// ConfigGT applies the GT predicate on the "config" field.
func ConfigGT(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfig), v))
	})
}

// ConfigGTE applies the GTE predicate on the "config" field.
func ConfigGTE(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfig), v))
	})
}

// ConfigLT applies the LT predicate on the "config" field.
func ConfigLT(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfig), v))
	})
}

// ConfigLTE applies the LTE predicate on the "config" field.
func ConfigLTE(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfig), v))
	})
}

// ConfigContains applies the Contains predicate on the "config" field.
func ConfigContains(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldConfig), v))
	})
}

// ConfigHasPrefix applies the HasPrefix predicate on the "config" field.
func ConfigHasPrefix(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldConfig), v))
	})
}

// ConfigHasSuffix applies the HasSuffix predicate on the "config" field.
func ConfigHasSuffix(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldConfig), v))
	})
}

// ConfigEqualFold applies the EqualFold predicate on the "config" field.
func ConfigEqualFold(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldConfig), v))
	})
}

// ConfigContainsFold applies the ContainsFold predicate on the "config" field.
func ConfigContainsFold(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldConfig), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Namespace {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Namespace {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// HasInodes applies the HasEdge predicate on the "inodes" edge.
func HasInodes() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InodesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InodesTable, InodesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInodesWith applies the HasEdge predicate on the "inodes" edge with a given conditions (other predicates).
func HasInodesWith(preds ...predicate.Inode) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InodesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InodesTable, InodesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkflows applies the HasEdge predicate on the "workflows" edge.
func HasWorkflows() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkflowsTable, WorkflowsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowsWith applies the HasEdge predicate on the "workflows" edge with a given conditions (other predicates).
func HasWorkflowsWith(preds ...predicate.Workflow) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkflowsTable, WorkflowsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMirrors applies the HasEdge predicate on the "mirrors" edge.
func HasMirrors() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MirrorsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MirrorsTable, MirrorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMirrorsWith applies the HasEdge predicate on the "mirrors" edge with a given conditions (other predicates).
func HasMirrorsWith(preds ...predicate.Mirror) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MirrorsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MirrorsTable, MirrorsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMirrorActivities applies the HasEdge predicate on the "mirror_activities" edge.
func HasMirrorActivities() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MirrorActivitiesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MirrorActivitiesTable, MirrorActivitiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMirrorActivitiesWith applies the HasEdge predicate on the "mirror_activities" edge with a given conditions (other predicates).
func HasMirrorActivitiesWith(preds ...predicate.MirrorActivity) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MirrorActivitiesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MirrorActivitiesTable, MirrorActivitiesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstances applies the HasEdge predicate on the "instances" edge.
func HasInstances() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstancesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InstancesTable, InstancesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstancesWith applies the HasEdge predicate on the "instances" edge with a given conditions (other predicates).
func HasInstancesWith(preds ...predicate.Instance) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstancesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InstancesTable, InstancesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLogs applies the HasEdge predicate on the "logs" edge.
func HasLogs() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LogsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LogsTable, LogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLogsWith applies the HasEdge predicate on the "logs" edge with a given conditions (other predicates).
func HasLogsWith(preds ...predicate.LogMsg) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LogsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LogsTable, LogsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVars applies the HasEdge predicate on the "vars" edge.
func HasVars() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VarsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VarsTable, VarsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVarsWith applies the HasEdge predicate on the "vars" edge with a given conditions (other predicates).
func HasVarsWith(preds ...predicate.VarRef) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VarsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VarsTable, VarsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCloudevents applies the HasEdge predicate on the "cloudevents" edge.
func HasCloudevents() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CloudeventsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CloudeventsTable, CloudeventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCloudeventsWith applies the HasEdge predicate on the "cloudevents" edge with a given conditions (other predicates).
func HasCloudeventsWith(preds ...predicate.CloudEvents) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CloudeventsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CloudeventsTable, CloudeventsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNamespacelisteners applies the HasEdge predicate on the "namespacelisteners" edge.
func HasNamespacelisteners() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespacelistenersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NamespacelistenersTable, NamespacelistenersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespacelistenersWith applies the HasEdge predicate on the "namespacelisteners" edge with a given conditions (other predicates).
func HasNamespacelistenersWith(preds ...predicate.Events) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespacelistenersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NamespacelistenersTable, NamespacelistenersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnnotations applies the HasEdge predicate on the "annotations" edge.
func HasAnnotations() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnnotationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnnotationsTable, AnnotationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnnotationsWith applies the HasEdge predicate on the "annotations" edge with a given conditions (other predicates).
func HasAnnotationsWith(preds ...predicate.Annotation) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnnotationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnnotationsTable, AnnotationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasServices applies the HasEdge predicate on the "services" edge.
func HasServices() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ServicesTable, ServicesFieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ServicesTable, ServicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasServicesWith applies the HasEdge predicate on the "services" edge with a given conditions (other predicates).
func HasServicesWith(preds ...predicate.Services) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ServicesInverseTable, ServicesFieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ServicesTable, ServicesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Namespace) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Namespace) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Namespace) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		p(s.Not())
	})
}

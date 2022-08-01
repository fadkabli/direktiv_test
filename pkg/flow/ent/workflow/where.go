// Code generated by ent, DO NOT EDIT.

package workflow

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Live applies equality check predicate on the "live" field. It's identical to LiveEQ.
func Live(v bool) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLive), v))
	})
}

// LogToEvents applies equality check predicate on the "logToEvents" field. It's identical to LogToEventsEQ.
func LogToEvents(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogToEvents), v))
	})
}

// ReadOnly applies equality check predicate on the "readOnly" field. It's identical to ReadOnlyEQ.
func ReadOnly(v bool) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadOnly), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// LiveEQ applies the EQ predicate on the "live" field.
func LiveEQ(v bool) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLive), v))
	})
}

// LiveNEQ applies the NEQ predicate on the "live" field.
func LiveNEQ(v bool) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLive), v))
	})
}

// LogToEventsEQ applies the EQ predicate on the "logToEvents" field.
func LogToEventsEQ(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsNEQ applies the NEQ predicate on the "logToEvents" field.
func LogToEventsNEQ(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsIn applies the In predicate on the "logToEvents" field.
func LogToEventsIn(vs ...string) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogToEvents), v...))
	})
}

// LogToEventsNotIn applies the NotIn predicate on the "logToEvents" field.
func LogToEventsNotIn(vs ...string) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogToEvents), v...))
	})
}

// LogToEventsGT applies the GT predicate on the "logToEvents" field.
func LogToEventsGT(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsGTE applies the GTE predicate on the "logToEvents" field.
func LogToEventsGTE(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsLT applies the LT predicate on the "logToEvents" field.
func LogToEventsLT(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsLTE applies the LTE predicate on the "logToEvents" field.
func LogToEventsLTE(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsContains applies the Contains predicate on the "logToEvents" field.
func LogToEventsContains(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsHasPrefix applies the HasPrefix predicate on the "logToEvents" field.
func LogToEventsHasPrefix(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsHasSuffix applies the HasSuffix predicate on the "logToEvents" field.
func LogToEventsHasSuffix(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsIsNil applies the IsNil predicate on the "logToEvents" field.
func LogToEventsIsNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLogToEvents)))
	})
}

// LogToEventsNotNil applies the NotNil predicate on the "logToEvents" field.
func LogToEventsNotNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLogToEvents)))
	})
}

// LogToEventsEqualFold applies the EqualFold predicate on the "logToEvents" field.
func LogToEventsEqualFold(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsContainsFold applies the ContainsFold predicate on the "logToEvents" field.
func LogToEventsContainsFold(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogToEvents), v))
	})
}

// ReadOnlyEQ applies the EQ predicate on the "readOnly" field.
func ReadOnlyEQ(v bool) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadOnly), v))
	})
}

// ReadOnlyNEQ applies the NEQ predicate on the "readOnly" field.
func ReadOnlyNEQ(v bool) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReadOnly), v))
	})
}

// ReadOnlyIsNil applies the IsNil predicate on the "readOnly" field.
func ReadOnlyIsNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldReadOnly)))
	})
}

// ReadOnlyNotNil applies the NotNil predicate on the "readOnly" field.
func ReadOnlyNotNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldReadOnly)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// HasInode applies the HasEdge predicate on the "inode" edge.
func HasInode() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InodeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, InodeTable, InodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInodeWith applies the HasEdge predicate on the "inode" edge with a given conditions (other predicates).
func HasInodeWith(preds ...predicate.Inode) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InodeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, InodeTable, InodeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRevisions applies the HasEdge predicate on the "revisions" edge.
func HasRevisions() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RevisionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RevisionsTable, RevisionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRevisionsWith applies the HasEdge predicate on the "revisions" edge with a given conditions (other predicates).
func HasRevisionsWith(preds ...predicate.Revision) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RevisionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RevisionsTable, RevisionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRefs applies the HasEdge predicate on the "refs" edge.
func HasRefs() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RefsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RefsTable, RefsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRefsWith applies the HasEdge predicate on the "refs" edge with a given conditions (other predicates).
func HasRefsWith(preds ...predicate.Ref) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RefsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RefsTable, RefsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstances applies the HasEdge predicate on the "instances" edge.
func HasInstances() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstancesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InstancesTable, InstancesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstancesWith applies the HasEdge predicate on the "instances" edge with a given conditions (other predicates).
func HasInstancesWith(preds ...predicate.Instance) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
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

// HasRoutes applies the HasEdge predicate on the "routes" edge.
func HasRoutes() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoutesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoutesTable, RoutesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoutesWith applies the HasEdge predicate on the "routes" edge with a given conditions (other predicates).
func HasRoutesWith(preds ...predicate.Route) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoutesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoutesTable, RoutesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLogs applies the HasEdge predicate on the "logs" edge.
func HasLogs() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LogsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LogsTable, LogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLogsWith applies the HasEdge predicate on the "logs" edge with a given conditions (other predicates).
func HasLogsWith(preds ...predicate.LogMsg) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
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
func HasVars() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VarsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VarsTable, VarsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVarsWith applies the HasEdge predicate on the "vars" edge with a given conditions (other predicates).
func HasVarsWith(preds ...predicate.VarRef) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
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

// HasWfevents applies the HasEdge predicate on the "wfevents" edge.
func HasWfevents() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WfeventsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WfeventsTable, WfeventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWfeventsWith applies the HasEdge predicate on the "wfevents" edge with a given conditions (other predicates).
func HasWfeventsWith(preds ...predicate.Events) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WfeventsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WfeventsTable, WfeventsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
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
func Not(p predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		p(s.Not())
	})
}

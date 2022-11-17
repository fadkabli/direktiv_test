// Code generated by ent, DO NOT EDIT.

package namespacesecret

import (
	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/pkg/secrets/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Ns applies equality check predicate on the "ns" field. It's identical to NsEQ.
func Ns(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNs), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Secret applies equality check predicate on the "secret" field. It's identical to SecretEQ.
func Secret(v []byte) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// NsEQ applies the EQ predicate on the "ns" field.
func NsEQ(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNs), v))
	})
}

// NsNEQ applies the NEQ predicate on the "ns" field.
func NsNEQ(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNs), v))
	})
}

// NsIn applies the In predicate on the "ns" field.
func NsIn(vs ...string) predicate.NamespaceSecret {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNs), v...))
	})
}

// NsNotIn applies the NotIn predicate on the "ns" field.
func NsNotIn(vs ...string) predicate.NamespaceSecret {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNs), v...))
	})
}

// NsGT applies the GT predicate on the "ns" field.
func NsGT(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNs), v))
	})
}

// NsGTE applies the GTE predicate on the "ns" field.
func NsGTE(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNs), v))
	})
}

// NsLT applies the LT predicate on the "ns" field.
func NsLT(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNs), v))
	})
}

// NsLTE applies the LTE predicate on the "ns" field.
func NsLTE(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNs), v))
	})
}

// NsContains applies the Contains predicate on the "ns" field.
func NsContains(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNs), v))
	})
}

// NsHasPrefix applies the HasPrefix predicate on the "ns" field.
func NsHasPrefix(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNs), v))
	})
}

// NsHasSuffix applies the HasSuffix predicate on the "ns" field.
func NsHasSuffix(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNs), v))
	})
}

// NsEqualFold applies the EqualFold predicate on the "ns" field.
func NsEqualFold(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNs), v))
	})
}

// NsContainsFold applies the ContainsFold predicate on the "ns" field.
func NsContainsFold(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNs), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.NamespaceSecret {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.NamespaceSecret {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SecretEQ applies the EQ predicate on the "secret" field.
func SecretEQ(v []byte) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// SecretNEQ applies the NEQ predicate on the "secret" field.
func SecretNEQ(v []byte) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecret), v))
	})
}

// SecretIn applies the In predicate on the "secret" field.
func SecretIn(vs ...[]byte) predicate.NamespaceSecret {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSecret), v...))
	})
}

// SecretNotIn applies the NotIn predicate on the "secret" field.
func SecretNotIn(vs ...[]byte) predicate.NamespaceSecret {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSecret), v...))
	})
}

// SecretGT applies the GT predicate on the "secret" field.
func SecretGT(v []byte) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecret), v))
	})
}

// SecretGTE applies the GTE predicate on the "secret" field.
func SecretGTE(v []byte) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecret), v))
	})
}

// SecretLT applies the LT predicate on the "secret" field.
func SecretLT(v []byte) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecret), v))
	})
}

// SecretLTE applies the LTE predicate on the "secret" field.
func SecretLTE(v []byte) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecret), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NamespaceSecret) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NamespaceSecret) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
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
func Not(p predicate.NamespaceSecret) predicate.NamespaceSecret {
	return predicate.NamespaceSecret(func(s *sql.Selector) {
		p(s.Not())
	})
}

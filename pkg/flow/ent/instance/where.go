// Code generated by entc, DO NOT EDIT.

package instance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// As applies equality check predicate on the "as" field. It's identical to AsEQ.
func As(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAs), v))
	})
}

// ErrorCode applies equality check predicate on the "errorCode" field. It's identical to ErrorCodeEQ.
func ErrorCode(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrorCode), v))
	})
}

// ErrorMessage applies equality check predicate on the "errorMessage" field. It's identical to ErrorMessageEQ.
func ErrorMessage(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrorMessage), v))
	})
}

// Invoker applies equality check predicate on the "invoker" field. It's identical to InvokerEQ.
func Invoker(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvoker), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndAt), v))
	})
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndAt), v...))
	})
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndAt), v...))
	})
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndAt), v))
	})
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndAt), v))
	})
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndAt), v))
	})
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndAt), v))
	})
}

// EndAtIsNil applies the IsNil predicate on the "end_at" field.
func EndAtIsNil() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndAt)))
	})
}

// EndAtNotNil applies the NotNil predicate on the "end_at" field.
func EndAtNotNil() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndAt)))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// AsEQ applies the EQ predicate on the "as" field.
func AsEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAs), v))
	})
}

// AsNEQ applies the NEQ predicate on the "as" field.
func AsNEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAs), v))
	})
}

// AsIn applies the In predicate on the "as" field.
func AsIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAs), v...))
	})
}

// AsNotIn applies the NotIn predicate on the "as" field.
func AsNotIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAs), v...))
	})
}

// AsGT applies the GT predicate on the "as" field.
func AsGT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAs), v))
	})
}

// AsGTE applies the GTE predicate on the "as" field.
func AsGTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAs), v))
	})
}

// AsLT applies the LT predicate on the "as" field.
func AsLT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAs), v))
	})
}

// AsLTE applies the LTE predicate on the "as" field.
func AsLTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAs), v))
	})
}

// AsContains applies the Contains predicate on the "as" field.
func AsContains(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAs), v))
	})
}

// AsHasPrefix applies the HasPrefix predicate on the "as" field.
func AsHasPrefix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAs), v))
	})
}

// AsHasSuffix applies the HasSuffix predicate on the "as" field.
func AsHasSuffix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAs), v))
	})
}

// AsEqualFold applies the EqualFold predicate on the "as" field.
func AsEqualFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAs), v))
	})
}

// AsContainsFold applies the ContainsFold predicate on the "as" field.
func AsContainsFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAs), v))
	})
}

// ErrorCodeEQ applies the EQ predicate on the "errorCode" field.
func ErrorCodeEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeNEQ applies the NEQ predicate on the "errorCode" field.
func ErrorCodeNEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeIn applies the In predicate on the "errorCode" field.
func ErrorCodeIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldErrorCode), v...))
	})
}

// ErrorCodeNotIn applies the NotIn predicate on the "errorCode" field.
func ErrorCodeNotIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldErrorCode), v...))
	})
}

// ErrorCodeGT applies the GT predicate on the "errorCode" field.
func ErrorCodeGT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeGTE applies the GTE predicate on the "errorCode" field.
func ErrorCodeGTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeLT applies the LT predicate on the "errorCode" field.
func ErrorCodeLT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeLTE applies the LTE predicate on the "errorCode" field.
func ErrorCodeLTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeContains applies the Contains predicate on the "errorCode" field.
func ErrorCodeContains(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeHasPrefix applies the HasPrefix predicate on the "errorCode" field.
func ErrorCodeHasPrefix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeHasSuffix applies the HasSuffix predicate on the "errorCode" field.
func ErrorCodeHasSuffix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeIsNil applies the IsNil predicate on the "errorCode" field.
func ErrorCodeIsNil() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldErrorCode)))
	})
}

// ErrorCodeNotNil applies the NotNil predicate on the "errorCode" field.
func ErrorCodeNotNil() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldErrorCode)))
	})
}

// ErrorCodeEqualFold applies the EqualFold predicate on the "errorCode" field.
func ErrorCodeEqualFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldErrorCode), v))
	})
}

// ErrorCodeContainsFold applies the ContainsFold predicate on the "errorCode" field.
func ErrorCodeContainsFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldErrorCode), v))
	})
}

// ErrorMessageEQ applies the EQ predicate on the "errorMessage" field.
func ErrorMessageEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageNEQ applies the NEQ predicate on the "errorMessage" field.
func ErrorMessageNEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageIn applies the In predicate on the "errorMessage" field.
func ErrorMessageIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldErrorMessage), v...))
	})
}

// ErrorMessageNotIn applies the NotIn predicate on the "errorMessage" field.
func ErrorMessageNotIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldErrorMessage), v...))
	})
}

// ErrorMessageGT applies the GT predicate on the "errorMessage" field.
func ErrorMessageGT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageGTE applies the GTE predicate on the "errorMessage" field.
func ErrorMessageGTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageLT applies the LT predicate on the "errorMessage" field.
func ErrorMessageLT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageLTE applies the LTE predicate on the "errorMessage" field.
func ErrorMessageLTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageContains applies the Contains predicate on the "errorMessage" field.
func ErrorMessageContains(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageHasPrefix applies the HasPrefix predicate on the "errorMessage" field.
func ErrorMessageHasPrefix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageHasSuffix applies the HasSuffix predicate on the "errorMessage" field.
func ErrorMessageHasSuffix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageIsNil applies the IsNil predicate on the "errorMessage" field.
func ErrorMessageIsNil() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldErrorMessage)))
	})
}

// ErrorMessageNotNil applies the NotNil predicate on the "errorMessage" field.
func ErrorMessageNotNil() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldErrorMessage)))
	})
}

// ErrorMessageEqualFold applies the EqualFold predicate on the "errorMessage" field.
func ErrorMessageEqualFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldErrorMessage), v))
	})
}

// ErrorMessageContainsFold applies the ContainsFold predicate on the "errorMessage" field.
func ErrorMessageContainsFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldErrorMessage), v))
	})
}

// InvokerEQ applies the EQ predicate on the "invoker" field.
func InvokerEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvoker), v))
	})
}

// InvokerNEQ applies the NEQ predicate on the "invoker" field.
func InvokerNEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInvoker), v))
	})
}

// InvokerIn applies the In predicate on the "invoker" field.
func InvokerIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInvoker), v...))
	})
}

// InvokerNotIn applies the NotIn predicate on the "invoker" field.
func InvokerNotIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInvoker), v...))
	})
}

// InvokerGT applies the GT predicate on the "invoker" field.
func InvokerGT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInvoker), v))
	})
}

// InvokerGTE applies the GTE predicate on the "invoker" field.
func InvokerGTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInvoker), v))
	})
}

// InvokerLT applies the LT predicate on the "invoker" field.
func InvokerLT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInvoker), v))
	})
}

// InvokerLTE applies the LTE predicate on the "invoker" field.
func InvokerLTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInvoker), v))
	})
}

// InvokerContains applies the Contains predicate on the "invoker" field.
func InvokerContains(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInvoker), v))
	})
}

// InvokerHasPrefix applies the HasPrefix predicate on the "invoker" field.
func InvokerHasPrefix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInvoker), v))
	})
}

// InvokerHasSuffix applies the HasSuffix predicate on the "invoker" field.
func InvokerHasSuffix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInvoker), v))
	})
}

// InvokerIsNil applies the IsNil predicate on the "invoker" field.
func InvokerIsNil() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInvoker)))
	})
}

// InvokerNotNil applies the NotNil predicate on the "invoker" field.
func InvokerNotNil() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInvoker)))
	})
}

// InvokerEqualFold applies the EqualFold predicate on the "invoker" field.
func InvokerEqualFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInvoker), v))
	})
}

// InvokerContainsFold applies the ContainsFold predicate on the "invoker" field.
func InvokerContainsFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInvoker), v))
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
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

// HasWorkflow applies the HasEdge predicate on the "workflow" edge.
func HasWorkflow() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowWith applies the HasEdge predicate on the "workflow" edge with a given conditions (other predicates).
func HasWorkflowWith(preds ...predicate.Workflow) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRevision applies the HasEdge predicate on the "revision" edge.
func HasRevision() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RevisionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RevisionTable, RevisionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRevisionWith applies the HasEdge predicate on the "revision" edge with a given conditions (other predicates).
func HasRevisionWith(preds ...predicate.Revision) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RevisionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RevisionTable, RevisionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLogs applies the HasEdge predicate on the "logs" edge.
func HasLogs() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LogsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LogsTable, LogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLogsWith applies the HasEdge predicate on the "logs" edge with a given conditions (other predicates).
func HasLogsWith(preds ...predicate.LogMsg) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
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
func HasVars() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VarsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VarsTable, VarsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVarsWith applies the HasEdge predicate on the "vars" edge with a given conditions (other predicates).
func HasVarsWith(preds ...predicate.VarRef) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
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

// HasRuntime applies the HasEdge predicate on the "runtime" edge.
func HasRuntime() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RuntimeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, RuntimeTable, RuntimeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRuntimeWith applies the HasEdge predicate on the "runtime" edge with a given conditions (other predicates).
func HasRuntimeWith(preds ...predicate.InstanceRuntime) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RuntimeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, RuntimeTable, RuntimeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildrenTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.InstanceRuntime) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildrenInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEventlisteners applies the HasEdge predicate on the "eventlisteners" edge.
func HasEventlisteners() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventlistenersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventlistenersTable, EventlistenersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventlistenersWith applies the HasEdge predicate on the "eventlisteners" edge with a given conditions (other predicates).
func HasEventlistenersWith(preds ...predicate.Events) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventlistenersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventlistenersTable, EventlistenersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Instance) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Instance) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
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
func Not(p predicate.Instance) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		p(s.Not())
	})
}

// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/vorteil/direktiv/ent"
)

// The CloudEventsFunc type is an adapter to allow the use of ordinary
// function as CloudEvents mutator.
type CloudEventsFunc func(context.Context, *ent.CloudEventsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CloudEventsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CloudEventsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CloudEventsMutation", m)
	}
	return f(ctx, mv)
}

// The NamespaceFunc type is an adapter to allow the use of ordinary
// function as Namespace mutator.
type NamespaceFunc func(context.Context, *ent.NamespaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NamespaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.NamespaceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NamespaceMutation", m)
	}
	return f(ctx, mv)
}

// The ServicesFunc type is an adapter to allow the use of ordinary
// function as Services mutator.
type ServicesFunc func(context.Context, *ent.ServicesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ServicesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ServicesMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ServicesMutation", m)
	}
	return f(ctx, mv)
}

// The WorkflowFunc type is an adapter to allow the use of ordinary
// function as Workflow mutator.
type WorkflowFunc func(context.Context, *ent.WorkflowMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkflowFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WorkflowMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkflowMutation", m)
	}
	return f(ctx, mv)
}

// The WorkflowEventsFunc type is an adapter to allow the use of ordinary
// function as WorkflowEvents mutator.
type WorkflowEventsFunc func(context.Context, *ent.WorkflowEventsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkflowEventsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WorkflowEventsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkflowEventsMutation", m)
	}
	return f(ctx, mv)
}

// The WorkflowEventsWaitFunc type is an adapter to allow the use of ordinary
// function as WorkflowEventsWait mutator.
type WorkflowEventsWaitFunc func(context.Context, *ent.WorkflowEventsWaitMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkflowEventsWaitFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WorkflowEventsWaitMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkflowEventsWaitMutation", m)
	}
	return f(ctx, mv)
}

// The WorkflowInstanceFunc type is an adapter to allow the use of ordinary
// function as WorkflowInstance mutator.
type WorkflowInstanceFunc func(context.Context, *ent.WorkflowInstanceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkflowInstanceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WorkflowInstanceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WorkflowInstanceMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}

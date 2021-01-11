// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/team15/app/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The CleanerNameQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CleanerNameQueryRuleFunc func(context.Context, *ent.CleanerNameQuery) error

// EvalQuery return f(ctx, q).
func (f CleanerNameQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CleanerNameQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CleanerNameQuery", q)
}

// The CleanerNameMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CleanerNameMutationRuleFunc func(context.Context, *ent.CleanerNameMutation) error

// EvalMutation calls f(ctx, m).
func (f CleanerNameMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CleanerNameMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CleanerNameMutation", m)
}

// The CleaningRoomQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CleaningRoomQueryRuleFunc func(context.Context, *ent.CleaningRoomQuery) error

// EvalQuery return f(ctx, q).
func (f CleaningRoomQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CleaningRoomQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CleaningRoomQuery", q)
}

// The CleaningRoomMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CleaningRoomMutationRuleFunc func(context.Context, *ent.CleaningRoomMutation) error

// EvalMutation calls f(ctx, m).
func (f CleaningRoomMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CleaningRoomMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CleaningRoomMutation", m)
}

// The DepositQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DepositQueryRuleFunc func(context.Context, *ent.DepositQuery) error

// EvalQuery return f(ctx, q).
func (f DepositQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DepositQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DepositQuery", q)
}

// The DepositMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DepositMutationRuleFunc func(context.Context, *ent.DepositMutation) error

// EvalMutation calls f(ctx, m).
func (f DepositMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DepositMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DepositMutation", m)
}

// The EmployeeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EmployeeQueryRuleFunc func(context.Context, *ent.EmployeeQuery) error

// EvalQuery return f(ctx, q).
func (f EmployeeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EmployeeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EmployeeQuery", q)
}

// The EmployeeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EmployeeMutationRuleFunc func(context.Context, *ent.EmployeeMutation) error

// EvalMutation calls f(ctx, m).
func (f EmployeeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EmployeeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EmployeeMutation", m)
}

// The EquipmentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EquipmentQueryRuleFunc func(context.Context, *ent.EquipmentQuery) error

// EvalQuery return f(ctx, q).
func (f EquipmentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EquipmentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EquipmentQuery", q)
}

// The EquipmentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EquipmentMutationRuleFunc func(context.Context, *ent.EquipmentMutation) error

// EvalMutation calls f(ctx, m).
func (f EquipmentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EquipmentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EquipmentMutation", m)
}

// The FacilitieQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FacilitieQueryRuleFunc func(context.Context, *ent.FacilitieQuery) error

// EvalQuery return f(ctx, q).
func (f FacilitieQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FacilitieQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FacilitieQuery", q)
}

// The FacilitieMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FacilitieMutationRuleFunc func(context.Context, *ent.FacilitieMutation) error

// EvalMutation calls f(ctx, m).
func (f FacilitieMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FacilitieMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FacilitieMutation", m)
}

// The LengthTimeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LengthTimeQueryRuleFunc func(context.Context, *ent.LengthTimeQuery) error

// EvalQuery return f(ctx, q).
func (f LengthTimeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LengthTimeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LengthTimeQuery", q)
}

// The LengthTimeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LengthTimeMutationRuleFunc func(context.Context, *ent.LengthTimeMutation) error

// EvalMutation calls f(ctx, m).
func (f LengthTimeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LengthTimeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LengthTimeMutation", m)
}

// The NearbyplaceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type NearbyplaceQueryRuleFunc func(context.Context, *ent.NearbyplaceQuery) error

// EvalQuery return f(ctx, q).
func (f NearbyplaceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NearbyplaceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.NearbyplaceQuery", q)
}

// The NearbyplaceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type NearbyplaceMutationRuleFunc func(context.Context, *ent.NearbyplaceMutation) error

// EvalMutation calls f(ctx, m).
func (f NearbyplaceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.NearbyplaceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.NearbyplaceMutation", m)
}

// The QuantityQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type QuantityQueryRuleFunc func(context.Context, *ent.QuantityQuery) error

// EvalQuery return f(ctx, q).
func (f QuantityQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.QuantityQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.QuantityQuery", q)
}

// The QuantityMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type QuantityMutationRuleFunc func(context.Context, *ent.QuantityMutation) error

// EvalMutation calls f(ctx, m).
func (f QuantityMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.QuantityMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.QuantityMutation", m)
}

// The RoomdetailQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RoomdetailQueryRuleFunc func(context.Context, *ent.RoomdetailQuery) error

// EvalQuery return f(ctx, q).
func (f RoomdetailQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RoomdetailQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RoomdetailQuery", q)
}

// The RoomdetailMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RoomdetailMutationRuleFunc func(context.Context, *ent.RoomdetailMutation) error

// EvalMutation calls f(ctx, m).
func (f RoomdetailMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RoomdetailMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RoomdetailMutation", m)
}

// The StatusdQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StatusdQueryRuleFunc func(context.Context, *ent.StatusdQuery) error

// EvalQuery return f(ctx, q).
func (f StatusdQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StatusdQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StatusdQuery", q)
}

// The StatusdMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StatusdMutationRuleFunc func(context.Context, *ent.StatusdMutation) error

// EvalMutation calls f(ctx, m).
func (f StatusdMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StatusdMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StatusdMutation", m)
}

// The StaytypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StaytypeQueryRuleFunc func(context.Context, *ent.StaytypeQuery) error

// EvalQuery return f(ctx, q).
func (f StaytypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StaytypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StaytypeQuery", q)
}

// The StaytypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StaytypeMutationRuleFunc func(context.Context, *ent.StaytypeMutation) error

// EvalMutation calls f(ctx, m).
func (f StaytypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StaytypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StaytypeMutation", m)
}

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

// The BedtypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BedtypeQueryRuleFunc func(context.Context, *ent.BedtypeQuery) error

// EvalQuery return f(ctx, q).
func (f BedtypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BedtypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BedtypeQuery", q)
}

// The BedtypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BedtypeMutationRuleFunc func(context.Context, *ent.BedtypeMutation) error

// EvalMutation calls f(ctx, m).
func (f BedtypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BedtypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BedtypeMutation", m)
}

// The BillQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BillQueryRuleFunc func(context.Context, *ent.BillQuery) error

// EvalQuery return f(ctx, q).
func (f BillQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BillQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BillQuery", q)
}

// The BillMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BillMutationRuleFunc func(context.Context, *ent.BillMutation) error

// EvalMutation calls f(ctx, m).
func (f BillMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BillMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BillMutation", m)
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

// The JobpositionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type JobpositionQueryRuleFunc func(context.Context, *ent.JobpositionQuery) error

// EvalQuery return f(ctx, q).
func (f JobpositionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.JobpositionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.JobpositionQuery", q)
}

// The JobpositionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type JobpositionMutationRuleFunc func(context.Context, *ent.JobpositionMutation) error

// EvalMutation calls f(ctx, m).
func (f JobpositionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.JobpositionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.JobpositionMutation", m)
}

// The LeaseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LeaseQueryRuleFunc func(context.Context, *ent.LeaseQuery) error

// EvalQuery return f(ctx, q).
func (f LeaseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LeaseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LeaseQuery", q)
}

// The LeaseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LeaseMutationRuleFunc func(context.Context, *ent.LeaseMutation) error

// EvalMutation calls f(ctx, m).
func (f LeaseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LeaseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LeaseMutation", m)
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

// The PaymentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PaymentQueryRuleFunc func(context.Context, *ent.PaymentQuery) error

// EvalQuery return f(ctx, q).
func (f PaymentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PaymentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PaymentQuery", q)
}

// The PaymentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PaymentMutationRuleFunc func(context.Context, *ent.PaymentMutation) error

// EvalMutation calls f(ctx, m).
func (f PaymentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PaymentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PaymentMutation", m)
}

// The PetruleQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PetruleQueryRuleFunc func(context.Context, *ent.PetruleQuery) error

// EvalQuery return f(ctx, q).
func (f PetruleQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PetruleQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PetruleQuery", q)
}

// The PetruleMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PetruleMutationRuleFunc func(context.Context, *ent.PetruleMutation) error

// EvalMutation calls f(ctx, m).
func (f PetruleMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PetruleMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PetruleMutation", m)
}

// The PledgeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PledgeQueryRuleFunc func(context.Context, *ent.PledgeQuery) error

// EvalQuery return f(ctx, q).
func (f PledgeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PledgeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PledgeQuery", q)
}

// The PledgeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PledgeMutationRuleFunc func(context.Context, *ent.PledgeMutation) error

// EvalMutation calls f(ctx, m).
func (f PledgeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PledgeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PledgeMutation", m)
}

// The RentalstatusQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RentalstatusQueryRuleFunc func(context.Context, *ent.RentalstatusQuery) error

// EvalQuery return f(ctx, q).
func (f RentalstatusQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RentalstatusQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RentalstatusQuery", q)
}

// The RentalstatusMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RentalstatusMutationRuleFunc func(context.Context, *ent.RentalstatusMutation) error

// EvalMutation calls f(ctx, m).
func (f RentalstatusMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RentalstatusMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RentalstatusMutation", m)
}

// The RepairinvoiceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RepairinvoiceQueryRuleFunc func(context.Context, *ent.RepairinvoiceQuery) error

// EvalQuery return f(ctx, q).
func (f RepairinvoiceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RepairinvoiceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RepairinvoiceQuery", q)
}

// The RepairinvoiceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RepairinvoiceMutationRuleFunc func(context.Context, *ent.RepairinvoiceMutation) error

// EvalMutation calls f(ctx, m).
func (f RepairinvoiceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RepairinvoiceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RepairinvoiceMutation", m)
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

// The SituationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SituationQueryRuleFunc func(context.Context, *ent.SituationQuery) error

// EvalQuery return f(ctx, q).
func (f SituationQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SituationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SituationQuery", q)
}

// The SituationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SituationMutationRuleFunc func(context.Context, *ent.SituationMutation) error

// EvalMutation calls f(ctx, m).
func (f SituationMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SituationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SituationMutation", m)
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

// The WifiQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type WifiQueryRuleFunc func(context.Context, *ent.WifiQuery) error

// EvalQuery return f(ctx, q).
func (f WifiQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.WifiQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.WifiQuery", q)
}

// The WifiMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type WifiMutationRuleFunc func(context.Context, *ent.WifiMutation) error

// EvalMutation calls f(ctx, m).
func (f WifiMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.WifiMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.WifiMutation", m)
}

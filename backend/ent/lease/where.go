// Code generated by entc, DO NOT EDIT.

package lease

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team15/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Addedtime applies equality check predicate on the "addedtime" field. It's identical to AddedtimeEQ.
func Addedtime(v time.Time) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedtime), v))
	})
}

// Tenant applies equality check predicate on the "tenant" field. It's identical to TenantEQ.
func Tenant(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenant), v))
	})
}

// Numbtenant applies equality check predicate on the "numbtenant" field. It's identical to NumbtenantEQ.
func Numbtenant(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumbtenant), v))
	})
}

// Pettenant applies equality check predicate on the "pettenant" field. It's identical to PettenantEQ.
func Pettenant(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPettenant), v))
	})
}

// AddedtimeEQ applies the EQ predicate on the "addedtime" field.
func AddedtimeEQ(v time.Time) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedtime), v))
	})
}

// AddedtimeNEQ applies the NEQ predicate on the "addedtime" field.
func AddedtimeNEQ(v time.Time) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddedtime), v))
	})
}

// AddedtimeIn applies the In predicate on the "addedtime" field.
func AddedtimeIn(vs ...time.Time) predicate.Lease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddedtime), v...))
	})
}

// AddedtimeNotIn applies the NotIn predicate on the "addedtime" field.
func AddedtimeNotIn(vs ...time.Time) predicate.Lease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddedtime), v...))
	})
}

// AddedtimeGT applies the GT predicate on the "addedtime" field.
func AddedtimeGT(v time.Time) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddedtime), v))
	})
}

// AddedtimeGTE applies the GTE predicate on the "addedtime" field.
func AddedtimeGTE(v time.Time) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddedtime), v))
	})
}

// AddedtimeLT applies the LT predicate on the "addedtime" field.
func AddedtimeLT(v time.Time) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddedtime), v))
	})
}

// AddedtimeLTE applies the LTE predicate on the "addedtime" field.
func AddedtimeLTE(v time.Time) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddedtime), v))
	})
}

// TenantEQ applies the EQ predicate on the "tenant" field.
func TenantEQ(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenant), v))
	})
}

// TenantNEQ applies the NEQ predicate on the "tenant" field.
func TenantNEQ(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenant), v))
	})
}

// TenantIn applies the In predicate on the "tenant" field.
func TenantIn(vs ...string) predicate.Lease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTenant), v...))
	})
}

// TenantNotIn applies the NotIn predicate on the "tenant" field.
func TenantNotIn(vs ...string) predicate.Lease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTenant), v...))
	})
}

// TenantGT applies the GT predicate on the "tenant" field.
func TenantGT(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenant), v))
	})
}

// TenantGTE applies the GTE predicate on the "tenant" field.
func TenantGTE(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenant), v))
	})
}

// TenantLT applies the LT predicate on the "tenant" field.
func TenantLT(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenant), v))
	})
}

// TenantLTE applies the LTE predicate on the "tenant" field.
func TenantLTE(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenant), v))
	})
}

// TenantContains applies the Contains predicate on the "tenant" field.
func TenantContains(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTenant), v))
	})
}

// TenantHasPrefix applies the HasPrefix predicate on the "tenant" field.
func TenantHasPrefix(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTenant), v))
	})
}

// TenantHasSuffix applies the HasSuffix predicate on the "tenant" field.
func TenantHasSuffix(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTenant), v))
	})
}

// TenantEqualFold applies the EqualFold predicate on the "tenant" field.
func TenantEqualFold(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTenant), v))
	})
}

// TenantContainsFold applies the ContainsFold predicate on the "tenant" field.
func TenantContainsFold(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTenant), v))
	})
}

// NumbtenantEQ applies the EQ predicate on the "numbtenant" field.
func NumbtenantEQ(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantNEQ applies the NEQ predicate on the "numbtenant" field.
func NumbtenantNEQ(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantIn applies the In predicate on the "numbtenant" field.
func NumbtenantIn(vs ...string) predicate.Lease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNumbtenant), v...))
	})
}

// NumbtenantNotIn applies the NotIn predicate on the "numbtenant" field.
func NumbtenantNotIn(vs ...string) predicate.Lease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNumbtenant), v...))
	})
}

// NumbtenantGT applies the GT predicate on the "numbtenant" field.
func NumbtenantGT(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantGTE applies the GTE predicate on the "numbtenant" field.
func NumbtenantGTE(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantLT applies the LT predicate on the "numbtenant" field.
func NumbtenantLT(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantLTE applies the LTE predicate on the "numbtenant" field.
func NumbtenantLTE(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantContains applies the Contains predicate on the "numbtenant" field.
func NumbtenantContains(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantHasPrefix applies the HasPrefix predicate on the "numbtenant" field.
func NumbtenantHasPrefix(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantHasSuffix applies the HasSuffix predicate on the "numbtenant" field.
func NumbtenantHasSuffix(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantEqualFold applies the EqualFold predicate on the "numbtenant" field.
func NumbtenantEqualFold(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNumbtenant), v))
	})
}

// NumbtenantContainsFold applies the ContainsFold predicate on the "numbtenant" field.
func NumbtenantContainsFold(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNumbtenant), v))
	})
}

// PettenantEQ applies the EQ predicate on the "pettenant" field.
func PettenantEQ(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPettenant), v))
	})
}

// PettenantNEQ applies the NEQ predicate on the "pettenant" field.
func PettenantNEQ(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPettenant), v))
	})
}

// PettenantIn applies the In predicate on the "pettenant" field.
func PettenantIn(vs ...string) predicate.Lease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPettenant), v...))
	})
}

// PettenantNotIn applies the NotIn predicate on the "pettenant" field.
func PettenantNotIn(vs ...string) predicate.Lease {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Lease(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPettenant), v...))
	})
}

// PettenantGT applies the GT predicate on the "pettenant" field.
func PettenantGT(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPettenant), v))
	})
}

// PettenantGTE applies the GTE predicate on the "pettenant" field.
func PettenantGTE(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPettenant), v))
	})
}

// PettenantLT applies the LT predicate on the "pettenant" field.
func PettenantLT(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPettenant), v))
	})
}

// PettenantLTE applies the LTE predicate on the "pettenant" field.
func PettenantLTE(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPettenant), v))
	})
}

// PettenantContains applies the Contains predicate on the "pettenant" field.
func PettenantContains(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPettenant), v))
	})
}

// PettenantHasPrefix applies the HasPrefix predicate on the "pettenant" field.
func PettenantHasPrefix(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPettenant), v))
	})
}

// PettenantHasSuffix applies the HasSuffix predicate on the "pettenant" field.
func PettenantHasSuffix(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPettenant), v))
	})
}

// PettenantEqualFold applies the EqualFold predicate on the "pettenant" field.
func PettenantEqualFold(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPettenant), v))
	})
}

// PettenantContainsFold applies the ContainsFold predicate on the "pettenant" field.
func PettenantContainsFold(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPettenant), v))
	})
}

// HasWifi applies the HasEdge predicate on the "Wifi" edge.
func HasWifi() predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WifiTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WifiTable, WifiColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWifiWith applies the HasEdge predicate on the "Wifi" edge with a given conditions (other predicates).
func HasWifiWith(preds ...predicate.Wifi) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WifiInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WifiTable, WifiColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoomdetail applies the HasEdge predicate on the "Roomdetail" edge.
func HasRoomdetail() predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomdetailTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, RoomdetailTable, RoomdetailColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomdetailWith applies the HasEdge predicate on the "Roomdetail" edge with a given conditions (other predicates).
func HasRoomdetailWith(preds ...predicate.Roomdetail) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomdetailInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, RoomdetailTable, RoomdetailColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLeases applies the HasEdge predicate on the "leases" edge.
func HasLeases() predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LeasesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LeasesTable, LeasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLeasesWith applies the HasEdge predicate on the "leases" edge with a given conditions (other predicates).
func HasLeasesWith(preds ...predicate.Deposit) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LeasesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LeasesTable, LeasesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBill applies the HasEdge predicate on the "bill" edge.
func HasBill() predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BillTable, BillColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillWith applies the HasEdge predicate on the "bill" edge with a given conditions (other predicates).
func HasBillWith(preds ...predicate.Bill) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BillTable, BillColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Lease) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Lease) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
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
func Not(p predicate.Lease) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		p(s.Not())
	})
}

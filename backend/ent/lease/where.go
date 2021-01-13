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

// Lease applies equality check predicate on the "lease" field. It's identical to LeaseEQ.
func Lease(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLease), v))
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

// LeaseEQ applies the EQ predicate on the "lease" field.
func LeaseEQ(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLease), v))
	})
}

// LeaseNEQ applies the NEQ predicate on the "lease" field.
func LeaseNEQ(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLease), v))
	})
}

// LeaseIn applies the In predicate on the "lease" field.
func LeaseIn(vs ...string) predicate.Lease {
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
		s.Where(sql.In(s.C(FieldLease), v...))
	})
}

// LeaseNotIn applies the NotIn predicate on the "lease" field.
func LeaseNotIn(vs ...string) predicate.Lease {
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
		s.Where(sql.NotIn(s.C(FieldLease), v...))
	})
}

// LeaseGT applies the GT predicate on the "lease" field.
func LeaseGT(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLease), v))
	})
}

// LeaseGTE applies the GTE predicate on the "lease" field.
func LeaseGTE(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLease), v))
	})
}

// LeaseLT applies the LT predicate on the "lease" field.
func LeaseLT(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLease), v))
	})
}

// LeaseLTE applies the LTE predicate on the "lease" field.
func LeaseLTE(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLease), v))
	})
}

// LeaseContains applies the Contains predicate on the "lease" field.
func LeaseContains(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLease), v))
	})
}

// LeaseHasPrefix applies the HasPrefix predicate on the "lease" field.
func LeaseHasPrefix(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLease), v))
	})
}

// LeaseHasSuffix applies the HasSuffix predicate on the "lease" field.
func LeaseHasSuffix(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLease), v))
	})
}

// LeaseEqualFold applies the EqualFold predicate on the "lease" field.
func LeaseEqualFold(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLease), v))
	})
}

// LeaseContainsFold applies the ContainsFold predicate on the "lease" field.
func LeaseContainsFold(v string) predicate.Lease {
	return predicate.Lease(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLease), v))
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

// Code generated by entc, DO NOT EDIT.

package repairinvoice

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team15/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Bequipment applies equality check predicate on the "bequipment" field. It's identical to BequipmentEQ.
func Bequipment(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBequipment), v))
	})
}

// Emtell applies equality check predicate on the "emtell" field. It's identical to EmtellEQ.
func Emtell(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmtell), v))
	})
}

// Num applies equality check predicate on the "num" field. It's identical to NumEQ.
func Num(v int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNum), v))
	})
}

// BequipmentEQ applies the EQ predicate on the "bequipment" field.
func BequipmentEQ(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBequipment), v))
	})
}

// BequipmentNEQ applies the NEQ predicate on the "bequipment" field.
func BequipmentNEQ(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBequipment), v))
	})
}

// BequipmentIn applies the In predicate on the "bequipment" field.
func BequipmentIn(vs ...string) predicate.Repairinvoice {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repairinvoice(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBequipment), v...))
	})
}

// BequipmentNotIn applies the NotIn predicate on the "bequipment" field.
func BequipmentNotIn(vs ...string) predicate.Repairinvoice {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repairinvoice(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBequipment), v...))
	})
}

// BequipmentGT applies the GT predicate on the "bequipment" field.
func BequipmentGT(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBequipment), v))
	})
}

// BequipmentGTE applies the GTE predicate on the "bequipment" field.
func BequipmentGTE(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBequipment), v))
	})
}

// BequipmentLT applies the LT predicate on the "bequipment" field.
func BequipmentLT(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBequipment), v))
	})
}

// BequipmentLTE applies the LTE predicate on the "bequipment" field.
func BequipmentLTE(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBequipment), v))
	})
}

// BequipmentContains applies the Contains predicate on the "bequipment" field.
func BequipmentContains(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBequipment), v))
	})
}

// BequipmentHasPrefix applies the HasPrefix predicate on the "bequipment" field.
func BequipmentHasPrefix(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBequipment), v))
	})
}

// BequipmentHasSuffix applies the HasSuffix predicate on the "bequipment" field.
func BequipmentHasSuffix(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBequipment), v))
	})
}

// BequipmentEqualFold applies the EqualFold predicate on the "bequipment" field.
func BequipmentEqualFold(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBequipment), v))
	})
}

// BequipmentContainsFold applies the ContainsFold predicate on the "bequipment" field.
func BequipmentContainsFold(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBequipment), v))
	})
}

// EmtellEQ applies the EQ predicate on the "emtell" field.
func EmtellEQ(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmtell), v))
	})
}

// EmtellNEQ applies the NEQ predicate on the "emtell" field.
func EmtellNEQ(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmtell), v))
	})
}

// EmtellIn applies the In predicate on the "emtell" field.
func EmtellIn(vs ...string) predicate.Repairinvoice {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repairinvoice(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmtell), v...))
	})
}

// EmtellNotIn applies the NotIn predicate on the "emtell" field.
func EmtellNotIn(vs ...string) predicate.Repairinvoice {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repairinvoice(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmtell), v...))
	})
}

// EmtellGT applies the GT predicate on the "emtell" field.
func EmtellGT(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmtell), v))
	})
}

// EmtellGTE applies the GTE predicate on the "emtell" field.
func EmtellGTE(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmtell), v))
	})
}

// EmtellLT applies the LT predicate on the "emtell" field.
func EmtellLT(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmtell), v))
	})
}

// EmtellLTE applies the LTE predicate on the "emtell" field.
func EmtellLTE(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmtell), v))
	})
}

// EmtellContains applies the Contains predicate on the "emtell" field.
func EmtellContains(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmtell), v))
	})
}

// EmtellHasPrefix applies the HasPrefix predicate on the "emtell" field.
func EmtellHasPrefix(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmtell), v))
	})
}

// EmtellHasSuffix applies the HasSuffix predicate on the "emtell" field.
func EmtellHasSuffix(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmtell), v))
	})
}

// EmtellEqualFold applies the EqualFold predicate on the "emtell" field.
func EmtellEqualFold(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmtell), v))
	})
}

// EmtellContainsFold applies the ContainsFold predicate on the "emtell" field.
func EmtellContainsFold(v string) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmtell), v))
	})
}

// NumEQ applies the EQ predicate on the "num" field.
func NumEQ(v int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNum), v))
	})
}

// NumNEQ applies the NEQ predicate on the "num" field.
func NumNEQ(v int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNum), v))
	})
}

// NumIn applies the In predicate on the "num" field.
func NumIn(vs ...int) predicate.Repairinvoice {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repairinvoice(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNum), v...))
	})
}

// NumNotIn applies the NotIn predicate on the "num" field.
func NumNotIn(vs ...int) predicate.Repairinvoice {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repairinvoice(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNum), v...))
	})
}

// NumGT applies the GT predicate on the "num" field.
func NumGT(v int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNum), v))
	})
}

// NumGTE applies the GTE predicate on the "num" field.
func NumGTE(v int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNum), v))
	})
}

// NumLT applies the LT predicate on the "num" field.
func NumLT(v int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNum), v))
	})
}

// NumLTE applies the LTE predicate on the "num" field.
func NumLTE(v int) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNum), v))
	})
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
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

// HasRentalstatus applies the HasEdge predicate on the "Rentalstatus" edge.
func HasRentalstatus() predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RentalstatusTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RentalstatusTable, RentalstatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRentalstatusWith applies the HasEdge predicate on the "Rentalstatus" edge with a given conditions (other predicates).
func HasRentalstatusWith(preds ...predicate.Rentalstatus) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RentalstatusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RentalstatusTable, RentalstatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLease applies the HasEdge predicate on the "Lease" edge.
func HasLease() predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LeaseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LeaseTable, LeaseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLeaseWith applies the HasEdge predicate on the "Lease" edge with a given conditions (other predicates).
func HasLeaseWith(preds ...predicate.Lease) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LeaseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LeaseTable, LeaseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Repairinvoice) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Repairinvoice) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
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
func Not(p predicate.Repairinvoice) predicate.Repairinvoice {
	return predicate.Repairinvoice(func(s *sql.Selector) {
		p(s.Not())
	})
}

// Code generated by entc, DO NOT EDIT.

package equipment

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team15/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Equipment applies equality check predicate on the "equipment" field. It's identical to EquipmentEQ.
func Equipment(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEquipment), v))
	})
}

// EquipmentEQ applies the EQ predicate on the "equipment" field.
func EquipmentEQ(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEquipment), v))
	})
}

// EquipmentNEQ applies the NEQ predicate on the "equipment" field.
func EquipmentNEQ(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEquipment), v))
	})
}

// EquipmentIn applies the In predicate on the "equipment" field.
func EquipmentIn(vs ...string) predicate.Equipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Equipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEquipment), v...))
	})
}

// EquipmentNotIn applies the NotIn predicate on the "equipment" field.
func EquipmentNotIn(vs ...string) predicate.Equipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Equipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEquipment), v...))
	})
}

// EquipmentGT applies the GT predicate on the "equipment" field.
func EquipmentGT(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEquipment), v))
	})
}

// EquipmentGTE applies the GTE predicate on the "equipment" field.
func EquipmentGTE(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEquipment), v))
	})
}

// EquipmentLT applies the LT predicate on the "equipment" field.
func EquipmentLT(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEquipment), v))
	})
}

// EquipmentLTE applies the LTE predicate on the "equipment" field.
func EquipmentLTE(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEquipment), v))
	})
}

// EquipmentContains applies the Contains predicate on the "equipment" field.
func EquipmentContains(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEquipment), v))
	})
}

// EquipmentHasPrefix applies the HasPrefix predicate on the "equipment" field.
func EquipmentHasPrefix(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEquipment), v))
	})
}

// EquipmentHasSuffix applies the HasSuffix predicate on the "equipment" field.
func EquipmentHasSuffix(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEquipment), v))
	})
}

// EquipmentEqualFold applies the EqualFold predicate on the "equipment" field.
func EquipmentEqualFold(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEquipment), v))
	})
}

// EquipmentContainsFold applies the ContainsFold predicate on the "equipment" field.
func EquipmentContainsFold(v string) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEquipment), v))
	})
}

// HasRoomdetail applies the HasEdge predicate on the "roomdetail" edge.
func HasRoomdetail() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomdetailTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomdetailTable, RoomdetailColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomdetailWith applies the HasEdge predicate on the "roomdetail" edge with a given conditions (other predicates).
func HasRoomdetailWith(preds ...predicate.Roomdetail) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomdetailInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomdetailTable, RoomdetailColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
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
func Not(p predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		p(s.Not())
	})
}

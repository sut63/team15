// Code generated by entc, DO NOT EDIT.

package roomdetail

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team15/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Roomnumber applies equality check predicate on the "roomnumber" field. It's identical to RoomnumberEQ.
func Roomnumber(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomnumber), v))
	})
}

// Roomtypename applies equality check predicate on the "roomtypename" field. It's identical to RoomtypenameEQ.
func Roomtypename(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomtypename), v))
	})
}

// Roomprice applies equality check predicate on the "roomprice" field. It's identical to RoompriceEQ.
func Roomprice(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomprice), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// Sleep applies equality check predicate on the "sleep" field. It's identical to SleepEQ.
func Sleep(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSleep), v))
	})
}

// Bed applies equality check predicate on the "bed" field. It's identical to BedEQ.
func Bed(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBed), v))
	})
}

// RoomnumberEQ applies the EQ predicate on the "roomnumber" field.
func RoomnumberEQ(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberNEQ applies the NEQ predicate on the "roomnumber" field.
func RoomnumberNEQ(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberIn applies the In predicate on the "roomnumber" field.
func RoomnumberIn(vs ...string) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoomnumber), v...))
	})
}

// RoomnumberNotIn applies the NotIn predicate on the "roomnumber" field.
func RoomnumberNotIn(vs ...string) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoomnumber), v...))
	})
}

// RoomnumberGT applies the GT predicate on the "roomnumber" field.
func RoomnumberGT(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberGTE applies the GTE predicate on the "roomnumber" field.
func RoomnumberGTE(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberLT applies the LT predicate on the "roomnumber" field.
func RoomnumberLT(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberLTE applies the LTE predicate on the "roomnumber" field.
func RoomnumberLTE(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberContains applies the Contains predicate on the "roomnumber" field.
func RoomnumberContains(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberHasPrefix applies the HasPrefix predicate on the "roomnumber" field.
func RoomnumberHasPrefix(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberHasSuffix applies the HasSuffix predicate on the "roomnumber" field.
func RoomnumberHasSuffix(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberEqualFold applies the EqualFold predicate on the "roomnumber" field.
func RoomnumberEqualFold(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoomnumber), v))
	})
}

// RoomnumberContainsFold applies the ContainsFold predicate on the "roomnumber" field.
func RoomnumberContainsFold(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoomnumber), v))
	})
}

// RoomtypenameEQ applies the EQ predicate on the "roomtypename" field.
func RoomtypenameEQ(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameNEQ applies the NEQ predicate on the "roomtypename" field.
func RoomtypenameNEQ(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameIn applies the In predicate on the "roomtypename" field.
func RoomtypenameIn(vs ...string) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoomtypename), v...))
	})
}

// RoomtypenameNotIn applies the NotIn predicate on the "roomtypename" field.
func RoomtypenameNotIn(vs ...string) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoomtypename), v...))
	})
}

// RoomtypenameGT applies the GT predicate on the "roomtypename" field.
func RoomtypenameGT(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameGTE applies the GTE predicate on the "roomtypename" field.
func RoomtypenameGTE(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameLT applies the LT predicate on the "roomtypename" field.
func RoomtypenameLT(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameLTE applies the LTE predicate on the "roomtypename" field.
func RoomtypenameLTE(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameContains applies the Contains predicate on the "roomtypename" field.
func RoomtypenameContains(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameHasPrefix applies the HasPrefix predicate on the "roomtypename" field.
func RoomtypenameHasPrefix(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameHasSuffix applies the HasSuffix predicate on the "roomtypename" field.
func RoomtypenameHasSuffix(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameEqualFold applies the EqualFold predicate on the "roomtypename" field.
func RoomtypenameEqualFold(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoomtypename), v))
	})
}

// RoomtypenameContainsFold applies the ContainsFold predicate on the "roomtypename" field.
func RoomtypenameContainsFold(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoomtypename), v))
	})
}

// RoompriceEQ applies the EQ predicate on the "roomprice" field.
func RoompriceEQ(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomprice), v))
	})
}

// RoompriceNEQ applies the NEQ predicate on the "roomprice" field.
func RoompriceNEQ(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoomprice), v))
	})
}

// RoompriceIn applies the In predicate on the "roomprice" field.
func RoompriceIn(vs ...int) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoomprice), v...))
	})
}

// RoompriceNotIn applies the NotIn predicate on the "roomprice" field.
func RoompriceNotIn(vs ...int) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoomprice), v...))
	})
}

// RoompriceGT applies the GT predicate on the "roomprice" field.
func RoompriceGT(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoomprice), v))
	})
}

// RoompriceGTE applies the GTE predicate on the "roomprice" field.
func RoompriceGTE(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoomprice), v))
	})
}

// RoompriceLT applies the LT predicate on the "roomprice" field.
func RoompriceLT(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoomprice), v))
	})
}

// RoompriceLTE applies the LTE predicate on the "roomprice" field.
func RoompriceLTE(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoomprice), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// SleepEQ applies the EQ predicate on the "sleep" field.
func SleepEQ(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSleep), v))
	})
}

// SleepNEQ applies the NEQ predicate on the "sleep" field.
func SleepNEQ(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSleep), v))
	})
}

// SleepIn applies the In predicate on the "sleep" field.
func SleepIn(vs ...int) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSleep), v...))
	})
}

// SleepNotIn applies the NotIn predicate on the "sleep" field.
func SleepNotIn(vs ...int) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSleep), v...))
	})
}

// SleepGT applies the GT predicate on the "sleep" field.
func SleepGT(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSleep), v))
	})
}

// SleepGTE applies the GTE predicate on the "sleep" field.
func SleepGTE(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSleep), v))
	})
}

// SleepLT applies the LT predicate on the "sleep" field.
func SleepLT(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSleep), v))
	})
}

// SleepLTE applies the LTE predicate on the "sleep" field.
func SleepLTE(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSleep), v))
	})
}

// BedEQ applies the EQ predicate on the "bed" field.
func BedEQ(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBed), v))
	})
}

// BedNEQ applies the NEQ predicate on the "bed" field.
func BedNEQ(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBed), v))
	})
}

// BedIn applies the In predicate on the "bed" field.
func BedIn(vs ...int) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBed), v...))
	})
}

// BedNotIn applies the NotIn predicate on the "bed" field.
func BedNotIn(vs ...int) predicate.Roomdetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomdetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBed), v...))
	})
}

// BedGT applies the GT predicate on the "bed" field.
func BedGT(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBed), v))
	})
}

// BedGTE applies the GTE predicate on the "bed" field.
func BedGTE(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBed), v))
	})
}

// BedLT applies the LT predicate on the "bed" field.
func BedLT(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBed), v))
	})
}

// BedLTE applies the LTE predicate on the "bed" field.
func BedLTE(v int) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBed), v))
	})
}

// HasPledge applies the HasEdge predicate on the "pledge" edge.
func HasPledge() predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PledgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PledgeTable, PledgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPledgeWith applies the HasEdge predicate on the "pledge" edge with a given conditions (other predicates).
func HasPledgeWith(preds ...predicate.Pledge) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PledgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PledgeTable, PledgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPetrule applies the HasEdge predicate on the "petrule" edge.
func HasPetrule() predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PetruleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PetruleTable, PetruleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPetruleWith applies the HasEdge predicate on the "petrule" edge with a given conditions (other predicates).
func HasPetruleWith(preds ...predicate.Petrule) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PetruleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PetruleTable, PetruleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBedtype applies the HasEdge predicate on the "bedtype" edge.
func HasBedtype() predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BedtypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BedtypeTable, BedtypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBedtypeWith applies the HasEdge predicate on the "bedtype" edge with a given conditions (other predicates).
func HasBedtypeWith(preds ...predicate.Bedtype) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BedtypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BedtypeTable, BedtypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
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

// HasJobposition applies the HasEdge predicate on the "jobposition" edge.
func HasJobposition() predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(JobpositionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobpositionTable, JobpositionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJobpositionWith applies the HasEdge predicate on the "jobposition" edge with a given conditions (other predicates).
func HasJobpositionWith(preds ...predicate.Jobposition) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(JobpositionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobpositionTable, JobpositionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStaytype applies the HasEdge predicate on the "staytype" edge.
func HasStaytype() predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StaytypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StaytypeTable, StaytypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStaytypeWith applies the HasEdge predicate on the "staytype" edge with a given conditions (other predicates).
func HasStaytypeWith(preds ...predicate.Staytype) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StaytypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StaytypeTable, StaytypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoomdetails applies the HasEdge predicate on the "roomdetails" edge.
func HasRoomdetails() predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomdetailsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, RoomdetailsTable, RoomdetailsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomdetailsWith applies the HasEdge predicate on the "roomdetails" edge with a given conditions (other predicates).
func HasRoomdetailsWith(preds ...predicate.Lease) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomdetailsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, RoomdetailsTable, RoomdetailsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCleaningrooms applies the HasEdge predicate on the "cleaningrooms" edge.
func HasCleaningrooms() predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CleaningroomsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CleaningroomsTable, CleaningroomsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCleaningroomsWith applies the HasEdge predicate on the "cleaningrooms" edge with a given conditions (other predicates).
func HasCleaningroomsWith(preds ...predicate.Cleaningroom) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CleaningroomsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CleaningroomsTable, CleaningroomsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Roomdetail) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Roomdetail) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
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
func Not(p predicate.Roomdetail) predicate.Roomdetail {
	return predicate.Roomdetail(func(s *sql.Selector) {
		p(s.Not())
	})
}

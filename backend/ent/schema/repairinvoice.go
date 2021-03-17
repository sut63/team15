package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Repairinvoice holds the schema definition for the Repairinvoice entity.
type Repairinvoice struct {
	ent.Schema
}

// Fields of the Repairinvoice.
func (Repairinvoice) Fields() []ent.Field {
	return []ent.Field{
		field.String("bequipment").Validate(func(s string) error {
			match, _ := regexp.MatchString("^([A-Z]{3})-([0-9]{4})$", s)
			if !match {
				return errors.New("The Equipment format is not valid.")
			}
			return nil
		}),
		field.String("emtell").MaxLen(12).MinLen(12),
		field.Int("num").Min(0).Max(10),
	}
}

// Edges of the Repairinvoice.
func (Repairinvoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("employee", Employee.Type).
			Ref("repairinvoices").
			Unique(),

		edge.From("Rentalstatus", Rentalstatus.Type).
			Ref("repairinvoices").
			Unique(),

			edge.From("Lease", Lease.Type).
			Ref("repairinvoices").
			Unique(),
	}
}

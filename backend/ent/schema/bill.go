package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedtime"),
		field.String("tell").MaxLen(12).MinLen(12),
		field.String("taxpayer").MaxLen(17).MinLen(17),
		field.String("total").Validate(func(s string) error {
			match, _ := regexp.MatchString("^(([0-9]{1}).([0-9]{2}))$|^(([0-9]{2}).([0-9]{2}))$|^(([0-9]{3}).([0-9]{2}))$|^(([0-9]{4}).([0-9]{2}))$", s)
			if !match {
				return errors.New("The price pattern does not match.")
			}
			return nil
		}),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Situation", Situation.Type).
			Ref("situations").
			Unique(),

		edge.From("Payment", Payment.Type).
			Ref("payments").
			Unique(),

		edge.From("Lease", Lease.Type).
			Ref("bill").
			Unique(),
	}
}

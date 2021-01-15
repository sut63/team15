package schema

import (
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
		field.Int("total"),
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


	}
}

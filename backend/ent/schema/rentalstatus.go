package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Rentalstatus holds the schema definition for the Rentalstatus entity.
type Rentalstatus struct {
	ent.Schema
}

// Fields of the Rentalstatus.
func (Rentalstatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("rentalstatus").Unique(),
	}
}

// Edges of the Rentalstatus.
func (Rentalstatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("repairinvoices", Repairinvoice.Type),
	}
}

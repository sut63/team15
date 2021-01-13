package schema

import (
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
		field.String("bequipment"),
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
	}
}

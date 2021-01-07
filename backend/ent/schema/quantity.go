package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Quantity holds the schema definition for the Quantity entity.
type Quantity struct {
	ent.Schema
}

// Fields of the Quantity.
func (Quantity) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity").Positive(),
	}
}

// Edges of the Quantity.
func (Quantity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("quantity"),
	}
}

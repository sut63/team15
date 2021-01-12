package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Nearbyplace holds the schema definition for the Nearbyplace entity.
type Nearbyplace struct {
	ent.Schema
}

// Fields of the Nearbyplace.
func (Nearbyplace) Fields() []ent.Field {
	return []ent.Field{
		field.String("nearbyplace").Unique(),
	}
}

// Edges of the Nearbyplace.
func (Nearbyplace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roomdetail", Roomdetail.Type),
	}
}

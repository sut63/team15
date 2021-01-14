package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Pledge holds the schema definition for the Pledge entity.
type Pledge struct {
	ent.Schema
}

// Fields of the Pledge.
func (Pledge) Fields() []ent.Field {
	return []ent.Field{
		field.String("provision").Unique(),
	}
}

// Edges of the Pledge.
func (Pledge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roomdetails", Roomdetail.Type),
	}
}

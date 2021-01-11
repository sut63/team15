package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Staytype holds the schema definition for the Staytype entity.
type Staytype struct {
	ent.Schema
}

// Fields of the Staytype.
func (Staytype) Fields() []ent.Field {
	return []ent.Field{
		field.String("staytype"),
	}
}

// Edges of the Staytype.
func (Staytype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roomdetails", Roomdetail.Type),
	}
}

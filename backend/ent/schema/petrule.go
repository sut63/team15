package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Petrule holds the schema definition for the Petrule entity.
type Petrule struct {
	ent.Schema
}

// Fields of the Petrule.
func (Petrule) Fields() []ent.Field {
	return []ent.Field{
		field.String("petrule").Unique(),
	}
}

// Edges of the Petrule.
func (Petrule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roomdetails", Roomdetail.Type),
	}
}

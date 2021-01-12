package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Facilitie holds the schema definition for the Facilitie entity.
type Facilitie struct {
	ent.Schema
}

// Fields of the Facilitie.
func (Facilitie) Fields() []ent.Field {
	return []ent.Field{
		field.String("facilitie").Unique(),
	}
}

// Edges of the Facilitie.
func (Facilitie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roomdetail", Roomdetail.Type),
	}
}

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// StayType holds the schema definition for the StayType entity.
type StayType struct {
	ent.Schema
}

// Fields of the StayType.
func (StayType) Fields() []ent.Field {
	return []ent.Field{
		field.String("staytype").Unique(),
	}
}

// Edges of the StayType.
func (StayType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("staytype"),
	}
}

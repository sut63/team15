package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// NearbyPlace holds the schema definition for the NearbyPlace entity.
type NearbyPlace struct {
	ent.Schema
}

// Fields of the NearbyPlace.
func (NearbyPlace) Fields() []ent.Field {
	return []ent.Field{
		field.String("placename").Unique(),
	}
}

// Edges of the NearbyPlace.
func (NearbyPlace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("nearbyplace"),
	}
}

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.Int("roomprice").Positive(),
		field.String("roomtypename").Unique(),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("quantity", Quantity.Type).Unique(),
		edge.To("staytype", StayType.Type).Unique(),
		edge.To("facilities", Facility.Type),
		edge.To("equipments", Equipment.Type),
		edge.To("nearbyplace", NearbyPlace.Type),
		edge.To("cleaningrooms", CleaningRoom.Type),
	}
}

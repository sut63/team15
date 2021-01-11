package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// CleaningRoom holds the schema definition for the CleaningRoom entity.
type CleaningRoom struct {
	ent.Schema
}

// Fields of the CleaningRoom.
func (CleaningRoom) Fields() []ent.Field {
	return []ent.Field{
		field.Time("dateandstarttime"),
		field.String("note"),
	}
}

// Edges of the CleaningRoom.
func (CleaningRoom) Edges() []ent.Edge {
	return []ent.Edge{
		/*edge.From("Room", Room.Type).
		  Ref("cleaningrooms").
		  Unique(),*/

		edge.From("CleanerName", CleanerName.Type).
			Ref("cleaningrooms").
			Unique(),

		edge.From("LengthTime", LengthTime.Type).
			Ref("cleaningrooms").
			Unique(),
	}
}

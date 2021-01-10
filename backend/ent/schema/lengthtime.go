package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// LengthTime holds the schema definition for the LengthTime entity.
type LengthTime struct {
	ent.Schema
}

// Fields of the LengthTime.
func (LengthTime) Fields() []ent.Field {
	return []ent.Field{
		field.String("lengthtime").Unique().NotEmpty(),
	}
}

// Edges of the LengthTime.
func (LengthTime) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cleaningrooms", CleaningRoom.Type). StorageKey(edge.Column("lengthtime_id")),
	}
}

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// CleanerName holds the schema definition for the CleanerName entity.
type CleanerName struct {
	ent.Schema
}

// Fields of the CleanerName.
func (CleanerName) Fields() []ent.Field {
	return []ent.Field{
		field.String("cleanername").Unique().NotEmpty(),
	}
}

// Edges of the CleanerName.
func (CleanerName) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cleaningrooms", CleaningRoom.Type). StorageKey(edge.Column("cleanerroom_id")),
	}
}

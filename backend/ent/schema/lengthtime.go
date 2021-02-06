package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// Lengthtime holds the schema definition for the Lengthtime entity.
type Lengthtime struct {
	ent.Schema
}

// Fields of the Lengthtime.
func (Lengthtime) Fields() []ent.Field {
	return []ent.Field{
		field.String("lengthtime").Unique().NotEmpty(),
	}
}

// Edges of the Lengthtime.
func (Lengthtime) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cleaningrooms", Cleaningroom.Type). StorageKey(edge.Column("lengthtime_id")),
	}
}

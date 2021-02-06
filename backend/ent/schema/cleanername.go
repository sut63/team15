package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// Cleanername holds the schema definition for the Cleanername entity.
type Cleanername struct {
	ent.Schema
}

// Fields of the Cleanername.
func (Cleanername) Fields() []ent.Field {
	return []ent.Field{
		field.String("cleanername").Unique().NotEmpty(),
	}
}

// Edges of the Cleanername.
func (Cleanername) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cleaningrooms", Cleaningroom.Type). StorageKey(edge.Column("cleanerroom_id")),
	}
}

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Jobposition holds the schema definition for the Jobposition entity.
type Jobposition struct {
	ent.Schema
}

// Fields of the Jobposition.
func (Jobposition) Fields() []ent.Field {
	return []ent.Field{
		field.String("position_name").NotEmpty().Unique(),
	}
}

// Edges of the Jobposition.
func (Jobposition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employees", Employee.Type).StorageKey(edge.Column("jobposition_id")),
	}
}

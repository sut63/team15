package schema

import(
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)
// Situation holds the schema definition for the Situation entity.
type Situation struct {
	ent.Schema
}

// Fields of the Situation.
func (Situation) Fields() []ent.Field {
	return []ent.Field{
		field.String("situationname").Unique().NotEmpty(),
		
	}
}

// Edges of the Situation.
func (Situation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("situations", Bill.Type).StorageKey(edge.Column("situation_id")),
		
	}
}

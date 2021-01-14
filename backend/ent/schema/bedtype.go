package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Bedtype holds the schema definition for the Bedtype entity.
type Bedtype struct {
	ent.Schema
}

// Fields of the Bedtype.
func (Bedtype) Fields() []ent.Field {
	return []ent.Field{
		field.String("bedtypename").Unique(),
	}
}

// Edges of the Bedtype.
func (Bedtype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roomdetails", Roomdetail.Type),
	}
}

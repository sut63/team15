package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Roomdetail holds the schema definition for the Roomdetail entity.
type Roomdetail struct {
	ent.Schema
}

// Fields of the Roomdetail.
func (Roomdetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("roomtypename").Unique(),
		field.String("roomprice").Unique(),
	}
}

// Edges of the Roomdetail.
func (Roomdetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipments", Equipment.Type),
		edge.To("facilities", Facilitie.Type),
		edge.To("nearbyplaces", Nearbyplace.Type),
		edge.From("quantity", Quantity.Type).
			Ref("roomdetails").
			Unique(),
		edge.From("staytype", Staytype.Type).
			Ref("roomdetails").
			Unique(),
	}
}

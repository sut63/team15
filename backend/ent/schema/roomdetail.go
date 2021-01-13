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
		field.String("roomnumber").NotEmpty().Unique(),
		field.String("roomtypename").NotEmpty(),
		field.String("roomprice").NotEmpty(),
	}
}

// Edges of the Roomdetail.
func (Roomdetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipments", Equipment.Type).Ref("roomdetail").Unique(),
		edge.From("facilities", Facilitie.Type).Ref("roomdetail").Unique(),
		edge.From("nearbyplaces", Nearbyplace.Type).Ref("roomdetail").Unique(),
		edge.From("employee", Employee.Type).
			Ref("roomdetails").
			Unique(),
		edge.From("quantity", Quantity.Type).
			Ref("roomdetails").
			Unique(),
		edge.From("staytype", Staytype.Type).
			Ref("roomdetails").
			Unique(),
		edge.To("roomdetails", Lease.Type).Unique().StorageKey(edge.Column("room_num")),
	}
}

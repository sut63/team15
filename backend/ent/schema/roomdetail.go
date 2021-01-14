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
		field.String("roomnumber").Unique(),
		field.String("roomtypename"),
		field.String("roomprice"),
		field.String("sleep"),
		field.String("bed"),
	}
}

// Edges of the Roomdetail.
func (Roomdetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pledge", Pledge.Type).Ref("roomdetails").Unique(),
		edge.From("petrule", Petrule.Type).Ref("roomdetails").Unique(),
		edge.From("bedtype", Bedtype.Type).Ref("roomdetails").Unique(),
		edge.From("employee", Employee.Type).Ref("roomdetails").Unique(),
		edge.From("staytype", Staytype.Type).Ref("roomdetails").Unique(),
		edge.To("roomdetails", Lease.Type).Unique().StorageKey(edge.Column("room_num")),
	}
}

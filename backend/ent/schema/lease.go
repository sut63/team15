package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Lease holds the schema definition for the Lease entity.
type Lease struct {
	ent.Schema
}

// Fields of the Lease.
func (Lease) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedtime"),
		field.String("tenant"),
	}
}

// Edges of the Lease.
func (Lease) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Wifi", Wifi.Type).
			Ref("wifis").
			Unique(),
		edge.From("Roomdetail", Roomdetail.Type).
			Ref("roomdetails").
			Unique().
			Required(),
		edge.From("employee", Employee.Type).Ref("leases").Unique(),
		edge.To("leases", Deposit.Type).StorageKey(edge.Column("lease_id")),
	}
}

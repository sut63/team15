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
		field.String("tenant").NotEmpty(),
		field.String("numbtenant").MaxLen(12).MinLen(12),
		field.String("idtenant").MaxLen(17).MinLen(17),
		field.Int("agetenant").Min(1),
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
		edge.From("employee", Employee.Type).Ref("leasess").Unique(),
		edge.To("leases", Deposit.Type).StorageKey(edge.Column("lease_id")),
		edge.To("bill", Bill.Type).StorageKey(edge.Column("lease_id")),
		edge.To("repairinvoices", Repairinvoice.Type).StorageKey(edge.Column("lease_id")),
	}
}

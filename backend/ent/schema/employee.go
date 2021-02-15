package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("password").NotEmpty(),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employees", Deposit.Type).StorageKey(edge.Column("employee_id")),
		edge.To("leasess", Lease.Type).StorageKey(edge.Column("employee_id")),
		edge.To("roomdetails", Roomdetail.Type).StorageKey(edge.Column("employee_id")),
		edge.From("jobposition", Jobposition.Type).
			Ref("employees").
			Unique(),
		edge.To("repairinvoices", Repairinvoice.Type).StorageKey(edge.Column("employee_id")),
		edge.To("cleaningrooms", Cleaningroom.Type).StorageKey(edge.Column("employee_id")),
	}
}

package schema

import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Deposit holds the schema definition for the Deposit entity.
type Deposit struct {
	ent.Schema
}

// Fields of the Deposit.
func (Deposit) Fields() []ent.Field {
	return []ent.Field{
	   field.Time("addedtime"),
	   field.String("info"),
   }
}

// Edges of the Deposit.
func (Deposit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Employee", Employee.Type).
            Ref("employees").
            Unique(),

		edge.From("Statusd", Statusd.Type).
            Ref("statusds").
			Unique(),

	    edge.From("Lease", Lease.Type).
            Ref("leases").
            Unique(),
	}
}

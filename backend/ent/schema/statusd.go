package schema

import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Statusd holds the schema definition for the Statusd entity.
type Statusd struct {
	ent.Schema
}

// Fields of the Statusd.
func (Statusd) Fields() []ent.Field {
	return []ent.Field{
       field.String("statusdname").Unique().NotEmpty(),
   }
}

// Edges of the Statusd.
func (Statusd) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("statusds", Deposit.Type).StorageKey(edge.Column("statusd_id")),
    }
}

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )
// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.String("paymentname").Unique().NotEmpty(),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("payments", Bill.Type).StorageKey(edge.Column("payment_id")),
    }
}

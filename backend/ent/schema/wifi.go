package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Wifi holds the schema definition for the Wifi entity.
type Wifi struct {
	ent.Schema
}

// Fields of the Wifi.
func (Wifi) Fields() []ent.Field {
	return []ent.Field{
		field.String("wifiname").Unique().NotEmpty(),
	}
}

// Edges of the Wifi.
func (Wifi) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("wifis", Lease.Type).StorageKey(edge.Column("wifi_id")),
	}
}

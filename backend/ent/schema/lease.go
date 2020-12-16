package schema

import "github.com/facebookincubator/ent"

// Lease holds the schema definition for the Lease entity.
type Lease struct {
	ent.Schema
}

// Fields of the Lease.
func (Lease) Fields() []ent.Field {
	return nil
}

// Edges of the Lease.
func (Lease) Edges() []ent.Edge {
	return nil
}

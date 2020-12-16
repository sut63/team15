package schema

import "github.com/facebookincubator/ent"

// Cleaning holds the schema definition for the Cleaning entity.
type Cleaning struct {
	ent.Schema
}

// Fields of the Cleaning.
func (Cleaning) Fields() []ent.Field {
	return nil
}

// Edges of the Cleaning.
func (Cleaning) Edges() []ent.Edge {
	return nil
}

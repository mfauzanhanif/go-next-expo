package schema

import "entgo.io/ent"

// Tenants holds the schema definition for the Tenants entity.
type Tenants struct {
	ent.Schema
}

// Fields of the Tenants.
func (Tenants) Fields() []ent.Field {
	return nil
}

// Edges of the Tenants.
func (Tenants) Edges() []ent.Edge {
	return nil
}

package schema

import "github.com/facebookincubator/ent"

// WasherEntity holds the schema definition for the WasherEntity entity.
type WasherEntity struct {
	ent.Schema
}

// Fields of the WasherEntity.
func (WasherEntity) Fields() []ent.Field {
	return nil
}

// Edges of the WasherEntity.
func (WasherEntity) Edges() []ent.Edge {
	return nil
}

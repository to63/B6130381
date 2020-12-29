package schema

import "github.com/facebookincubator/ent"

// Refund holds the schema definition for the Refund entity.
type Refund struct {
	ent.Schema
}

// Fields of the Refund.
func (Refund) Fields() []ent.Field {
	return nil
}

// Edges of the Refund.
func (Refund) Edges() []ent.Edge {
	return nil
}

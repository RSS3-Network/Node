package schema

import "entgo.io/ent"

// Feed holds the schema definition for the Feed entity.
type Feed struct {
	ent.Schema
}

// Fields of the Feed.
func (Feed) Fields() []ent.Field {
	return nil
}

// Edges of the Feed.
func (Feed) Edges() []ent.Edge {
	return nil
}

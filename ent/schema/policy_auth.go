package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
)

// PolicyAuth holds the schema definition for the PolicyAuth entity.
type PolicyAuth struct {
	ent.Schema
}

// Fields of the PolicyAuth.
func (PolicyAuth) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the PolicyAuth.
func (PolicyAuth) Edges() []ent.Edge {
	return nil
}

// Mixin of the PolicyAuth.
func (PolicyAuth) Mixin() []ent.Mixin {
    return nil
}

// Indexes of the PolicyAuth.
func (PolicyAuth) Indexes() []ent.Index {
    return nil
}

// Annotations of the PolicyAuth
func (PolicyAuth) Annotations() []schema.Annotation {
	return nil
}
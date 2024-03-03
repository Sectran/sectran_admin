package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
)

// UserGroup holds the schema definition for the UserGroup entity.
type UserGroup struct {
	ent.Schema
}

// Fields of the UserGroup.
func (UserGroup) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the UserGroup.
func (UserGroup) Edges() []ent.Edge {
	return nil
}

// Mixin of the UserGroup.
func (UserGroup) Mixin() []ent.Mixin {
	return nil
}

// Indexes of the UserGroup.
func (UserGroup) Indexes() []ent.Index {
	return nil
}

// Annotations of the UserGroup
func (UserGroup) Annotations() []schema.Annotation {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// login_logs holds the schema definition for the login_logs entity.
type login_logs struct {
	ent.Schema
}

// Fields of the login_logs.
func (login_logs) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the login_logs.
func (login_logs) Edges() []ent.Edge {
	return nil
}

// Mixin of the login_logs.
func (login_logs) Mixin() []ent.Mixin {
    return []ent.Mixin{}
}

// Indexes of the login_logs.
func (login_logs) Indexes() []ent.Index {
    return nil
}

// Annotations of the login_logs
func (login_logs) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "login_logs"},
	}
}
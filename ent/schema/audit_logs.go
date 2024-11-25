package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// audit_logs holds the schema definition for the audit_logs entity.
type audit_logs struct {
	ent.Schema
}

// Fields of the audit_logs.
func (audit_logs) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the audit_logs.
func (audit_logs) Edges() []ent.Edge {
	return nil
}

// Mixin of the audit_logs.
func (audit_logs) Mixin() []ent.Mixin {
    return []ent.Mixin{}
}

// Indexes of the audit_logs.
func (audit_logs) Indexes() []ent.Index {
    return nil
}

// Annotations of the audit_logs
func (audit_logs) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "audit_logs"},
	}
}
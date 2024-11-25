package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// acc_logs holds the schema definition for the acc_logs entity.
type acc_logs struct {
	ent.Schema
}

// Fields of the acc_logs.
func (acc_logs) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the acc_logs.
func (acc_logs) Edges() []ent.Edge {
	return nil
}

// Mixin of the acc_logs.
func (acc_logs) Mixin() []ent.Mixin {
    return []ent.Mixin{}
}

// Indexes of the acc_logs.
func (acc_logs) Indexes() []ent.Index {
    return nil
}

// Annotations of the acc_logs
func (acc_logs) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "acc_logs"},
	}
}
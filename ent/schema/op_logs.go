package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// op_logs holds the schema definition for the op_logs entity.
type op_logs struct {
	ent.Schema
}

// Fields of the op_logs.
func (op_logs) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the op_logs.
func (op_logs) Edges() []ent.Edge {
	return nil
}

// Mixin of the op_logs.
func (op_logs) Mixin() []ent.Mixin {
    return []ent.Mixin{}
}

// Indexes of the op_logs.
func (op_logs) Indexes() []ent.Index {
    return nil
}

// Annotations of the op_logs
func (op_logs) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "op_logs"},
	}
}
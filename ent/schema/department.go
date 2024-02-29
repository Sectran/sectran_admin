package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("The name of the department.").
			Annotations(entsql.WithComments(true)),

		field.String("area").
			NotEmpty().
			Comment("The area where the department is located.").
			Annotations(entsql.WithComments(true)),

		field.String("description").
			Comment("Description of the department.").
			Annotations(entsql.WithComments(true)),

		field.String("parent_departments_ids").
			NotEmpty().
			Comment("Comma-separated list of parent department IDs in ascending order.").
			Annotations(entsql.WithComments(true)),
	}
}

func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("departments"),
	}
}

func (Department) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

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
			Unique().
			MaxLen(32).
			Comment("The name of the department.|部门名称").
			Annotations(entsql.WithComments(true)),

		field.String("area").
			NotEmpty().
			Comment("The area where the department is located.|部门所在地区").
			MaxLen(128).
			Annotations(entsql.WithComments(true)),

		field.String("description").
			Comment("Description of the department.|部门描述").
			MaxLen(128).
			Annotations(entsql.WithComments(true)),

		field.String("parent_departments").
			NotEmpty().
			Comment("Comma-separated list of parent department IDs in ascending order.|上级部门集合逗号分隔升序排列").
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

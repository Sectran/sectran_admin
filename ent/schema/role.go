package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique().
			Comment("The name of the role.|角色名称").
			Annotations(entsql.WithComments(true)),
		field.Int("weight").
			Comment("The weight of the role. Smaller values indicate higher priority.|角色优先级，值越小优先级越高").
			Annotations(entsql.WithComments(true)),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("roles"),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

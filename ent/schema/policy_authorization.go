package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// PolicyAuth holds the schema definition for the PolicyAuth entity.
type PolicyAuth struct {
	ent.Schema
}

// Fields of the PolicyAuth.
func (PolicyAuth) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(64).
			Comment("policy name|策略名称").
			Annotations(entsql.WithComments(true)),
		field.Int32("power").
			Min(0).
			Comment("policy power|策略优先级").
			Annotations(entsql.WithComments(true)),
		field.Uint64("department_id").
			Optional().
			Min(1).
			Comment("ID of the policy's department.|策略所属部门").
			Annotations(entsql.WithComments(true)),
		field.String("users").
			NotEmpty().
			Comment("策略关联用户").
			Annotations(entsql.WithComments(true)),
		field.String("accounts").
			NotEmpty().
			Comment("策略关联账号").
			Annotations(entsql.WithComments(true)),
		field.Bool("direction").
			Comment("策略相关性方向,默认正向，即断言正向用户与账号").
			Default(true).
			Annotations(entsql.WithComments(true)),
	}
}

// Edges of the PolicyAuth.
func (PolicyAuth) Edges() []ent.Edge {
	return nil
}

// Mixin of the PolicyAuth.
func (PolicyAuth) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

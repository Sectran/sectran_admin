package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Policy holds the schema definition for the Policy entity.
type Policy struct {
	ent.Schema
}

// Fields of the Policy.
func (Policy) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(64).
			Comment("policy name|策略名称").
			Annotations(entsql.WithComments(true)),
		field.Int32("power").
			Optional().
			Min(0).
			Default(0).
			Comment("policy power|策略优先级、值越小优先级约高").
			Annotations(entsql.WithComments(true)),
		field.Uint64("department_id").
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
		field.Time("effecte_time_start").
			Optional().
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Policy effective time rangw start|策略生效时间开始").
			Annotations(entsql.WithComments(true)),
		field.Time("effecte_time_end").
			Optional().
			Default(time.Time{}.AddDate(9999, 0, 0)).
			UpdateDefault(time.Now).
			Comment("Policy effective time rangw end|策略生效时间结束").
			Annotations(entsql.WithComments(true)),
	}
}

// Edges of the Policy.
func (Policy) Edges() []ent.Edge {
	return nil
}

// Mixin of the Policy.
func (Policy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

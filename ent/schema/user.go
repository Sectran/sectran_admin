package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("account").
			Unique().
			NotEmpty().
			Comment("User account.").
			Annotations(entsql.WithComments(true)),
		field.String("name").
			NotEmpty().
			Comment("User name.").
			Annotations(entsql.WithComments(true)),
		field.String("password").
			Sensitive().
			Comment("User password.").
			Annotations(entsql.WithComments(true)),
		field.Uint64("department_id").
			Optional().
			Min(0).
			Comment("ID of the user's department.").
			Annotations(entsql.WithComments(true)),
		field.Uint64("role_id").
			Optional().
			Min(0).
			Comment("ID of the user's role.").
			Annotations(entsql.WithComments(true)),
		field.Bool("status").
			Default(true).
			Comment("User status (enabled(true) or disabled(false)).").
			Annotations(entsql.WithComments(true)),
		field.String("description").
			Optional().
			Comment("User description.").
			Annotations(entsql.WithComments(true)),
		field.String("email").
			Optional().
			Comment("User email.").
			Annotations(entsql.WithComments(true)),
		field.String("phone_number").
			Optional().
			Comment("User phone number.").
			Annotations(entsql.WithComments(true)),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("departments", Department.Type).Unique().Field("department_id"),
		edge.To("roles", Role.Type).Unique().Field("role_id"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

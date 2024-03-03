package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("The name of the device.").
			Annotations(entsql.WithComments(true)),
		field.String("host").
			Unique().
			NotEmpty().
			Comment("login host").
			Annotations(entsql.WithComments(true)),
		field.String("description").
			Comment("Description of the device.").
			Annotations(entsql.WithComments(true)),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("accounts", Account.Type).Ref("devices"),
	}
}

// Mixin of the Device.
func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

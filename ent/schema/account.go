package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

const (
	ProtocolSsh uint8 = 1 + iota
	PtotocolRdp
	PtotocolVnv
	PtotocolSftp
	PtotocolFtp
	PtotocolMysql
	PtotocolOracle
	PtotocolRedis
	ProtocolMax
)

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			NotEmpty().
			Comment("account username").
			Annotations(entsql.WithComments(true)),
		field.Uint32("port").
			Comment("account port").
			Annotations(entsql.WithComments(true)),
		field.Uint8("protocol").
			Min(ProtocolSsh).
			Max(ProtocolMax).
			Comment("protocol of the this account.").
			Annotations(entsql.WithComments(true)),
		field.String("password").
			Comment("account password").
			Annotations(entsql.WithComments(true)),
		field.String("private_key").
			Comment("private_key of the this account.").
			Annotations(entsql.WithComments(true)),
		field.Uint64("device_id").
			Optional().
			Comment("account belong to").
			Annotations(entsql.WithComments(true)),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type).Unique().Field("device_id"),
	}
}

// Mixin of the Account.
func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

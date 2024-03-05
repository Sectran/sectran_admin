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
			Comment("account username|账号名称").
			Annotations(entsql.WithComments(true)),
		field.Uint32("port").
			Comment("account port|端口").
			Annotations(entsql.WithComments(true)),
		field.Uint8("protocol").
			Min(ProtocolSsh).
			Max(ProtocolMax).
			Comment("protocol of the this account.|账号协议").
			Annotations(entsql.WithComments(true)),
		field.String("password").
			Comment("account password|账号密码").
			Annotations(entsql.WithComments(true)),
		field.String("private_key").
			Comment("private_key of the this account.|账号私钥").
			Annotations(entsql.WithComments(true)),
		field.Uint64("device_id").
			Optional().
			Comment("account belong to|账号所属设备").
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

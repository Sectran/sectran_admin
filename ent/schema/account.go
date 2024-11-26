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
			MaxLen(16).
			Annotations(entsql.WithComments(true)),
		field.Uint32("port").
			Comment("account port|端口").
			Min(1).Max(65534).
			Annotations(entsql.WithComments(true)),
		field.Uint8("protocol").
			Min(ProtocolSsh).Max(ProtocolMax).
			Comment("protocol of the this account.|账号协议").
			Annotations(entsql.WithComments(true)),
		field.String("password").
			Optional().
			Comment("account password|账号密码").
			MaxLen(128).
			Annotations(entsql.WithComments(true)),
		field.String("private_key").
			Optional().
			Comment("private_key of the this account.|账号私钥").
			MaxLen(4096).
			Annotations(entsql.WithComments(true)),
		field.String("private_key_password").
			Optional().
			Comment("private_key password of the this account.|私钥口令").
			MaxLen(4096).
			Annotations(entsql.WithComments(true)),
		field.Uint64("device_id").
			Comment("account belong to|账号所属设备").
			Min(1).
			Annotations(entsql.WithComments(true)),
		field.Uint64("department_id").
			Comment("account belong to|账号所属部门").
			Min(1).
			Annotations(entsql.WithComments(true)),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type).Required().Unique().Field("device_id"),
		edge.To("departments", Device.Type).Required().Unique().Field("department_id"),
	}
}

// Mixin of the Account.
func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

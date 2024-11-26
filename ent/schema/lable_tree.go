package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// LableTree holds the schema definition for the LableTree entity.
type LableTree struct {
	ent.Schema
}

const (
	LableTargetType_User uint8 = 0x2f + iota
	LableTargetType_Dept
	LableTargetType_Device
	LableTargetType_Account
	LableTargetType_Role
	LableTargetType_Log
	LableTargetType_Max
)

const (
	LableType_Group uint8 = 0x3f + iota
	LableType_Control
	LableType_Authorization
)

// Fields of the LableTree.
func (LableTree) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(64).
			Comment("lable name|标签名称").
			Annotations(entsql.WithComments(true)),
		field.Uint("type").
			Comment("lable type|标签类型（分组标签、控制标签、授权标签）").
			Annotations(entsql.WithComments(true)),
		field.String("icon").
			MaxLen(32).
			Comment("lable icon|标签图标").
			Annotations(entsql.WithComments(true)),
		field.String("content").
			MaxLen(1024).
			Comment("lable content|标签内容").
			Annotations(entsql.WithComments(true)),
		field.Uint8("ownership").
			Comment("lable ownership Level (Department Level/User Level)|标签所有权级别（部门级别/用户级别）").
			Annotations(entsql.WithComments(true)),
		field.Uint64("owner_id").
			Comment("lable owner,user ID,dept ID|标签所属者,用户ID,部门ID").
			Annotations(entsql.WithComments(true)),
		field.Uint64("parent_id").
			Comment("parent lable id|父标签id").
			Annotations(entsql.WithComments(true)),
		field.String("description").
			MaxLen(256).
			Comment("label description|标签描述").
			Annotations(entsql.WithComments(true)),
		//分组标签如果指定了目标类型、那么只能给这个类型的数据打标签，否则应该是混合类型（默认值）
		field.Uint16("target_type").
			Comment("lable target type|标签目标类型").
			Annotations(entsql.WithComments(true)),
		//限制树深不可超过64级
		field.String("parents").
			Comment("parent lables id,split by ',',lable tree deep cannot gather than 32|父标签id集合升序排列,逗号分隔,限制最多不可超过64级").
			Annotations(entsql.WithComments(true)),
		//子标签是否可以继承父标签,只有同类标签才可继承，继承标签不可以和父标签存在冲突属性（控制标签）注意标签不能形成循环依赖
		field.Bool("inherit").
			Comment("child lable can inherit parents|标签是否可以继承").
			Annotations(entsql.WithComments(true)),
	}
}

// Edges of the LableTree.
func (LableTree) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the LableTree.
func (LableTree) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

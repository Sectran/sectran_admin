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
		//分组标签、控制标签、授权标签
		field.Uint("type").
			Comment("lable type|标签类型").
			Annotations(entsql.WithComments(true)),
		field.String("icon").
			MaxLen(32).
			Comment("lable icon|标签图标").
			Annotations(entsql.WithComments(true)),
		field.String("content").
			MaxLen(1024).
			Comment("lable content|标签内容").
			Annotations(entsql.WithComments(true)),
		field.Uint64("parent_lable").
			Comment("parent lable id|父标签id").
			Annotations(entsql.WithComments(true)),
		//分组标签如果指定了目标类型、那么只能给这个类型的数据打标签
		field.Uint16("lable_target_type").
			Comment("lable target type|标签目标类型").
			Annotations(entsql.WithComments(true)),
		field.String("parent_lables").
			Comment("parent lables id,split by ','|父标签id集合升序排列,逗号分隔").
			Annotations(entsql.WithComments(true)),
		//标签所属者不可以转移
		field.Uint64("lable_owner").
			Comment("lable owner,user ID|标签所属者,用户ID").
			Annotations(entsql.WithComments(true)),
		//子标签是否可以继承父标签,只有同类标签才可继承，继承标签不可以和父标签存在冲突属性（控制标签）
		//注意标签不能形成循环依赖
		field.Bool("inherit").
			Comment("child lable can inherit parents|标签是否可以继承").
			Annotations(entsql.WithComments(true)),
		//关联标签,非排序
		field.String("related_lables").
			Comment("related labels id,split by ','|关联标签id集合升序排列,逗号分隔").
			Annotations(entsql.WithComments(true)),
		field.String("description").
			MaxLen(256).
			Comment("label description|标签描述").
			Annotations(entsql.WithComments(true)),
		field.String("ext1").
			Comment("label extented|标签拓展").
			Annotations(entsql.WithComments(true)),
		field.String("ext2").
			Comment("label extented|标签拓展").
			Annotations(entsql.WithComments(true)),
	}
}

// Edges of the LableTree.
func (LableTree) Edges() []ent.Edge {
	return nil
}

// Mixin of the LableTree.
func (LableTree) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

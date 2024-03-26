package types

import (
	"sectran_admin/ent"
)

// 登录请求
type LoginReq struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// 登录相应
type LoginRes struct {
	Token    string       `json:"token"`
	User     *ent.User    `json:"user"`
	Base     *BaseMsgResp `json:"base"`
	DeptName string       `json:"deptName"`
	RoleName string       `json:"roleName"`
}

// 查询树形结构的所有子节点
type ChildrenReq struct {
	PageInfo
	Deep int    `json:"deep" validate:"number"`
	Id   uint64 `json:"id" validate:"number"`
	// The name of the department.|部门名称
	Name *string `json:"name,optional"`
	// The area where the department is located.|部门所在地区
	Area *string `json:"area,optional"`
	// Description of the department.|部门描述
	Description *string `json:"description,optional"`
}

// The response data of api authorization | API授权数据
// swagger:model ApiAuthorityInfo
type ApiAuthorityInfo struct {
	// API path | API 路径
	// required : true
	// max length : 80
	Path string `json:"path" validate="required,max=80"`
	// API method | API请求方法
	// required : true
	// min length : 3
	// max length : 4
	Method string `json:"method" validate="required,min=3,max=4"`
}

// Create or update api authorization information request | 创建或更新API授权信息
// swagger:model UpdateApiAuthorityReq
type UpdateApiAuthorityReq struct {
	// Role ID | 角色ID
	// required : true
	// max : 1000
	RoleId uint64 `json:"roleId" validate:"required,lt=1000"`
	// API authorization list | API授权列表数据
	// Required: true
	Data []ApiAuthorityInfo `json:"data"`
}

type MenuListInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

type AuthorityRequestInfo struct {
	Id   uint64 `json:"id"`
	Type uint8  `json:"type"` //0 mune 1 api 2 all
}

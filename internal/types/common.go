package types

import "sectran_admin/ent"

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

// The response data of api authorization list | API授权列表返回数据
// swagger:model ApiAuthorityListResp
type ApiAuthorityListResp struct {
	BaseDataInfo
	// The api authorization list data | API授权列表数据
	Data ApiAuthorityListInfo `json:"data"`
}

// The  data of api authorization list | API授权列表数据
// swagger:model ApiAuthorityListInfo
type ApiAuthorityListInfo struct {
	BaseListInfo
	// The api authorization list data | API授权列表数据
	Data []ApiAuthorityInfo `json:"data"`
}

// The response data of department information | Department信息
// swagger:model DepartmentInfo
type DepartmentInfoDTO struct {
	DepartmentInfo
	HasChildren bool `json:"hasChildren,optional"`
}

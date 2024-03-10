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
	Id uint64 `json:"id" validate:"number"`
}

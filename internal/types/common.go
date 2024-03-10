package types

import "sectran_admin/ent"

type LoginReq struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type LoginRes struct {
	Token    string       `json:"token"`
	User     *ent.User    `json:"user"`
	Base     *BaseMsgResp `json:"base"`
	DeptName string       `json:"deptName"`
	RoleName string       `json:"roleName"`
}

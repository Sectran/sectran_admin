package types

// ----------------------------部门----------------------------
// The response data of department information | Department信息
// swagger:model DepartmentInfo
type DepartmentInfoRefer struct {
	DepartmentInfo
	HasChildren bool `json:"hasChildren,optional"` //是否存在下级部门
}

// Department list data | Department列表数据
// swagger:model DepartmentListInfo
type DepartmentListInfoRefer struct {
	BaseListInfo
	// The API list data | Department列表数据
	Data []DepartmentInfoRefer `json:"data"`
}

// The response data of department list | Department列表数据
// swagger:model DepartmentListResp
type DepartmentListRespRefer struct {
	BaseDataInfo
	// Department list data | Department列表数据
	Data DepartmentListInfoRefer `json:"data"`
}

// Get department list request params | Department列表请求参数
// swagger:model DepartmentListReq
type DepartmentListReqRefer struct {
	DepartmentListReq
	// the parent departmenr id |父部门id
	ParentDeptId *uint64 `json:"parentDeptId,optional"`
	// 查询一级子部门或者ParentDeptId部门下所有数据
	Flag *uint8 `json:"flag,optional"`
}

//----------------------------账号----------------------------

// Get account list request params | Account列表请求参数
// swagger:model AccountListReq
type AccountListReqRefer struct {
	PageInfo
	// account username|账号名称
	Username *string `json:"username,optional"`
	//所属设备id
	DeviceId *uint64 `json:"deviceId,optional"`
	//账号协议
	Protocol *uint8 `json:"protocol,optional"`
	//账号端口
	Port *uint32 `json:"port,optional"`
}

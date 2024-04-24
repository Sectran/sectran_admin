package types

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

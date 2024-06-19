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

// ----------------------------用户----------------------------
// User information response | User信息返回体

// User information response | User信息返回体
// swagger:model UserInfoResp
type UserInfoRespRefer struct {
	BaseDataInfo
	// User information | User数据
	Data UserInfoRefer `json:"data"`
}

// The response data of user information | User信息
// swagger:model UserInfo
type UserInfoRefer struct {
	BaseIDInfo
	// User account.|用户账号
	Account *string `json:"account,optional"`
	// User name.|用户姓名
	Name *string `json:"name,optional"`
	// User password.|用户密码
	Password *string `json:"password,optional"`
	// ID of the user's department.|用户所属部门
	DepartmentId *uint64 `json:"departmentId,optional"`
	// ID of the user's role.|用户所属角色
	RoleId *uint64 `json:"roleId,optional"`
	// User status (enabled(true) or disabled(false)).|用户账号状态
	Status *bool `json:"status,optional"`
	// User description.|用户账号描述
	Description *string `json:"description,optional"`
	// User email.|用户邮箱
	Email *string `json:"email,optional"`
	// User phone number.|用户手机号码
	PhoneNumber    *string `json:"phoneNumber,optional"`
	RoleName       *string `json:"roleName"`
	DepartmentName *string `json:"departmentName"`
}

// User list data | User列表数据
// swagger:model UserListInfo
type UserListInfoRefer struct {
	BaseListInfo
	// The API list data | User列表数据
	Data []UserInfoRefer `json:"data"`
}

// The response data of user list | User列表数据
// swagger:model UserListResp
type UserListRespRefer struct {
	BaseDataInfo
	// User list data | User列表数据
	Data UserListInfoRefer `json:"data"`
}

// Get user list request params | User列表请求参数
// swagger:model UserListReq
type UserListReqRefer struct {
	PageInfo
	// User account.|用户账号
	Account *string `json:"account,optional"`
	// User name.|用户姓名
	Name *string `json:"name,optional"`
	// User status (enabled(true) or disabled(false)).|用户账号状态
	Status *bool `json:"status,optional"`
	// User description.|用户账号描述
	Description *string `json:"description,optional"`
	// User email.|用户邮箱
	Email *string `json:"email,optional"`
	// User phone number.|用户手机号码
	PhoneNumber    *string `json:"phoneNumber,optional"`
	RoleName       *string `json:"roleName,optional"`
	DepartmentName *string `json:"departmentName,optional"`
}

// The response data of device list | Device列表数据
// swagger:model DeviceListResp
type DeviceListRespRefer struct {
	BaseDataInfo
	// Device list data | Device列表数据
	Data DeviceListInfoRefer `json:"data"`
}

// Device list data | Device列表数据
// swagger:model DeviceListInfo
type DeviceListInfoRefer struct {
	BaseListInfo
	// The API list data | Device列表数据
	Data []DeviceInfoRefer `json:"data"`
}

// The response data of device information | Device信息
// swagger:model DeviceInfo
type DeviceInfoRefer struct {
	BaseIDInfo
	// The name of the device.|设备名称
	Name *string `json:"name,optional"`
	// ID of the device's department.|设备所属部门
	DepartmentId *uint64 `json:"departmentId,optional"`
	// login host|设备地址
	Host *string `json:"host,optional"`
	// type of the device.|设备类型
	Type *string `json:"type,optional"`
	// Description of the device.|设备描述
	Description *string `json:"description,optional"`
	// the name of device department|设备所属部门名称
	DeptName *string `json:"deptName,optional"`
}

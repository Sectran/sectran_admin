// Code generated by goctl. DO NOT EDIT.
package types

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// required : true
	// min : 0
	Page uint64 `json:"page" validate:"required,number,gt=0"`
	// Page size | 单页数据行数
	// required : true
	// max : 100000
	PageSize uint64 `json:"pageSize" validate:"required,number,lt=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic UUID request in path | 基础UUID地址参数请求
// swagger:model UUIDPathReq
type UUIDPathReq struct {
	// ID
	// Required: true
	Id string `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// required : true
	// max length : 36
	// min length : 36
	Id string `json:"id" validate:"required,len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	Ids []string `json:"ids"`
}

// The base ID response data | 基础ID信息
// swagger:model BaseIDInfo
type BaseIDInfo struct {
	// ID
	Id *uint64 `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The base UUID response data | 基础UUID信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id *string `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The response data of department information | Department信息
// swagger:model DepartmentInfo
type DepartmentInfo struct {
	BaseIDInfo
	// The name of the department.|部门名称
	Name *string `json:"name,optional"`
	// The area where the department is located.|部门所在地区
	Area *string `json:"area,optional"`
	// Description of the department.|部门描述
	Description *string `json:"description,optional"`
	// parent department ID.|父亲部门id
	ParentDepartmentId *uint64 `json:"parentDepartmentId,optional"`
	// Comma-separated list of parent department IDs in ascending order.|上级部门集合逗号分隔升序排列
	ParentDepartments *string `json:"parentDepartments,optional"`
}

// The response data of department list | Department列表数据
// swagger:model DepartmentListResp
type DepartmentListResp struct {
	BaseDataInfo
	// Department list data | Department列表数据
	Data DepartmentListInfo `json:"data"`
}

// Department list data | Department列表数据
// swagger:model DepartmentListInfo
type DepartmentListInfo struct {
	BaseListInfo
	// The API list data | Department列表数据
	Data []DepartmentInfo `json:"data"`
}

// Get department list request params | Department列表请求参数
// swagger:model DepartmentListReq
type DepartmentListReq struct {
	PageInfo
	// The name of the department.|部门名称
	Name *string `json:"name,optional"`
	// The area where the department is located.|部门所在地区
	Area *string `json:"area,optional"`
	// Description of the department.|部门描述
	Description *string `json:"description,optional"`
}

// Department information response | Department信息返回体
// swagger:model DepartmentInfoResp
type DepartmentInfoResp struct {
	BaseDataInfo
	// Department information | Department数据
	Data DepartmentInfo `json:"data"`
}

// The response data of role information | Role信息
// swagger:model RoleInfo
type RoleInfo struct {
	BaseIDInfo
	// The name of the role.|角色名称
	Name *string `json:"name,optional"`
	// The weight of the role. Smaller values indicate higher priority.|角色优先级，值越小优先级越高
	Weight *int `json:"weight,optional"`
}

// The response data of role list | Role列表数据
// swagger:model RoleListResp
type RoleListResp struct {
	BaseDataInfo
	// Role list data | Role列表数据
	Data RoleListInfo `json:"data"`
}

// Role list data | Role列表数据
// swagger:model RoleListInfo
type RoleListInfo struct {
	BaseListInfo
	// The API list data | Role列表数据
	Data []RoleInfo `json:"data"`
}

// Get role list request params | Role列表请求参数
// swagger:model RoleListReq
type RoleListReq struct {
	PageInfo
	// The name of the role.|角色名称
	Name *string `json:"name,optional"`
}

// Role information response | Role信息返回体
// swagger:model RoleInfoResp
type RoleInfoResp struct {
	BaseDataInfo
	// Role information | Role数据
	Data RoleInfo `json:"data"`
}

// The response data of user information | User信息
// swagger:model UserInfo
type UserInfo struct {
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
	PhoneNumber *string `json:"phoneNumber,optional"`
}

// The response data of user list | User列表数据
// swagger:model UserListResp
type UserListResp struct {
	BaseDataInfo
	// User list data | User列表数据
	Data UserListInfo `json:"data"`
}

// User list data | User列表数据
// swagger:model UserListInfo
type UserListInfo struct {
	BaseListInfo
	// The API list data | User列表数据
	Data []UserInfo `json:"data"`
}

// Get user list request params | User列表请求参数
// swagger:model UserListReq
type UserListReq struct {
	PageInfo
	// User account.|用户账号
	Account *string `json:"account,optional"`
	// User name.|用户姓名
	Name *string `json:"name,optional"`
	// User password.|用户密码
	Password *string `json:"password,optional"`
}

// User information response | User信息返回体
// swagger:model UserInfoResp
type UserInfoResp struct {
	BaseDataInfo
	// User information | User数据
	Data UserInfo `json:"data"`
}

// The response data of account information | Account信息
// swagger:model AccountInfo
type AccountInfo struct {
	BaseIDInfo
	// account username|账号名称
	Username *string `json:"username,optional"`
	// account port|端口
	Port *uint32 `json:"port,optional"`
	// protocol of the this account.|账号协议
	Protocol *uint8 `json:"protocol,optional"`
	// account password|账号密码
	Password *string `json:"password,optional"`
	// private_key of the this account.|账号私钥
	PrivateKey *string `json:"privateKey,optional"`
	// private_key password of the this account.|私钥口令
	PrivateKeyPassword *string `json:"privateKeyPassword,optional"`
	// account belong to|账号所属设备
	DeviceId *uint64 `json:"deviceId,optional"`
	// account belong to|账号所属部门
	DepartmentId *uint64 `json:"departmentId,optional"`
}

// The response data of account list | Account列表数据
// swagger:model AccountListResp
type AccountListResp struct {
	BaseDataInfo
	// Account list data | Account列表数据
	Data AccountListInfo `json:"data"`
}

// Account list data | Account列表数据
// swagger:model AccountListInfo
type AccountListInfo struct {
	BaseListInfo
	// The API list data | Account列表数据
	Data []AccountInfo `json:"data"`
}

// Get account list request params | Account列表请求参数
// swagger:model AccountListReq
type AccountListReq struct {
	PageInfo
	// account username|账号名称
	Username *string `json:"username,optional"`
	// account password|账号密码
	Password *string `json:"password,optional"`
	// private_key of the this account.|账号私钥
	PrivateKey *string `json:"privateKey,optional"`
}

// Account information response | Account信息返回体
// swagger:model AccountInfoResp
type AccountInfoResp struct {
	BaseDataInfo
	// Account information | Account数据
	Data AccountInfo `json:"data"`
}

// The response data of device information | Device信息
// swagger:model DeviceInfo
type DeviceInfo struct {
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
}

// The response data of device list | Device列表数据
// swagger:model DeviceListResp
type DeviceListResp struct {
	BaseDataInfo
	// Device list data | Device列表数据
	Data DeviceListInfo `json:"data"`
}

// Device list data | Device列表数据
// swagger:model DeviceListInfo
type DeviceListInfo struct {
	BaseListInfo
	// The API list data | Device列表数据
	Data []DeviceInfo `json:"data"`
}

// Get device list request params | Device列表请求参数
// swagger:model DeviceListReq
type DeviceListReq struct {
	PageInfo
	// The name of the device.|设备名称
	Name *string `json:"name,optional"`
	// login host|设备地址
	Host *string `json:"host,optional"`
	// type of the device.|设备类型
	Type *string `json:"type,optional"`
}

// Device information response | Device信息返回体
// swagger:model DeviceInfoResp
type DeviceInfoResp struct {
	BaseDataInfo
	// Device information | Device数据
	Data DeviceInfo `json:"data"`
}

// The response data of lable tree information | LableTree信息
// swagger:model LableTreeInfo
type LableTreeInfo struct {
	BaseIDInfo
	// lable name|标签名称
	Name *string `json:"name,optional"`
	// lable type|标签类型
	Type *uint `json:"type,optional"`
	// lable icon|标签图标
	Icon *string `json:"icon,optional"`
	// lable content|标签内容
	Content *string `json:"content,optional"`
	// parent lable id|父标签id
	ParentLable *uint64 `json:"parentLable,optional"`
	// lable target type|标签目标类型
	LableTargetType *uint16 `json:"lableTargetType,optional"`
	// parent lables id,split by ','|父标签id集合升序排列,逗号分隔
	ParentLables *string `json:"parentLables,optional"`
	// lable owner,user ID|标签所属者,用户ID
	LableOwner *uint64 `json:"lableOwner,optional"`
	// child lable can inherit parents|标签是否可以继承
	Inherit *bool `json:"inherit,optional"`
	// related labels id,split by ','|关联标签id集合升序排列,逗号分隔
	RelatedLables *string `json:"relatedLables,optional"`
	// label description|标签描述
	Description *string `json:"description,optional"`
	// label extented|标签拓展
	Ext1 *string `json:"ext1,optional"`
	// label extented|标签拓展
	Ext2 *string `json:"ext2,optional"`
}

// The response data of lable tree list | LableTree列表数据
// swagger:model LableTreeListResp
type LableTreeListResp struct {
	BaseDataInfo
	// LableTree list data | LableTree列表数据
	Data LableTreeListInfo `json:"data"`
}

// LableTree list data | LableTree列表数据
// swagger:model LableTreeListInfo
type LableTreeListInfo struct {
	BaseListInfo
	// The API list data | LableTree列表数据
	Data []LableTreeInfo `json:"data"`
}

// Get lable tree list request params | LableTree列表请求参数
// swagger:model LableTreeListReq
type LableTreeListReq struct {
	PageInfo
	// lable name|标签名称
	Name *string `json:"name,optional"`
	// lable icon|标签图标
	Icon *string `json:"icon,optional"`
	// lable content|标签内容
	Content *string `json:"content,optional"`
}

// LableTree information response | LableTree信息返回体
// swagger:model LableTreeInfoResp
type LableTreeInfoResp struct {
	BaseDataInfo
	// LableTree information | LableTree数据
	Data LableTreeInfo `json:"data"`
}

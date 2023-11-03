package types

const (
	REQUEST_SUCCESS int = 200 + iota //请求成功
)

const (
	ERROR_REUQEST_FAILED    int = 500 + iota //请求失败
	ERROR_SUBJECT_NOT_EXSIT                  //请求主体不存在
	ERROR_ILLEGAL_PARAMS                     //不合法参数
	ERROR_ISSUE_TOKEN                        //颁发token失败
)

// -----------------response---------------
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CommonResponse struct {
	Response
	Data any `json:"data"`
}

func BuildCommonResponse(data any, msg string, code int) *CommonResponse {
	return &CommonResponse{
		Response: Response{
			Msg:  msg,
			Code: code,
		},
		Data: data,
	}
}

// -----------------auth---------------

type AuthRequest struct {
	Account  string `json:"account" validate:"required,alphanum,min=0,max=255"`
	Password string `json:"password" validate:"require,ascii"`
}

// -----------------page---------------
type PageInfo struct {
	PageStart int `json:"pageStart" validate:"required,gte=1"`
	PageEnd   int `json:"pageEnd" validate:"required,gte=1"`
	PageSize  int `json:"pageSize" validate:"required,gte=1"`
}

// -----------------user---------------
type UserDeleteRequest struct {
	UserIds []int `json:"user_ids"`
}

// https://github.com/go-playground/validator/blob/master/README.md#network
type UserVisibleInsertInfo struct {
	UserId      int64  `json:"userId" validate:"required,gte=0"`                   // 用户ID (自动填充)
	Account     string `json:"account" validate:"required,alphanum,min=5,max=255"` // 用户账号
	Username    string `json:"username" validate:"required,min=5,max=255"`         // 用户姓名
	DeptId      int64  `json:"deptId" validate:"required,gte=0"`                   // 用户所属部门ID
	Disable     uint8  `json:"disable" validate:"oneof=0 1"`                       // 账号是否禁用
	Description string `json:"description" validate:"required,min=0,max=255"`      // 账号描述
	Email       string `json:"email" validate:"email"`                             // 用户邮箱
	Telephone   string `json:"telephone" validate:"e164"`                          // 用户电话
	CreateTime  string `json:"createTime"`                                         // 创建时间 (自动填充)
	RoleId      int64  `json:"roleId" validate:"required,gte=0"`                   // 用户角色ID
}

type UserVisibleQueryInfo struct {
	UserId      int64  `json:"userId" validate:"gte=-1"`             // 用户ID
	Account     string `json:"account" validate:"min=0,max=255"`     // 用户账号
	Username    string `json:"username" validate:"min=0,max=255"`    // 用户姓名
	DeptId      int64  `json:"deptId" validate:"gte=-1"`             // 用户所属部门ID
	Disable     uint8  `json:"disable" validate:"oneof=0 1"`         // 账号是否禁用
	Description string `json:"description" validate:"min=0,max=255"` // 账号描述
	Email       string `json:"email" validate:"min=0,max=50"`        // 用户邮箱
	Telephone   string `json:"telephone" validate:"min=0,max=20"`    // 用户电话
	CreateTime  string `json:"createTime" validate:"min=0,max=20"`   // 创建时间
	RoleId      int64  `json:"roleId" validate:"gte=-1"`             // 用户角色ID
}

type UserQueryInfo struct {
	UserVisibleQueryInfo
	PageInfo
}

type UserAllInfo struct {
	UserVisibleInsertInfo
	CreateByUid int64  `json:"createByUid" validate:"gte=0"`       // 创建人
	IsDeleted   uint8  `json:"isDeleted"  validate:"oneof=0 1"`    // 是否被删除
	Password    string `json:"password" validate:"required,ascii"` // 用户密码
}

// -----------------role---------------
type RoleDeleteRequest struct {
	RoleIds []int `json:"RoleIds"`
}

type RoleQueryInfo struct {
	RoleVisibleQueryInfo
	PageInfo
}

type RoleVisibleQueryInfo struct {
	UserId      int64  `json:"userId" validate:"gte=-1"`             // 用户ID
	Account     string `json:"account" validate:"min=0,max=255"`     // 用户账号
	Username    string `json:"username" validate:"min=0,max=255"`    // 用户姓名
	DeptId      int64  `json:"deptId" validate:"gte=-1"`             // 用户所属部门ID
	Disable     uint8  `json:"disable" validate:"oneof=0 1"`         // 账号是否禁用
	Description string `json:"description" validate:"min=0,max=255"` // 账号描述
	Email       string `json:"email" validate:"min=0,max=50"`        // 用户邮箱
	Telephone   string `json:"telephone" validate:"min=0,max=20"`    // 用户电话
	CreateTime  string `json:"createTime" validate:"min=0,max=20"`   // 创建时间
	RoleId      int64  `json:"roleId" validate:"gte=-1"`             // 用户角色ID
}

type RoleVisibleInfo struct {
	PageInfo
	RoleId      int64  `json:"roleId"`      // 角色ID
	Name        string `json:"name"`        // 角色名称
	Description string `json:"description"` // 角色描述
	CreateTime  string `json:"createTime"`  // 创建时间
}

type RoleAllInfo struct {
	RoleVisibleInfo
	CreateByUid int64 `json:"createByUid"` // 创建者
	IsDelete    uint8 `json:"isDeleted"`   // 是否被删除
}

type RoleVisibleInfoArray struct {
	Response
	Data []RoleVisibleInfo `json:"data"`
}

// -----------------department---------------
type DeptDeleteRequest struct {
	DeptIds []int `json:"deptIds"`
}

type DeptVisibleInfo struct {
	PageInfo
	DeptId      int64  `json:"deptId"`      // 部门ID
	Name        string `json:"name"`        // 部门名称
	Description string `json:"description"` // 部门描述
	ParentId    int64  `json:"parentId"`    // 上级部门ID
	ChildIds    string `json:"childIds"`    // 下级部门ID集合，用逗号分隔
	Region      string `json:"region"`      // 部门所在地区
}

type DeptAllInfo struct {
	DeptVisibleInfo
	IsDelete    uint8 `json:"isDeleted"`   // 是否被删除
	CreateByUid int64 `json:"createByUid"` // 创建者
}

type DeptVisibleInfoArray struct {
	Response
	Data []DeptVisibleInfo `json:"data"`
}

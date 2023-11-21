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

type Pagination struct {
	PageNum    int64 `json:"pageNum" validate:"required,gte=0"`    //查询或者响应的开始页数
	PageSize   int64 `json:"pageSize" validate:"required,gte=1"`   //查询或者响应的页数大小
	TotalPages int64 `json:"totalPages" validate:"required,gte=1"` //查询或者响应的总页数
}

// -----------------auth---------------
type AuthRequest struct {
	Account  string `json:"account" validate:"required,alphanum,min=0,max=255"`
	Password string `json:"password" validate:"require,ascii"`
}

// -----------------user---------------
type UserDeleteRequest struct {
	UserIds []int `json:"user_ids"`
}

// https://github.com/go-playground/validator/blob/master/README.md#network
type UserInsertInfo struct {
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

type UserQueryInfo struct {
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

type UserAllInfo struct {
	CreateByUid int64  `json:"createByUid" validate:"gte=0"`       // 创建人
	IsDeleted   uint8  `json:"isDeleted"  validate:"oneof=0 1"`    // 是否被删除
	Password    string `json:"password" validate:"required,ascii"` // 用户密码
}

// -----------------role---------------
type RoleDeleteRequest struct {
	RoleId int64 `json:"role_id"  validate:"required,gte=0"` // 角色ID
}

type RoleQueryInfo struct {
	RoleVisibleQueryInfo
	PageInfo
}

type RoleVisibleQueryInfo struct {
	RoleId     int64  `json:"role_id" validate:"gte=-1"`          //角色ID
	Name       string `json:"name" validate:"min=0,max=255"`      //角色名称
	CreateTime string `json:"createTime" validate:"min=0,max=20"` // 创建时间

}

type RoleVisibleInfo struct {
	//RoleId      int64  `json:"role_id"  validate:"required,gte=0"`   // 角色ID
	Name        string `json:"name"  validate:"min=0,max=255"`       // 角色名称
	Description string `json:"description" validate:"min=0,max=255"` // 角色描述
}

type RoleAllInfo struct {
	RoleVisibleInfo
	CreateByUid int64  `json:"createByUid" validate:"gte=0"` // 创建者
	CreateTime  string `json:"createTime"  validate:"-"`     // 创建时间
	//IsDelete    uint8 `json:"isDeleted"`                    // 是否被删除
}

type RoleEditInfo struct {
	RoleId      int64  `json:"role_id"  validate:"required,gte=1"`   // 角色ID
	Name        string `json:"name"  validate:"min=1,max=255"`       // 角色名称
	Description string `json:"description" validate:"min=1,max=255"` // 角色描述
	CreateByUid int64  `json:"createByUid" validate:"gte=0"`         // 创建者
}

type RoleVisibleInfoArray struct {
	Response
	Data []RoleVisibleInfo `json:"data"`
}

// -----------------department---------------
type DeptAddRequest struct {
	Name        string `json:"name"  validate:"min=0,max=255"`        // 部门名称
	Description string `json:"description"  validate:"min=0,max=255"` // 部门描述
	ParentId    int64  `json:"parentId"  validate:"gte=0"`            // 上级部门ID
	ChildIds    string `json:"childIds"`                              // 下级部门ID集合，用逗号分隔
	Region      string `json:"region"`                                // 部门所在地区
	CreateByUid int64  `json:"createByUid" validate:"gte=0"`          // 创建者
}

type DeptQueryInfo struct {
	PageInfo
	DeptId   int64  `json:"dept_id" validate:"gte=-1"`       //角色ID
	Name     string `json:"name" validate:"min=0,max=255"`   //角色名称
	ParentId int64  `json:"parentId"  validate:"gte=-1"`     // 上级部门ID
	Region   string `json:"region" validate:"min=0,max=255"` // 部门所在地区
}

type DeptEditInfo struct {
	DeptId      int64  `json:"dept_id" validate:"gte=1"`             //部门ID
	Name        string `json:"name"  validate:"min=1,max=255"`       // 部门名称
	Description string `json:"description" validate:"min=1,max=255"` // 部门描述
	//ParentId    int64  `json:"parentId"  validate:"gte=0"`           // 上级部门ID
	Region string `json:"region"` // 部门所在地区
	//ChildIds    string `json:"childIds"`                             // 下级部门ID集合，用逗号分隔
	//CreateByUid int64  `json:"createByUid" validate:"gte=0"`         // 创建者
}

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
	//IsDelete    uint8 `json:"isDeleted"`   // 是否被删除
	CreateByUid int64 `json:"createByUid"` // 创建者
}

type DeptVisibleInfoArray struct {
	Response
	Data []DeptVisibleInfo `json:"data"`
}

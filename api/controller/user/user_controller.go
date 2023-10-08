package UserController

import (
	"github.com/gin-gonic/gin"
	response "sectran/api"
	"sectran/api/common"
)

// ListUser 查询用户列表
func ListUser(c *gin.Context) {
	if err, table, total := listUserImpl(c); err != nil {
		response.RequestError(c, "查询失败")
	} else {
		data := common.TableDto{Table: table, Total: total}
		response.RequestOk(c, data, "查询成功")
	}
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	p := common.UserDto{}
	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, "请输入参数")
		return
	}
	if !common.EmailRegexp(p.UserName) {
		response.RequestError(c, "请输入正确的用户名格式(邮箱)")
		return
	}
	if err, msg := addUserImpl(p); err != nil {
		response.RequestError(c, msg)
	} else {
		response.RequestOk(c, nil, msg)
	}
}

type EditDepartmentParameter struct {
	Id string `json:"id" gorm:"type:char(36);primary_key"` //用户ID
	common.UserDto
}

// EditDepartment 修改部门

func EditUser(c *gin.Context) {
	p := EditDepartmentParameter{}
	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, "请输入参数")
		return
	}
	if err := editUserImpl(p); err != nil {
		response.RequestError(c, "修改失败")
	} else {
		response.RequestOk(c, nil, "修改成功")
	}
}

// DeleteUser 删除部门
func DeleteUser(c *gin.Context) {
	p := common.DeleteDto{}
	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, "请输入参数")
		return
	}
	if len(p.Id) == 0 {
		response.RequestError(c, "请输入id")
		return
	}
	if err := deleteUserImpl(p); err != nil {
		response.RequestError(c, "删除失败")
	} else {
		response.RequestOk(c, nil, "删除成功")
	}
}

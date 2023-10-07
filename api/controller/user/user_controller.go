package user

import (
	"github.com/gin-gonic/gin"
	response "sectran/api"
)

type userParameter struct {
	UserName string `json:"userName"` //用户名
	Password string `json:"password"` //密码
}

// AddUser 添加部门
func AddUser(c *gin.Context) {
	p := userParameter{}
	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, "请输入参数")
		return
	}
	response.RequestOk(c, p, "添加成功")
	//if err := addDepartmentImpl(p); err != nil {
	//	response.RequestError(c, "添加失败")
	//} else {
	//	response.RequestOk(c, nil, "添加成功")
	//}
}

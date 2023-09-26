package router

import (
	"github.com/gin-gonic/gin"
	department "sectran/api/controller/department"
)

func DepartmentRouter(e *gin.Engine) {
	login := e.Group("department")
	//登录接口
	{
		login.POST("/list", department.List)
	}
}

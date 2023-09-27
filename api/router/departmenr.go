package router

import (
	"github.com/gin-gonic/gin"
	"sectran/api/controller/department"
)

func DepartmentRouter(e *gin.Engine) {

	login := e.Group("department")
	//请求列表
	{
		login.POST("/list", department.ListDepartment)
	}
	{
		login.POST("/add", department.AddDepartment)
	}
}

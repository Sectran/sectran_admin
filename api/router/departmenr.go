package router

import (
	"github.com/gin-gonic/gin"
	"sectran/api/controller/department"
)

func DepartmentRouter(e *gin.Engine) {

	login := e.Group("department")
	{
		login.POST("/add", department.AddDepartment)
	}

	//登录接口
	{
		login.POST("/list", department.List)
	}

}

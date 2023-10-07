package router

import (
	"github.com/gin-gonic/gin"
	"sectran/api/controller/user"
)

func UserRouter(e *gin.Engine) {
	User := e.Group("user")
	//请求列表
	//login.GET("/list", department.ListDepartment)
	User.POST("/add", user.AddUser)
	//login.POST("/edit", department.EditDepartment)
	//login.POST("/delete", department.DeleteDepartment)

}

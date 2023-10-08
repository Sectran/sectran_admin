package router

import (
	"github.com/gin-gonic/gin"
	UserController "sectran/api/controller/user"
)

func UserRouter(e *gin.Engine) {
	User := e.Group("user")
	//请求列表
	User.GET("/list", UserController.ListUser)
	User.POST("/add", UserController.AddUser)
	User.POST("/edit", UserController.EditUser)
	User.POST("/delete", UserController.DeleteUser)
}

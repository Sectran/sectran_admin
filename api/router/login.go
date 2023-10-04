package router

import (
	"github.com/gin-gonic/gin"
	loginController "sectran/api/controller/login"
)

func LoginRouter(e *gin.Engine) {
	login := e.Group("login")
	//登录接口
	{
		login.POST("/index", loginController.PostUserLogin)
	}
}

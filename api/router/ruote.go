package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sectran/api/common"
)

func InitRouter() {
	e := gin.Default()

	LoginRouter(e)
	e.Use(common.JWTAuthMiddleware())
	DepartmentRouter(e)
	UserRouter(e)
	e.GET("/userInfo", UserInfo)
	err := e.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func UserInfo(c *gin.Context) {
	//user, _ := c.Get("user")
	//claims := user.(*common.UserClaims)
	data := common.UserDto{UserName: c.GetString("username")}
	c.JSON(http.StatusOK, data)
}

package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"sectran/api/common"
)

func InitRouter() {
	e := gin.Default()
	LoginRouter(e)
	e.Use(common.JWTAuthMiddleware())
	DepartmentRouter(e)
	UserRouter(e)
	err := e.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

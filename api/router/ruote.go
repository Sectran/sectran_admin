package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	e := gin.Default()
	LoginRouter(e)

	err := e.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

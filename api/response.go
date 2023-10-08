package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
}

var Default = &response{}

func RequestOk(c *gin.Context, data interface{}, msg string) {
	fmt.Print("ok")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
		"msg":  msg,
	})
}

func RequestError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  msg,
	})
}

func TokenError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusUnauthorized,
		"msg":  msg,
	})
}

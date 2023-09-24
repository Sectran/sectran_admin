package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
}

var Default = &response{}

func RequestOk(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
		"msg":  msg,
	})
}

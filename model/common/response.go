package common

import (
	"net/http"

	"github.com/Sectran/sectran_pr/common"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int
	Data interface{}
	Msg  string
}

const (
	ERROR            = -1
	SUCCESS          = 0
	InvalidParameter = 100
	FailedOperation  = 101
)

// 数据异常
var ErrInvalidParameter = common.NewError(InvalidParameter, "invalid parameter")
var ErrFailedOperation = common.NewError(FailedOperation, "failed operation")

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "Success.", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "Success.", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "Operate Failed.", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithErr(c *gin.Context, err common.Error) {
	Result(ERROR, map[string]interface{}{}, err.Msg, c)
}

// 返回列表数据
func ResponseList(c *gin.Context, data interface{}, total int64) {
	res := &Response{
		Code: SUCCESS,
	}
	res.Data = map[string]interface{}{
		"Total": total, // data 不能为nil
		"Data":  data,
	}

	c.JSON(http.StatusOK, res)
	return
}

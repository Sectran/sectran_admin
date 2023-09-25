package user

import (
	"github.com/gin-gonic/gin"
	response "sectran/api"
)

//func UserList() string {
//	u, err := GetUserById(1)
//	if err != nil {
//		return api.ResponseError("找不到这个用户")
//	}
//
//	s, err := api.ResponseMsg(api.RSP_SUCCECC, "查找用户成功", u)
//	if err != nil {
//		return api.ResponseError("内部错误")
//	}
//
//	return s
//}

type loginParameter struct {
	UserName string `json:"userName"` //用户名
	Password string `json:"password"` //密码
}

func Login(c *gin.Context) {
	p := loginParameter{}
	err := c.BindJSON(&p)
	if err != nil {
		response.RequestError(c, nil, "请输入密码账号")
		return
	}

}

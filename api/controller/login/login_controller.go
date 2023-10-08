package loginController

import (
	"github.com/gin-gonic/gin"
	response "sectran/api"
	"sectran/api/common"
)

type loginParameter struct {
	UserName string `json:"userName"` //用户名
	Password string `json:"password"` //密码
}

func PostUserLogin(c *gin.Context) {
	p := loginParameter{}

	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, "请输入密码账号")
		return
	}
	if err, msg := LoginImpl(p); err != nil {
		response.RequestError(c, msg)
	} else {
		token, err := common.GenToken(p.UserName, p.Password)
		if err != nil {
			response.RequestError(c, "token生成失败")
			return
		}
		data := common.TokenDto{Token: token}
		response.RequestOk(c, data, msg)
	}
}

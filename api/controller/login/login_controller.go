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
	if err := LoginImpl(p); err != nil {
		response.RequestError(c, "请输入正确的密码账号")
	} else {
		token, err := common.GenToken(p.UserName, p.Password)
		if err != nil {
			response.RequestError(c, "登录失败")
		}

		data := common.TokenDto{Token: token}
		response.RequestOk(c, data, "登录成功")
	}

	//fmt.Printf("%v\n", p.UserName)
	//fmt.Printf("%v\n", err)
	//testString(p)
}

//func testString(s string) {
//	fmt.Printf("inner: %v, %v\n", s, &s)
//}

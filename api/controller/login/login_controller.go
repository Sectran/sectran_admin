package loginController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	response "sectran/api"
)

type loginParameter struct {
	UserName string `json:"userName"` //用户名
	Password string `json:"password"` //密码
}

func PostUserLogin(c *gin.Context) {
	p := loginParameter{}
	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, nil, "请输入密码账号")
		return
	}

	fmt.Printf("%v\n", p.UserName)
	//testString(p)
}

func testString(s string) {
	fmt.Printf("inner: %v, %v\n", s, &s)
}

package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	response "sectran/api"
	"sectran/api/common"
	"strings"
	"time"
)

func InitRouter() {
	e := gin.Default()

	LoginRouter(e)
	e.Use(JwtMiddleware())
	DepartmentRouter(e)
	err := e.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取token
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.RequestError(c, "用户不存在")
			c.Abort() //阻止执行
			return
		}
		//token格式错误
		tokenSlice := strings.SplitN(tokenStr, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			response.RequestError(c, "token格式错误")
			c.Abort()
			return
		}
		//验证token
		fmt.Printf("%#v\n", tokenSlice[1])
		tokenStruck, ok := CheckToken(tokenSlice[1])
		if !ok {
			response.RequestError(c, "无效token")
			c.Abort()
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			response.RequestError(c, "token过期")
			c.Abort()
			return
		}
		c.Set("id", tokenStruck.UserID)
		c.Next()
	}
}

var Hotkey = []byte("blog_jwt_key")

// CheckToken 验证token
func CheckToken(token string) (*common.JWTClaims, bool) {
	tokenObj, _ := jwt.ParseWithClaims(token, &common.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Hotkey, nil
	})
	if key, _ := tokenObj.Claims.(*common.JWTClaims); tokenObj.Valid {
		return key, true
	} else {
		return nil, false
	}
}

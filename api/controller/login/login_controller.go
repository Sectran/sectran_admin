package loginController

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	response "sectran/api"
	"sectran/api/common"
	"time"
)

var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f487")
	//该路由下不校验token
	noVerify = []interface{}{"/login", "/ping"}
	//token有效时间（纳秒）
	effectTime = 2 * time.Hour
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
		response.RequestError(c, "登录失败")
	} else {
		if err != nil {
			response.RequestError(c, "生成token失败")
			return
		} else {
			data := common.TokenDto{Token: common.GenerateToken(&common.UserClaims{
				UserName:       p.UserName,
				Password:       p.Password,
				StandardClaims: jwt.StandardClaims{},
			})}
			response.RequestOk(c, data, "登录成功")
		}
	}

	//fmt.Printf("%v\n", p.UserName)
	//fmt.Printf("%v\n", err)
	//testString(p)
}

//func testString(s string) {
//	fmt.Printf("inner: %v, %v\n", s, &s)
//}

// GetToken 生成token
//func GenerateToken(claims *UserClaims) string {
//	//设置token有效期，也可不设置有效期，采用redis的方式
//	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
//	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
//	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
//	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
//	//生成token
//	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
//	if err != nil {
//		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
//		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
//		panic(err)
//	}
//	return sign
//}

//func GenerateToken(claims *common.JWTClaims) string {
//	//设置token有效期，也可不设置有效期，采用redis的方式
//	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
//	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
//	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
//	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
//	//生成token
//	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
//	if err != nil {
//		return err
//	}
//	return sign
//}

package routers

import (
	"api/global"
	"fmt"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// StoreUser 用户表
type StoreUser struct {
	Id       string `json:"id" gorm:"type:char(36);primary_key"` //用户id
	UserName string `json:"userName"`                            //用户名
	Password string `json:"password"`                            //密码
}

//添加账号
func register(c echo.Context) error {
	u := new(StoreUser)
	err := c.Bind(u)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":  http.StatusOK,
			"msg":   "请输入账号和密码",
			"state": false,
		})
	}
	Db := global.DB
	//判断当前数据表是否存在
	if !Db.HasTable(&StoreUser{}) {
		Db.CreateTable(&StoreUser{})
	}
	//判断是否有此账号
	if err := Db.Where("user_name = ?", u.UserName).First(&u).Error; err != nil {
		//密码加密保存到数据库
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost) //加密处理
		if err != nil {
			fmt.Println(err)
		}
		var b = StoreUser{
			uuid.Must(uuid.NewV4(), nil).String(),
			u.UserName,
			string(hash),
		}
		if err := Db.Create(&b).Error; err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"code":  http.StatusOK,
				"msg":   "账号新建失败",
				"state": false,
			})
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"code":  http.StatusOK,
				"msg":   "账号新建成功",
				"state": true,
			})
		}

	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":  http.StatusOK,
			"msg":   "创建失败,已有此账号",
			"state": false,
		})

	}
}

func Login(e *echo.Echo) {
	Store := e.Group("Login")
	//注册接口
	{
		Store.POST("/register", register)
	}
	//登录接口
	//{
	//	Store.POST("/login", StoreLogin)
	//}
}

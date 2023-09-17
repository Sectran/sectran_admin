package main

import (
	"api/global"
	routers "api/router"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"net/http"
)

// User 定义
type User struct {
	Name  string `json:"name"` // 通过json标签定义struct字段转换成json字段的名字。
	Email string `json:"email"`
}

var totalRequests = 1

//中间件函数
func Count(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//在这里处理拦截请求的逻辑
		//累计访问量
		totalRequests++
		//在响应头中输出访问量
		c.Response().Header().Add("requests", fmt.Sprintf("%d", totalRequests))

		//执行下一个中间件或者执行控制器函数, 然后返回执行结果
		return next(c)
	}
}

func main() {
	//实例化echo对象。
	e := echo.New()
	u := &User{
		Name:  "tizi3651111",
		Email: "tizi@tizi365.com",
	}
	e.Use(Count)
	routers.Login(e)
	e.POST("/hello", func(c echo.Context) error {
		//控制器函数直接返回一个字符串，http响应状态为http.StatusOK，就是200状态。
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code": http.StatusOK,
			"data": u,
		})
	})

	//routers.Login(e)

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8082")
}

func init() {
	var err error
	global.DB, err = gorm.Open("mysql", "root:123456@/db_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	fmt.Println(1321)
	//打印查询的sql语句
	global.DB.LogMode(true)
}

//	sectran_admin
//
//	Description: sectran_admin service
//
//	Schemes: http, https
//	Host: localhost:8081
//	BasePath: /
//	Version: 0.0.1
//	SecurityDefinitions:
//	  Token:
//	    type: apiKey
//	    name: Authorization
//	    in: header
//	Security:
//	    - Token: []
//	Consumes:
//	  - application/json
//
//	Produces:
//	  - application/json
//
// swagger:meta
package main

import (
	"context"
	"flag"
	"fmt"
	"math"

	"sectran_admin/internal/config"
	"sectran_admin/internal/handler"
	"sectran_admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/sectran_admin.yaml", "the config file")

func initDept(ctx *svc.ServiceContext) {
	srvCtx := context.Background()
	c, err := ctx.DB.Department.Query().Count(srvCtx)

	if err != nil {
		fmt.Printf("项目初始化失败:%s\n", err)
		return
	}

	if c == 0 {
		_, err := ctx.DB.Department.Create().
			SetName("山川科技").
			SetArea("北京").
			SetDescription("北京山川科技股份有限公司根部门").
			SetParentDepartmentID(math.MaxInt - 1).
			SetParentDepartments("").
			Save(context.Background())
		if err != nil {
			fmt.Printf("项目初始化失败:%s\n", err)
		}
	}
}

func initRole(ctx *svc.ServiceContext) {
	srvCtx := context.Background()
	c, err := ctx.DB.Role.Query().Count(srvCtx)

	if err != nil {
		fmt.Printf("项目初始化失败:%s\n", err)
		return
	}

	if c == 0 {
		_, err := ctx.DB.Role.Create().
			SetName("开发者管理员").
			SetWeight(0).Save(srvCtx)
		if err != nil {
			fmt.Printf("项目初始化失败:%s\n", err)
		}
	}
}

func initUser(ctx *svc.ServiceContext) {
	srvCtx := context.Background()
	c, err := ctx.DB.User.Query().Count(srvCtx)

	if err != nil {
		fmt.Printf("项目初始化失败:%s\n", err)
		return
	}

	if c == 0 {
		_, err := ctx.DB.User.Create().
			SetAccount("administrator").
			SetName("admin").
			SetDepartmentID(1).
			SetRoleID(1).
			SetPassword("0okm)OKM").
			SetStatus(true).
			Save(srvCtx)
		if err != nil {
			fmt.Printf("项目初始化失败:%s\n", err)
		}
	}
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf, rest.WithCors(c.CROSConf.Address))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	handler.RegisterHandlersCustom(server, ctx)

	initDept(ctx)
	initRole(ctx)
	initUser(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

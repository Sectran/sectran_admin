package svc

import (
	"sectran/apiservice/internal/config"
	"sectran/apiservice/internal/middleware"
	"sectran/apiservice/model/st_dept"
	"sectran/apiservice/model/st_role"
	"sectran/apiservice/model/st_user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	AuthorizeMiddleware *middleware.AuthorizeMiddleware
	StDeptModel         st_dept.StDeptModel
	StRoleModel         st_role.StRoleModel
	StUserModel         st_user.StUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:              c,
		AuthorizeMiddleware: middleware.NewAuthorizeMiddleware(),
		StDeptModel:         st_dept.NewStDeptModel(mysqlConn),
		StRoleModel:         st_role.NewStRoleModel(mysqlConn),
		StUserModel:         st_user.NewStUserModel(mysqlConn),
	}
}

package svc

import (
	"sectran/internal/config"
	"sectran/internal/middleware"
	"sectran/model/st_dept"
	"sectran/model/st_role"
	"sectran/model/st_user"

	"github.com/go-playground/validator"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	AuthorizeMiddleware *middleware.AuthorizeMiddleware
	StDeptModel         st_dept.StDeptModel
	StRoleModel         st_role.StRoleModel
	StUserModel         st_user.StUserModel
	Validator           *validator.Validate
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	validate := validator.New()

	return &ServiceContext{
		Config:              c,
		AuthorizeMiddleware: middleware.NewAuthorizeMiddleware(),
		StDeptModel:         st_dept.NewStDeptModel(mysqlConn),
		StRoleModel:         st_role.NewStRoleModel(mysqlConn),
		StUserModel:         st_user.NewStUserModel(mysqlConn),
		Validator:           validate,
	}
}

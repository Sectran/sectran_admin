package svc

import (
	"sectran_admin/internal/config"
	i18n2 "sectran_admin/internal/i18n"
	"sectran_admin/internal/middleware"

	"sectran_admin/ent"
	_ "sectran_admin/ent/runtime"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/casbin/casbin/v2"
)

type ServiceContext struct {
	Config              config.Config
	Casbin              *casbin.Enforcer
	Authority           rest.Middleware
	AuthorityMiddleware *middleware.AuthorityMiddleware
	DB                  *ent.Client
	Trans               *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.CasbinDatabaseConf.Type, c.CasbinDatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	AuthorityMiddleware := middleware.NewAuthorityMiddleware(cbn, rds, trans)
	return &ServiceContext{
		Config:              c,
		AuthorityMiddleware: AuthorityMiddleware,
		Authority:           AuthorityMiddleware.Handle,
		Trans:               trans,
		DB:                  db,
		Casbin:              cbn,
	}

}

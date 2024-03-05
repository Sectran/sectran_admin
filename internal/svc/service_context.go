package svc

import (
	"sectran_admin/internal/config"
	i18n2 "sectran_admin/internal/i18n"
	"sectran_admin/internal/middleware"
	"sectran_admin/internal/utils/jwt"
	"time"

	"sectran_admin/ent"
	_ "sectran_admin/ent/runtime"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
	DB        *ent.Client
	Trans     *i18n.Translator
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

	if token, err := jwt.GenerateTokenUsingHs256(c.Auth.AccessSecret, time.Hour*time.Duration(c.Auth.AccessExpire)); err == nil {
		logx.Infof("token:%s", token)
	}

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Trans:     trans,
		DB:        db,
	}

}

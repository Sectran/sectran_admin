package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sectran_admin/ent"
	"sectran_admin/internal/config"
	"sectran_admin/internal/types"
	"strconv"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type AuthorityMiddleware struct {
	Cbn   *casbin.Enforcer
	Rds   redis.UniversalClient
	Trans *i18n.Translator
	Conf  *config.Config
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, rds redis.UniversalClient, trans *i18n.Translator, conf *config.Config) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Cbn:   cbn,
		Rds:   rds,
		Trans: trans,
		Conf:  conf,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj := r.URL.Path
		act := r.Method
		token := r.Header.Get("Authorization")
		ctx := context.Background()

		userJson, err := m.Rds.Get(ctx, token).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			logx.Errorw("从redis中查询token value失败", logx.Field("token", token))
			httpx.Error(w, types.ErrRedis)
			return
		}

		if len(userJson) == 0 {
			httpx.Error(w, types.ErrInvalidToken)
			return
		}

		user := &ent.User{}
		err = json.Unmarshal([]byte(userJson), user)
		if err != nil {
			logx.Errorw("从token中解析user json 失败", logx.Field("token", token))
			httpx.Error(w, types.ErrInternalError)
			return
		}

		if d, err := m.Rds.TTL(ctx, token).Result(); err != nil {
			logx.Errorw("查询token失效时间失败", logx.Field("token", token))
			httpx.Error(w, types.ErrInternalError)
			return
		} else {
			if d < time.Minute*10 {
				exp := time.Minute * time.Duration(m.Conf.Auth.AccessExpire-10)
				if _, err = m.Rds.Expire(ctx, token, exp).Result(); err != nil {
					logx.Errorw("无法延长token失效时间", logx.Field("token", token))
					httpx.Error(w, types.ErrInternalError)
					return
				}
			}
		}

		//开发者管理员
		if user.ID != 1 {
			result := batchCheck(m.Cbn, []uint64{user.RoleID}, act, obj)
			if !result {
				logx.Errorw("the role is not permitted to access the API", logx.Field("roleId", user.RoleID),
					logx.Field("path", obj), logx.Field("method", act))
				httpx.Error(w, types.ErrAccountHasNoRights)
				return
			}
		}

		logx.Infow("HTTP/HTTPS Request", logx.Field("UUID", user.ID),
			logx.Field("path", obj), logx.Field("method", act))
		r = r.WithContext(context.WithValue(r.Context(), "request_domain", user))
		next(w, r)
	}
}

func batchCheck(cbn *casbin.Enforcer, roleIds []uint64, act, obj string) bool {
	var checkReq [][]any
	for _, v := range roleIds {
		checkReq = append(checkReq, []any{strconv.FormatUint(v, 10), obj, act})
	}

	result, err := cbn.BatchEnforce(checkReq)
	if err != nil {
		logx.Errorw("Casbin enforce error", logx.Field("detail", err.Error()))
		return false
	}

	for _, v := range result {
		if v {
			return true
		}
	}

	return false
}

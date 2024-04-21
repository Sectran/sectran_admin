package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sectran_admin/ent"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"
)

type AuthorityMiddleware struct {
	Cbn   *casbin.Enforcer
	Rds   redis.UniversalClient
	Trans *i18n.Translator
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, rds redis.UniversalClient, trans *i18n.Translator) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Cbn:   cbn,
		Rds:   rds,
		Trans: trans,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the path
		obj := r.URL.Path
		// get the method
		act := r.Method

		token := r.Header.Get("Authorization")

		userJson, err := m.Rds.Get(context.Background(), token).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			logx.Errorw("redis error in jwt", logx.Field("detail", err.Error()))
			httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
			return
		}

		logx.Infof("user json :%s", userJson)

		user := &ent.User{}
		err = json.Unmarshal([]byte(userJson), user)
		if err != nil {
			logx.Errorw("user json unmarshal failed", logx.Field("detail", err.Error()))
			httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
			return
		}

		result := batchCheck(m.Cbn, []uint64{user.RoleID}, act, obj)
		if !result {
			logx.Errorw("the role is not permitted to access the API", logx.Field("roleId", user.RoleID),
				logx.Field("path", obj), logx.Field("method", act))
			httpx.Error(w, errorx.NewCodeError(errorcode.PermissionDenied, m.Trans.Trans(
				context.WithValue(context.Background(), "lang", r.Header.Get("Accept-Language")),
				"common.permissionDeny")))
			return
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

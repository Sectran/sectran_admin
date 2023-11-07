package middleware

import (
	"net/http"
	"sectran/apiservice/internal/types"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserAuthedInfo struct {
	*types.UserAllInfo
	ExpTime int64
}

type AuthorizeMiddleware struct {
	//token:user
	UserSessionPool map[string]*UserAuthedInfo
}

const auth_route_path string = "/sectran/auth/login"

func NewAuthorizeMiddleware() *AuthorizeMiddleware {
	auth := &AuthorizeMiddleware{
		UserSessionPool: make(map[string]*UserAuthedInfo),
	}

	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 10/* * * * *", func() {
		logx.Infof("calling token verify method")
		now := time.Now().Unix()
		for token, info := range auth.UserSessionPool {
			if now > info.ExpTime {
				delete(auth.UserSessionPool, token)
			}
		}
	})
	go c.Start()

	return auth
}

func (m *AuthorizeMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			msg string
		)

		if r.URL.Path != auth_route_path {
			token := r.Header.Get("Authorization")

			if token == "" {
				msg = "bad request without token in header."
				goto fatal
			}

			value, ok := m.UserSessionPool[token]
			if !ok {
				msg = "you are not login yet."
				logx.Errorf("%s", token)
				goto fatal
			}

			now := time.Now().Unix()
			//User session has timed out
			if value.ExpTime < now {
				msg = "this session is time out."
				goto fatal
			}
			//other verify

			//缓存用户的IP地址，不允许同时多地登陆。
			//校验请求来源、不允许apipot、curl、等等测试工具访问接口
			//增加接口的签名功能，防止中间人攻击

			//update token exp time
			value.ExpTime = time.Now().Unix() + 1800
		}
		next(w, r)

		return
	fatal:
		w.Write([]byte(msg))
	}

}

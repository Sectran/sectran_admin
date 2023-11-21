package user

import (
	"net/http"

	"sectran/internal/logic/user"
	"sectran/internal/svc"
	"sectran/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAllInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserEditLogic(r.Context(), svcCtx)
		resp, err := l.UserEdit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

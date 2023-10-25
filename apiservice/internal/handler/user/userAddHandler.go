package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sectran/apiservice/internal/logic/user"
	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"
)

func UserAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAllInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserAddLogic(r.Context(), svcCtx)
		resp, err := l.UserAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

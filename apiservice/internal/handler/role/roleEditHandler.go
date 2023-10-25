package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sectran/apiservice/internal/logic/role"
	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"
)

func RoleEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAllInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleEditLogic(r.Context(), svcCtx)
		resp, err := l.RoleEdit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

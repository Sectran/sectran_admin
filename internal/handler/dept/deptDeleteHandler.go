package dept

import (
	"net/http"

	"sectran/internal/logic/dept"
	"sectran/internal/svc"
	"sectran/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeptDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeptDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dept.NewDeptDeleteLogic(r.Context(), svcCtx)
		resp, err := l.DeptDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

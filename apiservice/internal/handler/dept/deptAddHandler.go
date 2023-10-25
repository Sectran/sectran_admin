package dept

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sectran/apiservice/internal/logic/dept"
	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"
)

func DeptAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeptAllInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dept.NewDeptAddLogic(r.Context(), svcCtx)
		resp, err := l.DeptAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

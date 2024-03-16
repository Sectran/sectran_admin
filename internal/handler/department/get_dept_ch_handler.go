package department

import (
	"net/http"
	"sectran_admin/internal/logic/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetChDepartmentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChildrenReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := department.NewChildrenDepartmentByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetChildrenDepartmentById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package department

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /department/list department GetDepartmentList
//
// Get department list | 获取Department列表
//
// Get department list | 获取Department列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DepartmentListReq
//
// Responses:
//  200: DepartmentListResp

func GetDepartmentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DepartmentListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := department.NewGetDepartmentListLogic(r.Context(), svcCtx)
		resp, err := l.GetDepartmentList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

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

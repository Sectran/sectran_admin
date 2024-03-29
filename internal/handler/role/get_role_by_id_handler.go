package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/role"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /role role GetRoleById
//
// Get role by ID | 通过ID获取Role
//
// Get role by ID | 通过ID获取Role
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: RoleInfoResp

func GetRoleByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewGetRoleByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

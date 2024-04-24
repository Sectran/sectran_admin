package access_policy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/access_policy"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /access_policy access_policy GetAccessPolicyById
//
// Get access policy by ID | 通过ID获取AccessPolicy
//
// Get access policy by ID | 通过ID获取AccessPolicy
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: AccessPolicyInfoResp

func GetAccessPolicyByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := access_policy.NewGetAccessPolicyByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetAccessPolicyById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

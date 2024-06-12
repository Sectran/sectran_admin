package policyauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/policyauth"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /policy_auth policyauth GetPolicyAuthById
//
// Get policy auth by ID | 通过ID获取PolicyAuth
//
// Get policy auth by ID | 通过ID获取PolicyAuth
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: PolicyAuthInfoResp

func GetPolicyAuthByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := policyauth.NewGetPolicyAuthByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetPolicyAuthById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

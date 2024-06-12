package policyauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/policyauth"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /policy_auth/create policyauth CreatePolicyAuth
//
// Create policy auth information | 创建PolicyAuth
//
// Create policy auth information | 创建PolicyAuth
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PolicyAuthInfo
//
// Responses:
//  200: BaseMsgResp

func CreatePolicyAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PolicyAuthInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := policyauth.NewCreatePolicyAuthLogic(r.Context(), svcCtx)
		resp, err := l.CreatePolicyAuth(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

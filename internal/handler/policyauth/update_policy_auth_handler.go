package policyauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/policyauth"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /policy_auth/update policyauth UpdatePolicyAuth
//
// Update policy auth information | 更新PolicyAuth
//
// Update policy auth information | 更新PolicyAuth
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PolicyAuthInfo
//
// Responses:
//  200: BaseMsgResp

func UpdatePolicyAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PolicyAuthInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := policyauth.NewUpdatePolicyAuthLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePolicyAuth(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package policyauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/policyauth"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /policy_auth/list policyauth GetPolicyAuthList
//
// Get policy auth list | 获取PolicyAuth列表
//
// Get policy auth list | 获取PolicyAuth列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PolicyAuthListReq
//
// Responses:
//  200: PolicyAuthListResp

func GetPolicyAuthListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PolicyAuthListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := policyauth.NewGetPolicyAuthListLogic(r.Context(), svcCtx)
		resp, err := l.GetPolicyAuthList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

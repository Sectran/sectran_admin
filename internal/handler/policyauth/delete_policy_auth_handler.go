package policyauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/policyauth"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /policy_auth/delete policyauth DeletePolicyAuth
//
// Delete policy auth information | 删除PolicyAuth信息
//
// Delete policy auth information | 删除PolicyAuth信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeletePolicyAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := policyauth.NewDeletePolicyAuthLogic(r.Context(), svcCtx)
		resp, err := l.DeletePolicyAuth(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

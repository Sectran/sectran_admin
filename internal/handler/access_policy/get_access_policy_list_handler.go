package access_policy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/access_policy"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /access_policy/list access_policy GetAccessPolicyList
//
// Get access policy list | 获取AccessPolicy列表
//
// Get access policy list | 获取AccessPolicy列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AccessPolicyListReq
//
// Responses:
//  200: AccessPolicyListResp

func GetAccessPolicyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccessPolicyListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := access_policy.NewGetAccessPolicyListLogic(r.Context(), svcCtx)
		resp, err := l.GetAccessPolicyList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package access_policy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/access_policy"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /access_policy/update access_policy UpdateAccessPolicy
//
// Update access policy information | 更新AccessPolicy
//
// Update access policy information | 更新AccessPolicy
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AccessPolicyInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateAccessPolicyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccessPolicyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := access_policy.NewUpdateAccessPolicyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAccessPolicy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

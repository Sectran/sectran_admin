package access_policy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/access_policy"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /access_policy/create access_policy CreateAccessPolicy
//
// Create access policy information | 创建AccessPolicy
//
// Create access policy information | 创建AccessPolicy
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AccessPolicyInfo
//
// Responses:
//  200: BaseMsgResp

func CreateAccessPolicyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccessPolicyInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := access_policy.NewCreateAccessPolicyLogic(r.Context(), svcCtx)
		resp, err := l.CreateAccessPolicy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

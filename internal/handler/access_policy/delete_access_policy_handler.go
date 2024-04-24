package access_policy

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/access_policy"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /access_policy/delete access_policy DeleteAccessPolicy
//
// Delete access policy information | 删除AccessPolicy信息
//
// Delete access policy information | 删除AccessPolicy信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteAccessPolicyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := access_policy.NewDeleteAccessPolicyLogic(r.Context(), svcCtx)
		resp, err := l.DeleteAccessPolicy(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

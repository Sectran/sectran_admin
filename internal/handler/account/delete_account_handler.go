package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/account"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /account/delete account DeleteAccount
//
// Delete account information | 删除Account信息
//
// Delete account information | 删除Account信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := account.NewDeleteAccountLogic(r.Context(), svcCtx)
		resp, err := l.DeleteAccount(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

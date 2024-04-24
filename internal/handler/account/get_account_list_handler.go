package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/account"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /account/list account GetAccountList
//
// Get account list | 获取Account列表
//
// Get account list | 获取Account列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AccountListReq
//
// Responses:
//  200: AccountListResp

func GetAccountListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccountListReqRefer
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := account.NewGetAccountListLogic(r.Context(), svcCtx)
		resp, err := l.GetAccountList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

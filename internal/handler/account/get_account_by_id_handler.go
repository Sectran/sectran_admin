package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/account"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /account account GetAccountById
//
// Get account by ID | 通过ID获取Account
//
// Get account by ID | 通过ID获取Account
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: AccountInfoResp

func GetAccountByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReqRefer
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := account.NewGetAccountByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetAccountById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

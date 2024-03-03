package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/account"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /account/create account CreateAccount
//
// Create account information | 创建Account
//
// Create account information | 创建Account
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AccountInfo
//
// Responses:
//  200: BaseMsgResp

func CreateAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccountInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := account.NewCreateAccountLogic(r.Context(), svcCtx)
		resp, err := l.CreateAccount(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

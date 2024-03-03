package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/user"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /user/create user CreateUser
//
// Create user information | 创建User
//
// Create user information | 创建User
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UserInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

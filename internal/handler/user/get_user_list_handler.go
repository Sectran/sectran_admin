package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/user"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /user/list user GetUserList
//
// Get user list | 获取User列表
//
// Get user list | 获取User列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UserListReq
//
// Responses:
//  200: UserListResp

func GetUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListReqRefer
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

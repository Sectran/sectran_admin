package labletree

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/labletree"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /lable_tree/list labletree GetLableTreeList
//
// Get lable tree list | 获取LableTree列表
//
// Get lable tree list | 获取LableTree列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: LableTreeListReq
//
// Responses:
//  200: LableTreeListResp

func GetLableTreeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LableTreeListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := labletree.NewGetLableTreeListLogic(r.Context(), svcCtx)
		resp, err := l.GetLableTreeList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

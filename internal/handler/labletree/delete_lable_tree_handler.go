package labletree

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/labletree"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /lable_tree/delete labletree DeleteLableTree
//
// Delete lable tree information | 删除LableTree信息
//
// Delete lable tree information | 删除LableTree信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteLableTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := labletree.NewDeleteLableTreeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteLableTree(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package LableTree

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/LableTree"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /lable_tree/update LableTree UpdateLableTree
//
// Update lable tree information | 更新LableTree
//
// Update lable tree information | 更新LableTree
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: LableTreeInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateLableTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LableTreeInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := LableTree.NewUpdateLableTreeLogic(r.Context(), svcCtx)
		resp, err := l.UpdateLableTree(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

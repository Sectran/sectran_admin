package labletree

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/labletree"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /lable_tree labletree GetLableTreeById
//
// Get lable tree by ID | 通过ID获取LableTree
//
// Get lable tree by ID | 通过ID获取LableTree
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: LableTreeInfoResp

func GetLableTreeByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := labletree.NewGetLableTreeByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetLableTreeById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package device

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/device"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /device/delete device DeleteDevice
//
// Delete device information | 删除Device信息
//
// Delete device information | 删除Device信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteDeviceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := device.NewDeleteDeviceLogic(r.Context(), svcCtx)
		resp, err := l.DeleteDevice(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

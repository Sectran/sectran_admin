package device

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/device"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /device/list device GetDeviceList
//
// Get device list | 获取Device列表
//
// Get device list | 获取Device列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DeviceListReq
//
// Responses:
//  200: DeviceListResp

func GetDeviceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := device.NewGetDeviceListLogic(r.Context(), svcCtx)
		resp, err := l.GetDeviceList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package device

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/device"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /device/update device UpdateDevice
//
// Update device information | 更新Device
//
// Update device information | 更新Device
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DeviceInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateDeviceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := device.NewUpdateDeviceLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDevice(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

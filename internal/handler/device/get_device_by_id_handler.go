package device

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"sectran_admin/internal/logic/device"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

// swagger:route post /device device GetDeviceById
//
// Get device by ID | 通过ID获取Device
//
// Get device by ID | 通过ID获取Device
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: DeviceInfoResp

func GetDeviceByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := device.NewGetDeviceByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetDeviceById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

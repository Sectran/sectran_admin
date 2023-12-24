package system

import (
	v1 "github.com/Sectran/sectran_admin/api/v1"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct{}

func (s *UserRouter) InitDeviceRouter(Router *gin.RouterGroup) {
	deviceRouter := Router.Group("device")
	deviceApi := v1.ApiGroupApp.SystemApiGroup.DeviceApi
	{
		deviceRouter.POST("listDevice", deviceApi.ListDevice)
		deviceRouter.POST("addDevice", deviceApi.AddDevice)
		deviceRouter.POST("updateDevice", deviceApi.UpdateDevice)
		deviceRouter.POST("deleteDevice", deviceApi.DeleteDevice)

		deviceRouter.POST("listDeviceAccount", deviceApi.ListDeviceAccount)
		deviceRouter.POST("addDeviceAccount", deviceApi.AddDeviceAccount)
		deviceRouter.POST("updateDeviceAccount", deviceApi.UpdateDeviceAccount)
		deviceRouter.POST("deleteDeviceAccount", deviceApi.DeleteDeviceAccount)
	}
}

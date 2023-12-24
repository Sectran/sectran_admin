package system

import "github.com/Sectran/sectran_admin/service"

type ApiGroup struct {
	UserApi
	DeviceApi
}

var (
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	deviceService = service.ServiceGroupApp.SystemServiceGroup.DeviceService

	// deptService = service.ServiceGroupApp.SystemServiceGroup.DeptService
)

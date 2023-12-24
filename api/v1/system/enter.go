package system

import "github.com/Sectran/sectran_admin/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	// deptService = service.ServiceGroupApp.SystemServiceGroup.DeptService
)

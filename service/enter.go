package service

import (
	"github.com/Sectran/sectran_admin/service/example"
	"github.com/Sectran/sectran_admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

package router

import (
	"github.com/Sectran/sectran_admin/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

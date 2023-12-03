package router

import (
	"github.com/Sectran/sectran_admin/router/example"
	"github.com/Sectran/sectran_admin/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

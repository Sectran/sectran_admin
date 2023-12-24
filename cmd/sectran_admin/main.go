package main

import (
	"go.uber.org/zap"

	"github.com/Sectran/sectran_admin/core"
	"github.com/Sectran/sectran_admin/global"
	"github.com/Sectran/sectran_admin/initialize"
)

func main() {
	// load conf
	global.GVA_VP = core.Viper("./conf/sectran_admin.yaml") // 初始化Viper

	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()

	core.RunWindowsServer()
}

package initialize

import (
	"net/http"

	swaggerFiles "github.com/swaggo/files"

	"github.com/Sectran/sectran_admin/docs"
	"github.com/Sectran/sectran_admin/global"
	"github.com/Sectran/sectran_admin/middleware"
	"github.com/Sectran/sectran_admin/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由

func Routers() *gin.Engine {

	// 设置为发布模式
	if global.GVA_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) //DebugMode ReleaseMode TestMode
	}

	Router := gin.New()

	if global.GVA_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger(), gin.Recovery())
	}

	// InstallPlugin(Router) // 安装插件
	systemRouter := router.RouterGroupApp.System
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	PrivateGroup.Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitDeviceRouter(PrivateGroup)

	}

	global.GVA_LOG.Info("router register success")
	return Router
}

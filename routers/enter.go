package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func Init_router() *gin.Engine {
	// 解决router警告信息
	gin.SetMode(global.Config.System.Env)
	// 默认路由
	router := gin.Default()
	// 路由分组
	apiRouterGroup := router.Group("api")
	appRouterGroup := RouterGroup{apiRouterGroup}
	appRouterGroup.SettingsRouter() // 系统配置路由
	appRouterGroup.ImagesRouter()   // 图片相关路由
	appRouterGroup.AdvertRouter()   // 广告相关路由
	return router
}

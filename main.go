package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	// 初始化配置文件
	core.Init_conf()
	// 初始化log
	global.Log = core.InitLogger()
	// 初始化数据库连接
	global.DB = core.Init_gorm()
	// 初始化路由
	router := routers.Init_router()
	addr := global.Config.System.Addr()
	global.Log.Info(fmt.Sprintf("server listen at %s", addr))
	router.Run(addr)
}

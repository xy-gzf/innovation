package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
	"innovation/middleware"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("system").Use(middleware.OperationRecord())
	{
		SystemRouter.POST("getSystemConfig", api.GetSystemConfig) // 获取配置文件内容
		SystemRouter.POST("setSystemConfig", api.SetSystemConfig) // 设置配置文件内容
		SystemRouter.POST("getServerInfo", api.GetServerInfo)     // 获取服务器信息
		SystemRouter.POST("reloadSystem", api.ReloadSystem)       // 重启服务
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
)

func InitInitRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("init")
	{
		ApiRouter.POST("initdb", api.InitDB)   // 创建Api
		ApiRouter.POST("checkdb", api.CheckDB) // 创建Api
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
	"innovation/middleware"
)

func InitDynamicRouter(Router *gin.RouterGroup) {
	DynamicRouter := Router.Group("dynamic").Use(middleware.OperationRecord())
	{
		DynamicRouter.POST("createDynamic", api.CreateDynamic)             // 新建Dynamic
		DynamicRouter.DELETE("deleteDynamic", api.DeleteDynamic)           // 删除Dynamic
		DynamicRouter.DELETE("deleteDynamicByIds", api.DeleteDynamicByIds) // 批量删除Dynamic
		DynamicRouter.PUT("updateDynamic", api.UpdateDynamic)              // 更新Dynamic
		DynamicRouter.GET("findDynamic", api.FindDynamic)                  // 根据ID获取Dynamic
		DynamicRouter.GET("getDynamicList", api.GetDynamicList)            // 获取Dynamic列表
	}
}

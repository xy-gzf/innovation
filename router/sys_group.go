package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
	"innovation/middleware"
)

func InitGroupRouter(Router *gin.RouterGroup) {
	GroupRouter := Router.Group("group").Use(middleware.OperationRecord())
	{
		GroupRouter.POST("createGroup", api.CreateGroup)             // 新建Group
		GroupRouter.DELETE("deleteGroup", api.DeleteGroup)           // 删除Group
		GroupRouter.DELETE("deleteGroupByIds", api.DeleteGroupByIds) // 批量删除Group
		GroupRouter.PUT("updateGroup", api.UpdateGroup)              // 更新Group
		GroupRouter.GET("findGroup", api.FindGroup)                  // 根据ID获取Group
		GroupRouter.GET("getGroupList", api.GetGroupList)            // 获取Group列表
		GroupRouter.GET("getMyGroups", api.GetMyGroups)              // 获取我的Group列表
	}
}

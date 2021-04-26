package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api/v1"
	"innovation/middleware"
)

func InitGroupRouter(Router *gin.RouterGroup) {
	GroupRouter := Router.Group("group").Use(middleware.OperationRecord())
	{
		GroupRouter.POST("createGroup", v1.CreateGroup)             // 新建Group
		GroupRouter.DELETE("deleteGroup", v1.DeleteGroup)           // 删除Group
		GroupRouter.DELETE("deleteGroupByIds", v1.DeleteGroupByIds) // 批量删除Group
		GroupRouter.PUT("updateGroup", v1.UpdateGroup)              // 更新Group
		GroupRouter.GET("findGroup", v1.FindGroup)                  // 根据ID获取Group
		GroupRouter.GET("getGroupList", v1.GetGroupList)            // 获取Group列表
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
	"innovation/middleware"
)

func InitCustomerRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("customer").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("customer", api.CreateExaCustomer)     // 创建客户
		ApiRouter.PUT("customer", api.UpdateExaCustomer)      // 更新客户
		ApiRouter.DELETE("customer", api.DeleteExaCustomer)   // 删除客户
		ApiRouter.GET("customer", api.GetExaCustomer)         // 获取单一客户信息
		ApiRouter.GET("customerList", api.GetExaCustomerList) // 获取客户列表
	}
}

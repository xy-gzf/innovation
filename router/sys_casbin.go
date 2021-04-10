package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
	"innovation/middleware"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	{
		CasbinRouter.POST("updateCasbin", api.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", api.GetPolicyPathByAuthorityId)
	}
}

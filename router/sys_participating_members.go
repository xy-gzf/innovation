package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
	"innovation/middleware"
)

func InitParticipatingMembersRouter(Router *gin.RouterGroup) {
	MemberRouter := Router.Group("member").Use(middleware.OperationRecord())
	{
		MemberRouter.POST("createMember", api.CreateMember)             // 新建Member
		MemberRouter.DELETE("deleteMember", api.DeleteMember)           // 删除Member
		MemberRouter.DELETE("deleteMemberByIds", api.DeleteMemberByIds) // 批量删除Member
		MemberRouter.PUT("updateMember", api.UpdateMember)              // 更新Member
		MemberRouter.GET("findMember", api.FindMember)                  // 根据ID获取Member
		MemberRouter.GET("getMemberList", api.GetMemberList)            // 获取Member列表
	}
}

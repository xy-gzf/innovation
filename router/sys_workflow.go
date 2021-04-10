package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
	"innovation/middleware"
)

func InitWorkflowProcessRouter(Router *gin.RouterGroup) {
	WorkflowProcessRouter := Router.Group("workflowProcess").Use(middleware.OperationRecord())
	{
		WorkflowProcessRouter.POST("createWorkflowProcess", api.CreateWorkflowProcess)             // 新建WorkflowProcess
		WorkflowProcessRouter.DELETE("deleteWorkflowProcess", api.DeleteWorkflowProcess)           // 删除WorkflowProcess
		WorkflowProcessRouter.DELETE("deleteWorkflowProcessByIds", api.DeleteWorkflowProcessByIds) // 批量删除WorkflowProcess
		WorkflowProcessRouter.PUT("updateWorkflowProcess", api.UpdateWorkflowProcess)              // 更新WorkflowProcess
		WorkflowProcessRouter.GET("findWorkflowProcess", api.FindWorkflowProcess)                  // 根据ID获取WorkflowProcess
		WorkflowProcessRouter.GET("findWorkflowStep", api.FindWorkflowStep)                        // 根据ID获取工作流步骤
		WorkflowProcessRouter.GET("getWorkflowProcessList", api.GetWorkflowProcessList)            // 获取WorkflowProcess列表
		WorkflowProcessRouter.POST("startWorkflow", api.StartWorkflow)                             // 开启工作流
		WorkflowProcessRouter.POST("completeWorkflowMove", api.CompleteWorkflowMove)               // 提交工作流
		WorkflowProcessRouter.GET("getMyStated", api.GetMyStated)                                  // 获取我发起的工作流
		WorkflowProcessRouter.GET("getMyNeed", api.GetMyNeed)                                      // 获取我的待办
		WorkflowProcessRouter.GET("getWorkflowMoveByID", api.GetWorkflowMoveByID)                  // 获取我的待办
	}
}

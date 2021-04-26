package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"innovation/global"
	"innovation/model"
	"innovation/model/request"
	"innovation/model/response"
	"innovation/service"
)

// @Tags Group
// @Summary 创建Group
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Group true "创建Group"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /group/createGroup [post]
func CreateGroup(c *gin.Context) {
	var group model.Group
	_ = c.ShouldBindJSON(&group)
	if err := service.CreateGroup(group); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Group
// @Summary 删除Group
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Group true "删除Group"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /group/deleteGroup [delete]
func DeleteGroup(c *gin.Context) {
	var group model.Group
	_ = c.ShouldBindJSON(&group)
	if err := service.DeleteGroup(group); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Group
// @Summary 批量删除Group
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Group"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /group/deleteGroupByIds [delete]
func DeleteGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteGroupByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Group
// @Summary 更新Group
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Group true "更新Group"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /group/updateGroup [put]
func UpdateGroup(c *gin.Context) {
	var group model.Group
	_ = c.ShouldBindJSON(&group)
	if err := service.UpdateGroup(group); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Group
// @Summary 用id查询Group
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Group true "用id查询Group"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /group/findGroup [get]
func FindGroup(c *gin.Context) {
	var group model.Group
	_ = c.ShouldBindQuery(&group)
	if err, regroup := service.GetGroup(group.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regroup": regroup}, c)
	}
}

// @Tags Group
// @Summary 分页获取Group列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GroupSearch true "分页获取Group列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /group/getGroupList [get]
func GetGroupList(c *gin.Context) {
	var pageInfo request.GroupSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetGroupInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

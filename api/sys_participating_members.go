package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"innovation/global"
	"innovation/model"
	"innovation/model/request"
	"innovation/model/response"
	"innovation/service"
)

// @Tags Member
// @Summary 创建Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Member true "创建Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Member/createMember [post]
func CreateMember(c *gin.Context) {
	var Member model.SysParticipatingMembers
	_ = c.ShouldBindJSON(&Member)
	if err := service.CreateMember(Member); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Member
// @Summary 删除Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Member true "删除Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Member/deleteMember [delete]
func DeleteMember(c *gin.Context) {
	var Member model.SysParticipatingMembers
	_ = c.ShouldBindJSON(&Member)
	if err := service.DeleteMember(Member); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Member
// @Summary 批量删除Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Member/deleteMemberByIds [delete]
func DeleteMemberByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteMemberByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Member
// @Summary 更新Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Member true "更新Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Member/updateMember [put]
func UpdateMember(c *gin.Context) {
	var Member model.SysParticipatingMembers
	_ = c.ShouldBindJSON(&Member)
	if err := service.UpdateMember(Member); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Member
// @Summary 用id查询Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Member true "用id查询Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Member/findMember [get]
func FindMember(c *gin.Context) {
	var Member model.SysParticipatingMembers
	_ = c.ShouldBindQuery(&Member)
	if err, reMember := service.GetMember(Member.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reMember": reMember}, c)
	}
}

// @Tags Member
// @Summary 分页获取Member列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MemberSearch true "分页获取Member列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Member/getMemberList [get]
func GetMemberList(c *gin.Context) {
	var pageInfo request.MemberSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMemberInfoList(pageInfo); err != nil {
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

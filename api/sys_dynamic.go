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

// @Tags Dynamic
// @Summary 创建Dynamic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dynamic true "创建Dynamic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Dynamic/createDynamic [post]
func CreateDynamic(c *gin.Context) {
	var Dynamic model.SysDynamic
	_ = c.ShouldBindJSON(&Dynamic)
	if err := service.CreateDynamic(Dynamic); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Dynamic
// @Summary 删除Dynamic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dynamic true "删除Dynamic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Dynamic/deleteDynamic [delete]
func DeleteDynamic(c *gin.Context) {
	var Dynamic model.SysDynamic
	_ = c.ShouldBindJSON(&Dynamic)
	if err := service.DeleteDynamic(Dynamic); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Dynamic
// @Summary 批量删除Dynamic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Dynamic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Dynamic/deleteDynamicByIds [delete]
func DeleteDynamicByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDynamicByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Dynamic
// @Summary 更新Dynamic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dynamic true "更新Dynamic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Dynamic/updateDynamic [put]
func UpdateDynamic(c *gin.Context) {
	var Dynamic model.SysDynamic
	_ = c.ShouldBindJSON(&Dynamic)
	if err := service.UpdateDynamic(Dynamic); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Dynamic
// @Summary 用id查询Dynamic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Dynamic true "用id查询Dynamic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Dynamic/findDynamic [get]
func FindDynamic(c *gin.Context) {
	var Dynamic model.SysDynamic
	_ = c.ShouldBindQuery(&Dynamic)
	if err, reDynamic := service.GetDynamic(Dynamic.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDynamic": reDynamic}, c)
	}
}

// @Tags Dynamic
// @Summary 分页获取Dynamic列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DynamicSearch true "分页获取Dynamic列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Dynamic/getDynamicList [get]
func GetDynamicList(c *gin.Context) {
	var pageInfo request.DynamicSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDynamicInfoList(pageInfo); err != nil {
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

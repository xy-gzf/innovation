package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"innovation/global"
	"innovation/model"
	"innovation/model/response"
	"innovation/service"
	"innovation/utils"
	"strconv"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/uploadPdf [post]
func UploadPdf(c *gin.Context) {
	var file model.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	if header.Size > 10*1000*1000 {
		global.GVA_LOG.Error("Pdf文件过大 超过10Mb!", zap.Any("err", err))
		response.FailWithMessage("Pdf文件过大 超过10Mb!", c)
		return
	}
	if header.Header.Get("Content-Type") != "application/pdf" {
		global.GVA_LOG.Error("上传文件类型不为Pdf!", zap.Any("err", err))
		response.FailWithMessage("上传文件类型不为Pdf!", c)
		return
	}

	err, file = service.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("文件上传服务器失败!", zap.Any("err", err))
		response.FailWithMessage("文件上传服务器失败", c)
		return
	}

	// 上传的文件放置于组内
	groupId, _ := strconv.Atoi(c.Request.Header.Get("Group-Id"))
	var group model.SysGroup
	group.ID = uint(groupId)
	group.FilePath = file.Url
	if err = service.AddFilePathToGroup(group); err != nil {
		global.GVA_LOG.Error("文件绑定组失败  请重试!", zap.Any("err", err))
		response.FailWithMessage("文件绑定组失败  请重试!", c)
		return
	}

	response.OkWithDetailed(response.ExaFileResponse{File: file}, "上传成功", c)
}

func DownloadPdf(c *gin.Context) {
	filePath := c.Query("filePath")
	ok, err := utils.PathExists(filePath)
	if !ok || err != nil {
		global.GVA_LOG.Error("文件不存在", zap.Any("err", err))
		response.FailWithMessage("文件不存在", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

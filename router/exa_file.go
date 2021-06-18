package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
)

func InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	FileUploadAndDownloadGroup := Router.Group("fileUploadAndDownload")
	{
		FileUploadAndDownloadGroup.POST("/upload", api.UploadFile)                                 // 上传文件
		FileUploadAndDownloadGroup.POST("/getFileList", api.GetFileList)                           // 获取上传文件列表
		FileUploadAndDownloadGroup.POST("/deleteFile", api.DeleteFile)                             // 删除指定文件
		FileUploadAndDownloadGroup.POST("/breakpointContinue", api.BreakpointContinue)             // 断点续传
		FileUploadAndDownloadGroup.GET("/findFile", api.FindFile)                                  // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/breakpointContinueFinish", api.BreakpointContinueFinish) // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/removeChunk", api.RemoveChunk)                           // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/uploadPdf", api.UploadPdf)                               // 上传Pdf
		FileUploadAndDownloadGroup.GET("/downloadPdf", api.DownloadPdf)                            // 查看Pdf
	}
}

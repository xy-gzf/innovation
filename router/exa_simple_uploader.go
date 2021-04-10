package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
)

func InitSimpleUploaderRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("simpleUploader")
	{
		ApiRouter.POST("upload", api.SimpleUploaderUpload) // 上传功能
		ApiRouter.GET("checkFileMd5", api.CheckFileMd5)    // 文件完整度验证
		ApiRouter.GET("mergeFileMd5", api.MergeFileMd5)    // 合并文件
	}
}

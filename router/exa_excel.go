package router

import (
	"github.com/gin-gonic/gin"
	"innovation/api"
)

func InitExcelRouter(Router *gin.RouterGroup) {
	FileUploadAndDownloadGroup := Router.Group("excel")
	{
		FileUploadAndDownloadGroup.POST("/importExcel", api.ImportExcel)          // 导入Excel
		FileUploadAndDownloadGroup.GET("/loadExcel", api.LoadExcel)               // 加载Excel数据
		FileUploadAndDownloadGroup.POST("/exportExcel", api.ExportExcel)          // 导出Excel
		FileUploadAndDownloadGroup.GET("/downloadTemplate", api.DownloadTemplate) // 下载模板文件
	}
}

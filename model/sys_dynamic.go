package model

import (
	"innovation/global"
)

type SysDynamic struct {
	global.GVA_MODEL
	Title   string `json:"title" gorm:"comment:标题"`
	Author  string `json:"author"  gorm:"comment:作者"`
	Content string `json:"content" gorm:"default:无;comment:内容"`
}

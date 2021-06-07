package model

import (
	"innovation/global"
)

type SysGroup struct {
	global.GVA_MODEL
	GroupName       string          `json:"groupName" gorm:"comment:小组名称"`
	Master          string          `json:"master"  gorm:"comment:组长"`
	Mentor          string          `json:"mentor" gorm:"default:无;comment:导师" `
	CompetitionItem CompetitionItem `json:"competitionItem" gorm:"default:0;comment:比赛项目 1:三创, 2:互联网+, 3:大创"`
	Period          string          `json:"period" gorm:"comment:比赛届 年份"`
	FilePath        string          `json:"filePath" gorm:"comment:文件路径"`
}

type CompetitionItem uint32

const (
	Competition_Default      CompetitionItem = 0
	Competition_Innovation   CompetitionItem = 1
	Competition_InternetPlus CompetitionItem = 2
	Competition_StudentCxcy  CompetitionItem = 3
)

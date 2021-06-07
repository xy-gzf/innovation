package model

import (
	"innovation/global"
)

type SysParticipatingMembers struct {
	global.GVA_MODEL
	GroupId string `json:"groupId" gorm:"comment:小组id"`
	UserId  string `json:"userId"  gorm:"comment:用户id"`
}

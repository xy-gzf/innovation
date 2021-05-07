package model

import (
	"github.com/satori/go.uuid"
	"innovation/global"
)

type SysParticipatingMembers struct {
	global.GVA_MODEL
	UUID    uuid.UUID `json:"uuid" gorm:"comment:参赛成员UUID"`
	GroupId string    `json:"group_id" gorm:"comment:小组id"`
	UserId  string    `json:"user_id"  gorm:"comment:用户id"`
}

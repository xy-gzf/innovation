package request

import "innovation/model"

type MemberSearch struct {
	model.SysParticipatingMembers
	PageInfo
}

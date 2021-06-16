package service

import (
	"innovation/global"
	"innovation/model"
	"innovation/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMember
//@description: 创建Member记录
//@param: Member model.Member
//@return: err error

func CreateMember(Member model.SysParticipatingMembers) (err error) {
	err = global.GVA_DB.Create(&Member).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMember
//@description: 删除Member记录
//@param: Member model.Member
//@return: err error

func DeleteMember(Member model.SysParticipatingMembers) (err error) {
	err = global.GVA_DB.Delete(&Member).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMemberByIds
//@description: 批量删除Member记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMemberByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SysParticipatingMembers{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMember
//@description: 更新Member记录
//@param: Member *model.Member
//@return: err error

func UpdateMember(Member model.SysParticipatingMembers) (err error) {
	err = global.GVA_DB.Save(&Member).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMember
//@description: 根据id获取Member记录
//@param: id uint
//@return: err error, Member model.Member

func GetMember(id uint) (err error, Member model.SysParticipatingMembers) {
	err = global.GVA_DB.Where("id = ?", id).First(&Member).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMemberInfoList
//@description: 分页获取Member记录
//@param: info request.MemberSearch
//@return: err error, list interface{}, total int64

func GetMemberInfoList(info request.MemberSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysParticipatingMembers{})
	var Members []model.SysParticipatingMembers
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != "" {
		db = db.Where("`user_id` = ?", info.UserId)
	}
	if info.GroupId != "" {
		db = db.Where("`group_id` = ?", info.GroupId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Members).Error
	return err, Members, total
}

//func GetMembersByUserId(userId string) (err error, Member []model.SysParticipatingMembers) {
//	err = global.GVA_DB.Where("user_id = ?", userId).First(&Member).Error
//	return
//}

func GetMembersByGroupId(groupId string) (err error, Member []model.SysParticipatingMembers) {
	err = global.GVA_DB.Where("group_id = ?", groupId).First(&Member).Error
	return
}

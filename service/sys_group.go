package service

import (
	"innovation/global"
	"innovation/model"
	"innovation/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateGroup
//@description: 创建Group记录
//@param: group model.Group
//@return: err error

func CreateGroup(group model.SysGroup) (err error) {
	err = global.GVA_DB.Create(&group).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteGroup
//@description: 删除Group记录
//@param: group model.Group
//@return: err error

func DeleteGroup(group model.SysGroup) (err error) {
	err = global.GVA_DB.Delete(&group).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteGroupByIds
//@description: 批量删除Group记录
//@param: ids request.IdsReq
//@return: err error

func DeleteGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SysGroup{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateGroup
//@description: 更新Group记录
//@param: group *model.Group
//@return: err error

func UpdateGroup(group model.SysGroup) (err error) {
	err = global.GVA_DB.Save(&group).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetGroup
//@description: 根据id获取Group记录
//@param: id uint
//@return: err error, group model.Group

func GetGroup(id uint) (err error, group model.SysGroup) {
	err = global.GVA_DB.Where("id = ?", id).First(&group).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetGroupInfoList
//@description: 分页获取Group记录
//@param: info request.GroupSearch
//@return: err error, list interface{}, total int64

func GetGroupInfoList(info request.GroupSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysGroup{})
	var groups []model.SysGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.GroupName != "" {
		db = db.Where("`group_name` = ?", info.GroupName)
	}
	if info.Mentor != "" {
		db = db.Where("`mentor` = ?", info.Mentor)
	}
	if info.Master != "" {
		db = db.Where("`master` = ?", info.Master)
	}
	if info.CompetitionItem != 0 {
		db = db.Where("`competition_item` = ?", info.CompetitionItem)
	}
	if info.Period != "" {
		db = db.Where("`period` = ?", info.Period)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&groups).Error
	return err, groups, total
}

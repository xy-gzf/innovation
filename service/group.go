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

func CreateGroup(group model.Group) (err error) {
	err = global.GVA_DB.Create(&group).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteGroup
//@description: 删除Group记录
//@param: group model.Group
//@return: err error

func DeleteGroup(group model.Group) (err error) {
	err = global.GVA_DB.Delete(&group).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteGroupByIds
//@description: 批量删除Group记录
//@param: ids request.IdsReq
//@return: err error

func DeleteGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Group{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateGroup
//@description: 更新Group记录
//@param: group *model.Group
//@return: err error

func UpdateGroup(group model.Group) (err error) {
	err = global.GVA_DB.Save(&group).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetGroup
//@description: 根据id获取Group记录
//@param: id uint
//@return: err error, group model.Group

func GetGroup(id uint) (err error, group model.Group) {
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
	db := global.GVA_DB.Model(&model.Group{})
	var groups []model.Group
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` = ?", info.Name)
	}
	if info.Teacher != "" {
		db = db.Where("`teacher` = ?", info.Teacher)
	}
	if info.Group_leader != "" {
		db = db.Where("`group_leader` = ?", info.Group_leader)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&groups).Error
	return err, groups, total
}

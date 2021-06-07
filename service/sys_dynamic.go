package service

import (
	"innovation/global"
	"innovation/model"
	"innovation/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDynamic
//@description: 创建Dynamic记录
//@param: Dynamic model.Dynamic
//@return: err error

func CreateDynamic(dynamic model.SysDynamic) (err error) {
	err = global.GVA_DB.Create(&dynamic).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDynamic
//@description: 删除Dynamic记录
//@param: Dynamic model.Dynamic
//@return: err error

func DeleteDynamic(Dynamic model.SysDynamic) (err error) {
	err = global.GVA_DB.Delete(&Dynamic).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDynamicByIds
//@description: 批量删除Dynamic记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDynamicByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SysDynamic{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDynamic
//@description: 更新Dynamic记录
//@param: Dynamic *model.Dynamic
//@return: err error

func UpdateDynamic(Dynamic model.SysDynamic) (err error) {
	err = global.GVA_DB.Save(&Dynamic).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDynamic
//@description: 根据id获取Dynamic记录
//@param: id uint
//@return: err error, Dynamic model.Dynamic

func GetDynamic(id uint) (err error, Dynamic model.SysDynamic) {
	err = global.GVA_DB.Where("id = ?", id).First(&Dynamic).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDynamicInfoList
//@description: 分页获取Dynamic记录
//@param: info request.DynamicSearch
//@return: err error, list interface{}, total int64

func GetDynamicInfoList(info request.DynamicSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysDynamic{})
	var Dynamics []model.SysDynamic
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Title != "" {
		db = db.Where("`title` = ?", info.Title)
	}
	if info.Author != "" {
		db = db.Where("`author` = ?", info.Author)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Dynamics).Error
	return err, Dynamics, total
}

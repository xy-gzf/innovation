package service

import (
	"innovation/global"
	"innovation/model"
	"innovation/model/request"
	"innovation/model/response"
	"strconv"
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

func CreateGroupTx(group model.SysGroup, member model.SysParticipatingMembers) (err error) {
	tx := global.GVA_DB.Begin()
	if err = tx.Create(&group).Error; err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Create(&member).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return nil
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

func GetNewGroupId(group model.SysGroup) string {
	tx := global.GVA_DB.Last(&group)
	tx.Commit()
	idNum := int(group.ID) + 1
	idStr := strconv.Itoa(idNum)
	return idStr
}

func GetMyGroups(user model.SysUser, info request.GroupSearch) (err error, list interface{}, total int64) {
	if user.AuthorityId == "788" {
		// 学生
		var memberInfo request.MemberSearch
		memberInfo.PageInfo = info.PageInfo
		memberInfo.UserId = strconv.Itoa(int(user.ID))
		if err, list, total = GetMemberInfoList(memberInfo); err != nil {
			return
		}
		members := list.([]model.SysParticipatingMembers)
		listMyGroup := make([]response.MyGroupsResponse, 0)
		for _, member := range members {
			// 通过参赛列表获取到组的信息
			groupId, _ := strconv.Atoi(member.GroupId)
			_, group := GetGroup(uint(groupId))

			// 查找同组成员信息
			users, errGetUsers := getUsersFromGroup(member.GroupId)
			if errGetUsers != nil {
				err = errGetUsers
				return
			}

			var myGroup = &response.MyGroupsResponse{
				Group: group,
				Users: users,
			}
			listMyGroup = append(listMyGroup, *myGroup)
		}
		list = listMyGroup
		return
	}
	if user.AuthorityId == "832" {
		// 导师
		info.Mentor = user.NickName
		if err, list, total = GetGroupInfoList(info); err != nil {
			return
		}
		groups := list.([]model.SysGroup)
		listMyGroup := make([]response.MyGroupsResponse, 0)
		for _, group := range groups {
			users, errGetUsers := getUsersFromGroup(strconv.Itoa(int(group.ID)))
			if errGetUsers != nil {
				return
			}
			var myGroup = &response.MyGroupsResponse{
				Group: group,
				Users: users,
			}
			listMyGroup = append(listMyGroup, *myGroup)
		}
		list = listMyGroup
		return
	}

	// 管理员
	if err, list, total = GetGroupInfoList(info); err != nil {
		return
	}
	groups := list.([]model.SysGroup)
	listMyGroup := make([]response.MyGroupsResponse, 0)
	for _, group := range groups {
		users, errGetUsers := getUsersFromGroup(strconv.Itoa(int(group.ID)))
		if errGetUsers != nil {
			return
		}
		var myGroup = &response.MyGroupsResponse{
			Group: group,
			Users: users,
		}
		listMyGroup = append(listMyGroup, *myGroup)
	}
	list = listMyGroup
	return
}

func getUsersFromGroup(groupId string) (users []model.SysUser, err error) {
	err, members := GetMembersByGroupId(groupId)
	if err != nil {
		return
	}
	for _, member := range members {
		userId, _ := strconv.Atoi(member.UserId)
		errUser, user := FindUserById(userId)
		if errUser != nil {
			err = errUser
			return
		}
		users = append(users, *user)
	}
	return
}

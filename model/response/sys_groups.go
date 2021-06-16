package response

import "innovation/model"

type MyGroupsResponse struct {
	Group model.SysGroup  `json:"groups"`
	Users []model.SysUser `json:"users"`
}

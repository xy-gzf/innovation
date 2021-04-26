// 自动生成模板Group
package model

import (
	"innovation/global"
)

// 如果含有time.Time 请自行import time包
type Group struct {
	global.GVA_MODEL
	Name         string `json:"name" form:"name" gorm:"column:name;comment:小组名称;type:varchar(191);size:191;"`
	Teacher      string `json:"teacher" form:"teacher" gorm:"column:teacher;comment:导师;type:varchar(191);size:191;"`
	Group_leader string `json:"group_leader" form:"group_leader" gorm:"column:group_leader;comment:组长;type:varchar(191);size:191;"`
}

func (Group) TableName() string {
	return "groups"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type GroupWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Group   `json:"business"`
// }

// func (Group) TableName() string {
// 	return "groups"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["group"] = func() model.GVA_Workflow {
//   return new(model.GroupWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["group"] = func() interface{} {
// 	return new(model.Group)
// }

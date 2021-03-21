package model

import (
	"gin-vue-admin/global"
)

type ResearchPaper struct {
	global.GVA_MODEL
	Title       string `json:"title" gorm:"comment:论文名称"`
}

// type Meta struct {
// 	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`
// 	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"`
// 	Title       string `json:"title" gorm:"comment:菜单名"`
// 	Icon        string `json:"icon" gorm:"comment:菜单图标"`
// 	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`
// }

// type SysBaseMenuParameter struct {
// 	global.GVA_MODEL
// 	SysBaseMenuID uint
// 	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
// 	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`
// 	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`
// }

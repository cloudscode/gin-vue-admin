package model

import (
	"gin-vue-admin/global"
	"time"
)


type ResearchPaper struct {
	global.GVA_MODEL
	Title      string `json:"title" gorm:"comment:论文名称"`
	Kind       int `json:"kind" gorm:"comment:论文类型"`
	PublicationType int `json:"publicationType" gorm:"comment:刊物类别"`
	PublicationName string `json:"publicationName" gorm:"comment:刊物名称"`
	Period string `json:"period" gorm:"comment:期号"`
	VolumeNo string `json:"volumeNo" gorm:"comment:卷号"`
	PublishingDate time.Time `json:"publishingDate" gorm:"comment:发表/出版时间"`
	WordLength       int `json:"wordLength" gorm:"comment:字数"`
	KnowledgeClass string `json:"knowledgeClass" gorm:"comment:学科门类"`
	FirstKnowledge string `json:"firstKnowledge" gorm:"comment:一级学科"`
	Source string `json:"source" gorm:"comment:成果来源"`
	PublishingRange string `json:"publishingRange" gorm:"comment:发表范围"`
	Issn string `json:"issn" gorm:"comment:issn号"`
	Cn string `json:"cn" gorm:"comment:cn号"`
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

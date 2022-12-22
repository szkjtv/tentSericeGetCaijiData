package QQmodel

import "github.com/jinzhu/gorm"

type QQEntertainment struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

package Wymodel

import "github.com/jinzhu/gorm"

// 娱乐
type WEntertainment struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:longtext"` //阅读量
	Itm   string `gorm:"type:text"`     //时间
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 体育
type WSports struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 科技
type WTechnology struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 财经
type WFinance struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 要闻
type WFocusnews struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 热点
type WHotspot struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 时尚
type WFashion struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 热榜
type WHotlist struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 旅游
type WTravel struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 军事
type WMilitary struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 知否
type WKnowsno struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 亲子
type WParenting struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 历史
type WHistory struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 游戏
type WGame struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 数码
type WDigital struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 独家
type Wsole struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 教育
type Weducation struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 圈子
type Wcircle struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	// Itm   string `gorm:"type:text"`     //时间戳
	Cmt string `gorm:"type:int"` //评论数
	// Auth  string `gorm:"type:longtext"` //作者
}

// 汽车
type WCar struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 股票
type Wshares struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 家居
type WFurnishing struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// wnba
type Wnba struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 健康
type Whealthy struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 影视
type WMovies struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 星座
type Wconstellation struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 二次元
type Wpixiv struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 手机
type Wphone struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 女人
type Wwoman struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 艺术
type WArt struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 国际足球
type WInternationalFootball struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 双创
type WMass struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:text"`     //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

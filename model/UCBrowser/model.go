package model

import (
	"log"
	// "os/user"

	"github.com/jinzhu/gorm"
	// 	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Username string
	Account  string
	Password string
	Vip      string
	Redate   string
}

func Dbinit() (db *gorm.DB) { //tests   world  golang
	// db, err = gorm.Open("mysql", "ucuc:7itSjY5T6Pbmz8i3@tcp(182.254.202.251:3306)/ucuc?charset=utf8&parseTime=True&loc=Local")
	// db, err = gorm.Open("mysql", "ucuc:aa1451418@tcp(127.0.0.1:3306)/ucuc?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("mysql", "ucuc:aa1451418@tcp(127.0.0.1:3306)/ucuc?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error() + "数据库连接出错了！！")
	}
	db.AutoMigrate(&User{})
	return db

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 娱乐领域
type Ucyl struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 动漫领域
type Animation struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 汽车
type Car struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 游戏
type Game struct {
	// gorm.Model
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 历史
type Historylist struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 热点
type Hotspot struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 社会
type Ucsh struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 体育
type Sports struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 科技
type Technology struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 财经领域
type Economics struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	// Itm  int    // `gorm:"type:int"`      //时间戳
	Cmt  string `gorm:"type:int"`      //评论数
	Auth string `gorm:"type:longtext"` //作者
}

// 图片
type Img struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 综艺
type Variety struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 科学探索
type ScientificExploration struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 健康
type Healthy struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 教育
type Cducation struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 美容瘦身
type BeautyAndSlimming struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 文玩
type Wenwan struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 美食
type Deliciousfood struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 奇闻
type Anecdote struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 职场
type Workplace struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 热榜
type Hotlist struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 钓鱼
type Gofishing struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 影视
type Movies struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 家居
type Furnishing struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 房产
type Property struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 健身
type Bodybuilding struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 数码
type Digital struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 星座
type Constellation struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 互联网
type Internet struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 电影
type Film struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 搞笑
type Funny struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 足球
type Football struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

// 旅游
type Travel struct {
	gorm.Model
	Title string `gorm:"type:longtext"` //标题
	Link  string `gorm:"type:longtext"` //内容链接
	Cnt   string `gorm:"type:int"`      //阅读量
	Itm   string `gorm:"type:int"`      //时间戳
	Cmt   string `gorm:"type:int"`      //评论数
	Auth  string `gorm:"type:longtext"` //作者
}

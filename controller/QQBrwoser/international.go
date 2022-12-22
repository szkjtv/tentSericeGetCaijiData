package qqbrwoser

import (

	// "example.com/m/v2/model"
	"strconv"
	"time"

	QQmodel "example.com/m/v2/model/QQBrowser"
	model "example.com/m/v2/model/UCBrowser"
	"example.com/m/v2/session"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 国际
func Internationadeile(c *gin.Context) {
	db := model.Dbinit()

	var ucSh []QQmodel.Internationa
	var cmt []QQmodel.Internationa
	var user model.User
	Cmt := c.Query("cmt")
	title := c.Query("title")
	auth := c.Query("auth")
	datequuery := c.Query("date")
	xiaoyu := c.Query("xiaoyu")
	UserID := int(session.GetSessionUserId(c))
	date := time.Now().Format("2006-01-02")

	db.First(&user, UserID).Find(&user) //判断用户名session的id
	if date > user.Vip {                //判断用户是否到期
		c.HTML(200, "QQguoji.html", gin.H{
			"daoqi":        "VIP已到期",
			"username":     user.Username,
			"account":      user.Account,
			"vip":          user.Vip,
			"ordinaryuser": "./static/img/ordinaryuser.png", //普通用户头像
		})
		return
	}
	currentPage, err := strconv.Atoi(c.Query("currentPage")) //第几页
	if err != nil {
		currentPage = 1 //第几页
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize")) //一页显示多少条
	if err != nil {
		pageSize = 15 //默认一页显示多少条
	}

	var total int = 0 //总条数据

	// db.Select([]string{"link", "title", "cmt", "cnt", "auth", "FROM_UNIXTIME(itm,'%Y-%m-%d') as itm "}).Order("id desc").Where("cmt >= ?", Cmt).Where("title LIKE ?", "%"+title+"%").Where("auth LIKE ?", "%"+auth+"%").Model(&ucSh).Where("cmt >= ?", Cmt).Count(&total).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&ucSh)

	// 让时间格式转换成时间戳进行搜索  搜索时间大于的
	timeStr := datequuery
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	startTime := t.Unix()

	timeXiaoyu := xiaoyu //搜索时间小于的
	x, _ := time.ParseInLocation("2006-01-02", timeXiaoyu, time.Local)
	endTime := x.Unix()

	if timeXiaoyu == "" {
		db.Select([]string{"link", "title", "cmt", "cnt", "auth", "FROM_UNIXTIME(itm,'%Y-%m-%d') as itm"}).Order("id desc").Where("cnt >= ?", Cmt).Where("title LIKE ?", "%"+title+"%").Where("auth LIKE ?", "%"+auth+"%").Model(&cmt).Count(&total).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&ucSh)

	} else {
		db.Select([]string{"link", "title", "cmt", "cnt", "auth", "FROM_UNIXTIME(itm,'%Y-%m-%d') as itm"}).Order("id desc").Where("cnt >= ?", Cmt).Where("title LIKE ?", "%"+title+"%").Where("auth LIKE ?", "%"+auth+"%").Where("itm >= ? and itm <= ?  ", startTime, endTime).Model(&cmt).Count(&total).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&ucSh)
	}

	a := total/15 + 1
	if currentPage > a {
		renovate := "/user/auth/internationadeile?cmt=&pageSize=15&currentPage=1"
		c.HTML(200, "QQguoji.html", gin.H{
			"datanull":   "没有数据了,不要再点了！！！",
			"renovate":   renovate,
			"username":   user.Username,
			"account":    user.Account,
			"vip":        user.Vip,
			"vipuserimg": "./static/img/icon_svip.png", //vip用户ico图标
		})
		return
	}

	if currentPage == 0 {
		renovate := "/user/auth/internationadeile?cmt=&pageSize=15&currentPage=1"
		c.HTML(200, "QQguoji.html", gin.H{
			"zuosi":      "不要作死",
			"renovate":   renovate,
			"username":   user.Username,
			"account":    user.Account,
			"vip":        user.Vip,
			"vipuserimg": "./static/img/icon_svip.png", //vip用户ico图标
		})
		return
	}

	if Cmt == "" {
		renovate := "/user/auth/internationadeile?cmt=&pageSize=15&currentPage=1"
		c.HTML(200, "QQguoji.html", gin.H{
			"ucsh":        ucSh,            //获取结果
			"searUcCntt":  cmt,             //获取结果
			"total":       total,           //总条数
			"add":         currentPage + 1, //下一页加1
			"reduce":      currentPage - 1, //上一页减1
			"renovate":    renovate,        //刷新
			"currentPage": currentPage,     //前端显示在第几页
			"pagenumber":  a,               //共多少页
			"inputvalee":  Cmt,             //搜索阅读量框架回显
			"nullccmt":    "",              //让搜索参数为空
			"username":    user.Username,
			"account":     user.Account,
			"vip":         user.Vip,
			"vipuserimg":  "./static/img/icon_svip.png", //vip用户ico图标
			"title":       title,                        //搜索的标题回显
			"auth":        auth,                         //搜索的作者回显
			"startTime":   datequuery,                   //搜索日期大于回显
			"endTime":     xiaoyu,                       //搜索日期小于回显
		})
		return
	} else {
		renovate := "/user/auth/internationadeile?cmt=&pageSize=15&currentPage=1"
		c.HTML(200, "QQguoji.html", gin.H{
			"searUcCntt":  cmt,             //输出结果
			"ucsh":        ucSh,            //输出结果
			"ccmt":        "cmt=",          //获取搜索的参数
			"mmcmt":       Cmt,             //获取搜索的参数
			"inputvalee":  Cmt,             //搜索阅读量框架回显
			"add":         currentPage + 1, //下一页加1
			"reduce":      currentPage - 1, //上一页减1
			"currentPage": currentPage,     //前端显示在第几页
			"pagenumber":  a,               //共多少页
			"seratotal":   total,
			"renovate":    renovate, //刷新
			"username":    user.Username,
			"account":     user.Account,
			"vip":         user.Vip,
			"vipuserimg":  "./static/img/icon_svip.png", //vip用户ico图标
			"title":       title,                        //搜索的标题回显
			"auth":        auth,                         //搜索的作者回显
			"startTime":   datequuery,                   //搜索日期大于回显
			"endTime":     xiaoyu,                       //搜索日期小于回显
		})
	}

	defer db.Close()

}

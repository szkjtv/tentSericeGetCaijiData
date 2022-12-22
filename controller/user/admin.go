package user

import (
	"net/http"
	"strconv"

	model "example.com/m/v2/model/UCBrowser"
	"example.com/m/v2/session"
	"github.com/gin-gonic/gin"
)

//查询和搜索用户信息
func Getadmin(c *gin.Context) {
	db := model.Dbinit()
	var info []model.User
	var admin model.User
	Account := c.Query("account")
	Username := c.Query("username")
	currentPage, err := strconv.Atoi(c.Query("currentPage")) //第几页
	if err != nil {
		currentPage = 1 //第几页
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize")) //一页显示多少条
	if err != nil {
		pageSize = 20 //默认一页显示多少条
	}

	var total int = 0 //总条数据

	db.Where("id >= ?", 0).Model(&info).Where("account LIKE  ?", "%"+Account+"%").Where("username LIKE  ?", "%"+Username+"%").Count(&total).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&info)

	// db.Model(&info).Where("id >= ?", 0).Count(&total)

	a := total/20 + 1
	if currentPage > a {
		renovate := "/adminhuang?"
		c.HTML(200, "admin.html", gin.H{
			"datanull": "没有数据了,不要再点了！！！",
			"renovate": renovate,
		})
		return
	}

	if currentPage == 0 {
		renovate := "/admin?"
		c.HTML(200, "admin.html", gin.H{
			"zuosi":    "不要作死",
			"renovate": renovate,
		})
		return
	}

	renovate := "/adminhuang?"
	c.HTML(200, "admin.html", gin.H{
		"info":         info,
		"admin":        admin.Username,
		"adminaccount": admin.Account,
		"total":        total,
		"add":          currentPage + 1,
		"reduce":       currentPage - 1,
		"renovate":     renovate,
		"currentPage":  currentPage, //前端显示在第几页
		"pagenumber":   a,
		"phone":        Account,  //搜索账号回显在输入框架内
		"username":     Username, //搜索用户名回显在输入框内
	})

	defer db.Close()
}

// 显示用户充值页，修改
func ShowEditor(c *gin.Context) {
	db := model.Dbinit()
	var queryuser model.User
	id := c.Param("id")
	db.Find(&queryuser, id)
	c.HTML(http.StatusOK, "editor.html", queryuser)
	// c.HTML(http.StatusOK, "admin.html", queryuser)
}

// 给用户充值
func AdminEditor(c *gin.Context) {
	db := model.Dbinit()
	id := c.Param("id")
	username := c.PostForm("username")
	account := c.PostForm("account")
	vip := c.PostForm("vip")
	db.Model(&model.User{}).Where(id).Update(&model.User{
		Username: username,
		Account:  account,
		Vip:      vip,
	})
	// c.JSON(200, "充值成功")
	c.Redirect(http.StatusMovedPermanently, "/adminhuang")
	defer db.Close()
}

//删除用户
func Deleteadmin(c *gin.Context) {
	db := model.Dbinit()
	id := c.Param("id")
	var delete model.User
	db.Where(id).Delete(&delete) //调试前端时添加了id
	// c.HTML(200, "admin.html", gin.H{
	// 	"deletesussce": "删除成功",
	// })
	c.Redirect(http.StatusMovedPermanently, "/adminhuang")
	defer db.Close()

}
func VIPnotice(c *gin.Context) {
	c.HTML(200, "vip.html", nil)
}

// 获取VIP公告提示页面
func VipNotice(c *gin.Context) {
	// c.HTML(200, "vip.html", nil)
	db := model.Dbinit()
	var user model.User
	UserID := int(session.GetSessionUserId(c))
	db.First(&user, UserID).Find(&user) //判断用户名session的id
	c.HTML(200, "vip.html", gin.H{
		"daoqi":        "VIP已到期",
		"id":           user.ID,
		"username":     user.Username,
		"account":      user.Account,
		"vip":          user.Vip,
		"ordinaryuser": "./static/img/ordinaryuser.png", //普通用户头像
		"Rushtobuy":    "立即抢购",
		"crushtobuy":   "crushtobuy",
	})

}

// 不需要登录的
func NologinVipNotice(c *gin.Context) {

	c.HTML(200, "vip.html", gin.H{})

}

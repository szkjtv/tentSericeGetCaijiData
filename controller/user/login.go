package user

import (
	"net/http"
	"net/url"

	model "example.com/m/v2/model/UCBrowser"
	"example.com/m/v2/session"
	"github.com/gin-gonic/gin"
	"gopkg.in/cas.v2"
)

func GetpageLogin(c *gin.Context) {
	db := model.Dbinit()
	var user model.User
	UserID := int(session.GetSessionUserId(c))

	db.First(&user, UserID).Find(&user) //判断用户名session的id
	if UserID == UserID {               //判断用户是否到期
		c.Redirect(http.StatusMovedPermanently, "/user/auth/index")
	} else {
		c.HTML(200, "notice.html", gin.H{
			"notice": "您还没有登录,点击登录!",
		})
	}

	c.Redirect(http.StatusMovedPermanently, "/user/auth/index")
	defer db.Close()
}

// 登录
func Login(c *gin.Context) {
	db := model.Dbinit()
	var user model.User
	// username := c.Request.FormValue("username")
	// account := c.Request.FormValue("account")
	// password := c.Request.FormValue("password")
	account := c.PostForm("account")
	password := c.PostForm("password")

	db.Where("account = ?", account).First(&user)
	if user.ID == 0 {
		c.HTML(200, "notice.html", gin.H{
			"accountno": "账号不存在",
			"notice":    "请重新登录",
		})
		// c.JSON(203, "账号不存在,请核对账号")
		return
	}

	// if hasSession := session.HasSession(c); hasSession == true {
	// 	c.String(200, "用户登录过了")
	// 	// c.Redirect(http.StatusMovedPermanently, "/user/auth")
	// 	return
	// }

	//user := mysql.UserDetailByName(account)  //无效只能用这个判断来进行操作
	db.Where("password = ?", password).First(&user)

	if err := session.Compare(user.Password, password); err != nil {
		c.HTML(200, "notice.html", gin.H{
			"passworderro": "密码错误",
			"notice":       "请重新登录",
		})
		return
	}

	session.SaveAuthSession(c, user.ID) //登录保存session信息
	session.SaveAuthSessionUserName(c, user.ID, user.Account)
	// c.JSON(200, gin.H{
	// 	"id:":  user.ID,
	// 	"用户名:": user.Username,
	// 	"账号:":  user.Account,
	// })

	// c.JSON(200, "登录成功")

	u, _ := url.Parse("https://cas.example.com")
	client := cas.NewClient(&cas.Options{URL: u})

	h := client.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		if !cas.IsAuthenticated(r) {
			client.RedirectToLogin(w, r)
		} else {
			c.String(200, "已登录")
		}
	})
	h.ServeHTTP(c.Writer, c.Request)
	// c.String(200, "登录成功")

	c.Redirect(http.StatusMovedPermanently, "/user/auth/index")
	defer db.Close()
}

func Logout(c *gin.Context) {
	if hasSession := session.HasSession(c); hasSession == false {
		c.String(401, "您还没登陆")
		return
	}
	session.ClearAuthSession(c) //清除session信息
	c.HTML(200, "notice.html", gin.H{
		"notice": "您还没有登录,点击登录!",
	})
	// c.JSON(200, "退出 ")
	// c.String(200, "已退出")
	// c.Redirect(http.StatusMovedPermanently, "/")
	// return
	// c.String(200, "已经退出了")
	// session := sessions.Default(c)
	// if session == nil {
	// 	session.Clear()
	// 	c.Redirect(http.StatusMovedPermanently, "/user/auth/car")
	// }
}

// 已实现用验证码登录验证码功能
func Captlogin(c *gin.Context) {

	account := c.PostForm("account")
	password := c.PostForm("password")
	value := c.PostForm("code")
	if CaptchaVerify(c, value) {
		db := model.Dbinit()
		var user model.User
		db.Where("account = ?", account).First(&user)
		if user.ID == 0 {
			c.HTML(200, "notice.html", gin.H{
				"accountno": "账号不存在",
				"notice":    "请重新登录",
			})
			return
		}
		db.Where("password = ?", password).First(&user)
		if err := session.Compare(user.Password, password); err != nil {
			c.HTML(200, "notice.html", gin.H{
				"passworderro": "密码错误",
				"notice":       "请重新登录",
			})
			return
		}
		session.SaveAuthSession(c, user.ID) //登录保存session信息
		session.SaveAuthSessionUserName(c, user.ID, user.Account)
		u, _ := url.Parse("https://cas.example.com")
		client := cas.NewClient(&cas.Options{URL: u})
		h := client.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
			if !cas.IsAuthenticated(r) {
				client.RedirectToLogin(w, r)
			} else {
				// c.String(200, "已登录")
			}
		})
		h.ServeHTTP(c.Writer, c.Request)
		c.Redirect(http.StatusMovedPermanently, "/user/auth/index")
		defer db.Close()

	} else {
		c.HTML(http.StatusOK, "notice.html", gin.H{
			"status":     1,
			"noticecode": "验证码输入错误",
		})
	}
}

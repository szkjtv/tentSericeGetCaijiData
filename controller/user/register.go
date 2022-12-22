package user

import (
	"bytes"
	"math/rand"
	"net/http"
	"time"

	model "example.com/m/v2/model/UCBrowser"
	"example.com/m/v2/session"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// RandomString 生成随机字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqlstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]

	}
	return string(result)

}

func Register(c *gin.Context) {
	db := model.Dbinit()
	var loginuser model.User
	loginuser.Username = c.PostForm("username")
	loginuser.Account = c.PostForm("account")
	loginuser.Password = c.PostForm("password")

	//如果用户名为空,生成随机字符串
	if len(loginuser.Username) == 0 {
		loginuser.Username = RandomString(10)
	}

	if len(loginuser.Account) == 0 {
		c.HTML(200, "notice.html", gin.H{
			"acconull": "账号不能为空",
			"notice":   "请输入账号重新注册",
		})
		return
	}

	if len(loginuser.Account) < 11 {
		c.HTML(200, "notice.html", gin.H{
			"numbererro": "您输入的手机号码有误",
			"notice":     "输入手机号错误请重新注册",
		})
		return
	}
	if len(loginuser.Password) == 0 {
		c.HTML(200, "notice.html", gin.H{
			"passnull": "密码不能为空",
			"notice":   "输入密码重新注册",
		})
		return
	}

	if len(loginuser.Username) > 12 {
		c.HTML(200, "notice.html", gin.H{
			"username": "用户名最大为12位",
			"notice":   "输入用户名错误请重新注册",
		})
		return
	}

	if len(loginuser.Account) >= 12 {
		c.HTML(200, "notice.html", gin.H{
			"phonenumbe": "您输入的不是11位数的手机号",
			"notice":     "输入手机号错误请重新注册",
		})
		return
	}

	if len(loginuser.Password) < 6 {
		c.HTML(200, "notice.html", gin.H{
			"passowrd": "密码不能小于6位",
			"notice":   "密码错误请重新注册",
		})
		return
	}

	db.Where("account = ?", loginuser.Account).First(&loginuser)
	if loginuser.ID != 0 {
		c.HTML(200, "notice.html", gin.H{
			"accoundz": "账号已存在请重新更换一个新账号",
			"notice":   "请重新注册",
		})

		return
	}

	// 给密码加密
	if pwd, err := session.Encrypt(c.Request.FormValue("password")); err == nil {
		loginuser.Password = pwd
	}

	// date := time.Now().Format("2006-01-02")
	currentTime := time.Now()                                    //获取当前时间
	oldTime := currentTime.AddDate(0, 0, 3).Format("2006-01-02") //年 月 日

	newAdduser := model.User{Username: loginuser.Username, Account: loginuser.Account, Password: loginuser.Password, Redate: oldTime, Vip: oldTime}
	db.Create(&newAdduser)
	session.SaveAuthSession(c, loginuser.ID)
	// c.Redirect(http.StatusMovedPermanently, "/login")
	c.HTML(200, "notice.html", gin.H{
		"regiternotice": "注册成功请登录!",
		"notice":        "点击登录",
	})
}

// 配置session
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "Psych"
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge,
		Path:   "/",
	})
	return store
}

// 创建中间件
// 处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}

// 生成图片
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, hight int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer

	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, hight)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))

	return nil
}

// 生成验证码
func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 50
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Recode(c *gin.Context) {
	db := model.Dbinit()
	var loginuser model.User
	Username := c.PostForm("username")
	Account := c.PostForm("account")
	Password := c.PostForm("password")
	confirm := c.PostForm("confirm")
	value := c.PostForm("code")
	if CaptchaVerify(c, value) {
		// c.JSON(200, value)
		//如果用户名为空,生成随机字符串
		if len(Username) == 0 {
			Username = RandomString(10)
		}

		if len(Account) == 0 {
			c.HTML(200, "notice.html", gin.H{
				"acconull": "账号不能为空",
				"notice":   "请输入账号重新注册",
			})
			return
		}

		if len(Account) < 11 {
			c.HTML(200, "notice.html", gin.H{
				"numbererro": "您输入的手机号码有误",
				"notice":     "输入手机号错误请重新注册",
			})
			return
		}
		if len(Password) == 0 {
			c.HTML(200, "notice.html", gin.H{
				"passnull": "密码不能为空",
			})
			return
		}

		if confirm != Password {
			c.HTML(200, "notice.html", gin.H{
				"passnull": "确认密码不一致",
			})
			return
		}
		if len(Username) > 12 {
			c.HTML(200, "notice.html", gin.H{
				"username": "用户名最大为12位",
				"notice":   "输入用户名错误请重新注册",
			})
			return
		}

		if len(Account) >= 12 {
			c.HTML(200, "notice.html", gin.H{
				"phonenumbe": "您输入的不是11位数的手机号",
				"notice":     "输入手机号错误请重新注册",
			})
			return
		}

		if len(Password) < 6 {
			c.HTML(200, "notice.html", gin.H{
				"passowrd": "密码不能小于6位",
				"notice":   "密码错误请重新注册",
			})
			return
		}

		db.Where("account = ?", Account).First(&loginuser)
		if loginuser.ID != 0 {
			c.HTML(200, "notice.html", gin.H{
				"accoundz": "账号已存在请重新更换一个新账号",
				"notice":   "请重新注册",
			})

			return
		}

		// 给密码加密
		if pwd, err := session.Encrypt(c.Request.FormValue("password")); err == nil {
			Password = pwd
		}

		currentTime := time.Now()                                    //获取当前时间
		oldTime := currentTime.AddDate(0, 0, 3).Format("2006-01-02") //年 月 日
		newAdduser := model.User{Username: Username, Account: Account, Password: Password, Redate: oldTime, Vip: oldTime}
		db.Create(&newAdduser)
		session.SaveAuthSession(c, loginuser.ID)
		c.HTML(200, "notice.html", gin.H{
			"regiternotice": "注册成功请登录!",
			"notice":        "点击登录",
		})
	} else {
		c.HTML(http.StatusOK, "notice.html", gin.H{
			"status":   1,
			"renotice": "注册验证码输入错误",
		})
	}

}

func Captchare(c *gin.Context) {
	Captcha(c, 4)
}

package router

import (
	// "example.com/m/v2/UCBrowser/controller"

	"time"

	qqbrwoser "example.com/m/v2/controller/QQBrwoser"
	controller "example.com/m/v2/controller/UCBrowser"
	Wycontroller "example.com/m/v2/controller/Wangyi"
	alipay "example.com/m/v2/controller/alipay"
	"example.com/m/v2/controller/user"
	model "example.com/m/v2/model/UCBrowser"
	"example.com/m/v2/session"
	"github.com/gin-gonic/gin"
)

func Header(c *gin.Context) {
	db := model.Dbinit()
	var user model.User
	UserID := int(session.GetSessionUserId(c))
	date := time.Now().Format("2006-01-02")
	if UserID == UserID { //判断用户名的id
		db.First(&user, UserID).Find(&user)
		// db.First(&name, UserID).Find(&name)
	}
	if date > user.Vip { //判断用户是否到期
		c.HTML(200, "header", gin.H{
			"daoqi":        "VIP到期了请联系客服开通！",
			"tranusernmae": user.Username,
			"tranaccount":  user.Account,
			"promtvip":     user.Vip,
		})
		return
	} else {
		c.HTML(200, "header", gin.H{
			"usernmae": user.Username,
			"account":  user.Account,
			"vip":      user.Vip,
		})
	}

}

func Router() {
	r := gin.Default()
	r.LoadHTMLGlob("view/**/*")
	r.Static("/static", "./static")
	alipay.InitAliConfig()
	alipay.InitAlipayClient()
	r.GET("/vip", user.NologinVipNotice)
	r.GET("/adminhuang", user.Getadmin)
	r.GET("/showeditor/:id", user.ShowEditor)
	r.POST("/adminRecharge/:id", user.AdminEditor)
	r.GET("/deleuser/:id", user.Deleteadmin)
	r.POST("/alipayid/", alipay.Notify)
	r.GET("/ping", alipay.Pong)
	r.GET("/getpay", alipay.GetPay)
	r.POST("/alipayid", alipay.Notify)
	r.Use(user.Session("Psych"))           //这个和验证码有关session
	r.GET("/captcha", user.Captchare)      //这是给注册的时获取的图形验证码，不能和登录的使用同一个接口就可以
	r.POST("/captcha/verify", user.Recode) //注册要输入验证码验证才能，提交注册成功
	tt := r.Group("/", session.EnableCookieSession())
	{
		tt.GET("/", user.GetpageLogin) //获取登录页面
	}

	u := r.Group("/user", session.EnableCookieSession())
	{
		u.Static("/static", "./static")

		// u.POST("/register", user.Register)      //不需要输入验证码的登录接口，可以先不要了
		u.POST("/captcha/verify", user.Recode)  //需要输入验证码的注册接口
		u.POST("/login", user.Captlogin)        //需要输入验证码的登录接口
		u.GET("/captcha/login", user.Captchare) //这是给登录时获取的图片验证码，不能和注册的使用同一个接口就可以 /captcha/login
		u.GET("/", user.GetpageLogin)           //获取登录页面
		u.GET("/vip", user.VipNotice)           //获取登录页面
		u.POST("/alipayid/:id", alipay.Notify)  //用户使用支付宝支付成功回调修改用户信息的，接口
		//登录后才能访问的路由
		ft := u.Group("/auth", session.MiddleWare())
		{
			ft.Static("/static", "./static")
			ft.GET("/vip", user.VipNotice)
			ft.POST("/alipayid", alipay.Notify)
			ft.GET("/getpay", alipay.GetPay)     //获取付款链接
			ft.GET("/getthree", alipay.GetThree) //获取三个月的付款链接
			ft.GET("/halfyear", alipay.HalfYear) //获取半年的付款链接
			ft.GET("/getyear", alipay.GetYear)   //获取一年的付款链接
			// uc大鱼号的
			ft.GET("/index", controller.Index) //uc娱乐
			ft.GET("/", Header)
			ft.GET("/technology", controller.Technology)                       //科技
			ft.GET("/animation", controller.Animation)                         //动漫
			ft.GET("/car", controller.Car)                                     //汽车
			ft.GET("/game", controller.Game)                                   //游戏
			ft.GET("/historylist", controller.Historylist)                     //历史
			ft.GET("/hotspot", controller.Hotspot)                             //热点
			ft.GET("/ucsh", controller.Ucsh)                                   //社会
			ft.GET("/sports", controller.Sports)                               //体育
			ft.GET("/economics", controller.Economics)                         //财经领域查看
			ft.GET("/img", controller.Img)                                     //图片
			ft.GET("/variety", controller.Variety)                             //综艺
			ft.GET("/scientificExploration", controller.ScientificExploration) //科学探索
			ft.GET("/healthy", controller.Healthy)                             //健康
			ft.GET("/cducation", controller.Cducation)                         //教育
			ft.GET("/beautyAndSlimming", controller.BeautyAndSlimming)         //美容瘦身
			ft.GET("/wenwan", controller.Wenwan)                               //文玩
			ft.GET("/deliciousfood", controller.Deliciousfood)                 //美食
			ft.GET("/anecdote", controller.Anecdote)                           //奇闻
			ft.GET("/workplace", controller.Workplace)                         //奇闻
			ft.GET("/hotlist", controller.Hotlist)                             //热榜
			ft.GET("/gofishing", controller.Gofishing)                         //钓鱼
			ft.GET("/movies", controller.Movies)                               //影视
			ft.GET("/furnishing", controller.Furnishing)                       //家居
			ft.GET("/property", controller.Property)                           //房产
			ft.GET("/bodybuilding", controller.Bodybuilding)                   //健身
			ft.GET("/digital", controller.Digital)                             //数码
			ft.GET("/constellation", controller.Constellation)                 //星座
			ft.GET("/internet", controller.Internet)                           //互联网
			ft.GET("/film", controller.Film)                                   //电影
			ft.GET("/funny", controller.Funny)                                 //搞笑
			ft.GET("/football", controller.Football)                           //足球
			ft.GET("/travel", controller.Travel)                               //旅游
			// QQ浏览器
			ft.GET("/qqcardeile", qqbrwoser.QQcardeile)                     //汽车
			ft.GET("/qqEntertainmentdeile", qqbrwoser.QQEntertainmentdeile) // 娱乐
			ft.GET("/qqhousedeile", qqbrwoser.QQhousedeile)                 // 房产
			ft.GET("/internationadeile", qqbrwoser.Internationadeile)       // 国际
			ft.GET("/qqmoviesdeile", qqbrwoser.QQMoviesdeile)               // 影视
			ft.GET("/qqeconomicsdeile", qqbrwoser.QQEconomicsdeile)         // 财经
			ft.GET("/qqmilitarydeile", qqbrwoser.QQmilitarydeile)           // 军事
			ft.GET("/qqnbadeile", qqbrwoser.QQNBAdeile)                     // NBA
			ft.GET("/qqsportsdeile", qqbrwoser.QQSportsdeile)               // 体育
			ft.GET("/qqtechnolygydeile", qqbrwoser.QQtechnolygydeile)       // 科技
			ft.GET("/qqsharesdeile", qqbrwoser.QQsharesdeile)               // 股票
			// 网易
			ft.GET("/wyyl", Wycontroller.WEntertainment)                           // 娱乐
			ft.GET("/wytiyi", Wycontroller.WSports)                                // 体育
			ft.GET("/wtechnology", Wycontroller.WTechnology)                       // 科技
			ft.GET("/wfinance", Wycontroller.WFinance)                             // 财经
			ft.GET("/wfocusnews", Wycontroller.WFocusnews)                         // 要闻
			ft.GET("/whotspot", Wycontroller.WHotspot)                             // 热点
			ft.GET("/wfashion", Wycontroller.WFashion)                             // 时尚
			ft.GET("/whotlist", Wycontroller.WHotlist)                             // 热榜
			ft.GET("/wtravel", Wycontroller.WTravel)                               // 旅游
			ft.GET("/wmilitary", Wycontroller.WMilitary)                           // 军事
			ft.GET("/wknowsno", Wycontroller.WKnowsno)                             // 知否
			ft.GET("/wparenting", Wycontroller.WParenting)                         // 亲子
			ft.GET("/whistory", Wycontroller.WHistory)                             // 历史
			ft.GET("/WGame", Wycontroller.WGame)                                   // 游戏
			ft.GET("/WDigital", Wycontroller.WDigital)                             // 数码
			ft.GET("/Wsole", Wycontroller.Wsole)                                   // 独家
			ft.GET("/Weducation", Wycontroller.Weducation)                         // 教育
			ft.GET("/Wcircle", Wycontroller.Wcircle)                               // 圈子
			ft.GET("/WCar", Wycontroller.WCar)                                     // 汽车
			ft.GET("/Wshares", Wycontroller.Wshares)                               // 股票
			ft.GET("/WFurnishing", Wycontroller.WFurnishing)                       // 家居
			ft.GET("/Wnba", Wycontroller.Wnba)                                     // nba
			ft.GET("/Whealthy", Wycontroller.Whealthy)                             // 健康
			ft.GET("/WMovies", Wycontroller.WMovies)                               // 影视
			ft.GET("/Wconstellation", Wycontroller.Wconstellation)                 // 星座
			ft.GET("/Wpixiv", Wycontroller.Wpixiv)                                 // 二次元
			ft.GET("/Wphone", Wycontroller.Wphone)                                 // 手机
			ft.GET("/Wwoman", Wycontroller.Wwoman)                                 // 女人
			ft.GET("/WArt", Wycontroller.WArt)                                     // 艺术
			ft.GET("/WInternationalFootball", Wycontroller.WInternationalFootball) // 国际足球
			ft.GET("/WMass", Wycontroller.WMass)                                   // 双创

			ft.GET("/loginout", user.Logout) //退出登录

		}
	}

	// 运行端口
	r.Run(":8071")

}

// ////// 跨域
// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method               //请求方法
// 		origin := c.Request.Header.Get("Origin") //请求头部
// 		var headerKeys []string                  // 声明请求头keys
// 		for k, _ := range c.Request.Header {
// 			headerKeys = append(headerKeys, k)
// 		}
// 		headerStr := strings.Join(headerKeys, ", ")
// 		if headerStr != "" {
// 			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
// 		} else {
// 			headerStr = "access-control-allow-origin, access-control-allow-headers"
// 		}
// 		if origin != "" {
// 			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
// 			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
// 			//  header的类型
// 			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
// 			//              允许跨域设置                                                                                                      可以返回其他子段
// 			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
// 			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
// 			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
// 			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
// 		}

// 		//放行所有OPTIONS方法
// 		if method == "OPTIONS" {
// 			c.JSON(http.StatusOK, "Options Request!")
// 		}
// 		// 处理请求
// 		c.Next() //  处理请求
// 	}
// }

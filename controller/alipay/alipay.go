package alipay

import (
	// "/xueyibootsteap/alipay/config"
	// "controller/config"

	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	// "github.com/satori/go.uuid"

	"example.com/m/v2/config"
	model "example.com/m/v2/model/UCBrowser"
	"example.com/m/v2/session"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
)

var (
	AliPayClient *alipay.Client
	AliPayConfig config.AliConfig
)

// func main() {
// 	r := gin.Default()
// 	InitAliConfig()
// 	InitAlipayClient()
// 	r.GET("/ping", Pong)
// 	r.GET("/getpay", GetPay)
// 	r.POST("/alipay/notify", Notify)

// 	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

// InitAliConfig 初始化配置文件
func InitAliConfig() {

	v := viper.New()
	// 路径必须要写相对路径,相对于项目的路径
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %s", err.Error())
	}
	// 映射到结构体
	if err := v.Unmarshal(&AliPayConfig); err != nil {
		log.Fatalf("映射结构体失败: %s", err.Error())
	}
}

// InitAlipayClient 初始化*alipay.Client
// 生成支付url和回调都需要用,只初始化一次就可以了
func InitAlipayClient() {
	var err error
	// isProduction - 是否为生产环境，传 false 的时候为沙箱环境，用于开发测试，正式上线的时候需要改为 true
	AliPayClient, err = alipay.New(AliPayConfig.AppID, AliPayConfig.PrivateKey, true) // 为false时是沙箱环境https://openapi.alipaydev.com/gateway.do
	if err != nil {
		log.Fatalf("加载alipay.Client失败: %s", err.Error())
	}
	// 加载支付宝公钥
	err = AliPayClient.LoadAliPayPublicKey(AliPayConfig.AliPublicKey)
	if err != nil {
		log.Fatalf("加载支付宝公钥失败: %s", err.Error())
	}
}

// Pong ReturnURL:支付成功后会请求一次这个
func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// GetPay 获取支付url 充值一个月的接口一个月的
func GetPay(ctx *gin.Context) {
	OutTradeNo := fmt.Sprintf("%s", strings.ReplaceAll(uuid.NewV4().String(), "-", "")) //生成订单号
	var p = alipay.TradePagePay{}
	UserID := session.GetSessionUserId(ctx)
	stringA := strconv.Itoa(int(UserID))           //将uint转换成文本类型给回调接口拼接
	p.NotifyURL = AliPayConfig.NotifyURL + stringA // 回调地址;用来通知我们支付结果的,好去修改状态
	p.ReturnURL = AliPayConfig.ReturnURL           // 返回地址;支付成功后,浏览器内跳转地址
	p.Subject = AliPayConfig.Subject + "标准版（月付）（15.00/1月-1端）"
	p.OutTradeNo = OutTradeNo                //生成订单号
	p.TotalAmount = "15.00"                  //支付金额
	p.ProductCode = AliPayConfig.ProductCode //产品标题支付的时候显示

	url, err := AliPayClient.TradePagePay(p)
	if err != nil {
		log.Fatalf("生成支付url失败: %s", err.Error())
	}
	// fmt.Println("拼接id", p.NotifyURL)
	// fmt.Println("用户id", UserID)
	var payURL = url.String()
	// log.Printf("订单号:%s创建支付url成功;", OutTradeNo)
	// log.Printf(payURL)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"url":  payURL,
	// 	"pay":  "付款",
	// 	"用户id": UserID,
	// })
	db := model.Dbinit()
	var user model.User
	db.First(&user, UserID)
	timex := time.Now().Format("2006-01-02")
	if user.Vip < timex {
		tm2, _ := time.Parse("2006-01-02", timex)
		oldTime := tm2.AddDate(0, 1, 0).Format("2006-01-02")
		ctx.HTML(200, "noticepay.html", gin.H{
			"url":           payURL,
			"pay":           "提交订单",
			"Paymentamount": p.TotalAmount,
			"After":         15,
			"day":           0.5,
			"vip":           oldTime,
			"payment":       "标准版（1月）",
		})
		// return
	} else {
		tm2, _ := time.Parse("2006-01-02", user.Vip)
		oldTime := tm2.AddDate(0, 1, 0).Format("2006-01-02")
		ctx.HTML(200, "noticepay.html", gin.H{
			"url":           payURL,
			"pay":           "提交订单",
			"Paymentamount": p.TotalAmount,
			"After":         15,
			"day":           0.5,
			"vip":           oldTime,
			"payment":       "标准版（1月）",
		})
	}
	defer db.Close()
}

// GetPay 获取支付url 充值3个月的接口
func GetThree(ctx *gin.Context) {
	OutTradeNo := fmt.Sprintf("%s", strings.ReplaceAll(uuid.NewV4().String(), "-", "")) //生成订单号
	var p = alipay.TradePagePay{}
	UserID := session.GetSessionUserId(ctx)

	stringA := strconv.Itoa(int(UserID))           //将uint转换成文本类型给回调接口拼接
	p.NotifyURL = AliPayConfig.NotifyURL + stringA // 回调地址;用来通知我们支付结果的,好去修改状态
	p.ReturnURL = AliPayConfig.ReturnURL           // 返回地址;支付成功后,浏览器内跳转地址
	p.Subject = AliPayConfig.Subject + "标准版（季度）（45.00/3月-1端）"
	p.OutTradeNo = OutTradeNo                //生成订单号
	p.TotalAmount = "45.00"                  //支付金额
	p.ProductCode = AliPayConfig.ProductCode //产品标题支付的时候显示

	url, err := AliPayClient.TradePagePay(p)
	if err != nil {
		log.Fatalf("生成支付url失败: %s", err.Error())
	}
	var payURL = url.String()
	db := model.Dbinit()
	var user model.User
	db.First(&user, UserID)
	timex := time.Now().Format("2006-01-02")
	if user.Vip < timex {
		tm2, _ := time.Parse("2006-01-02", timex)
		oldTime := tm2.AddDate(0, 3, 0).Format("2006-01-02")
		ctx.HTML(200, "noticepay.html", gin.H{
			"url":           payURL,
			"pay":           "提交订单",
			"Paymentamount": p.TotalAmount,
			"After":         15,
			"day":           0.5,
			"vip":           oldTime,
			"payment":       "标准版（3月）",
		})
		// return
	} else {
		tm2, _ := time.Parse("2006-01-02", user.Vip)
		oldTime := tm2.AddDate(0, 3, 0).Format("2006-01-02")
		ctx.HTML(200, "noticepay.html", gin.H{
			"url":           payURL,
			"pay":           "提交订单",
			"Paymentamount": p.TotalAmount,
			"After":         15,
			"day":           0.5,
			"vip":           oldTime,
			"payment":       "标准版（3月）",
		})
	}
	defer db.Close()
}

// GetPay 获取支付url 充值半年的接口
func HalfYear(ctx *gin.Context) {
	OutTradeNo := fmt.Sprintf("%s", strings.ReplaceAll(uuid.NewV4().String(), "-", "")) //生成订单号
	var p = alipay.TradePagePay{}
	UserID := session.GetSessionUserId(ctx)
	stringA := strconv.Itoa(int(UserID))           //将uint转换成文本类型给回调接口拼接
	p.NotifyURL = AliPayConfig.NotifyURL + stringA // 回调地址;用来通知我们支付结果的,好去修改状态
	p.ReturnURL = AliPayConfig.ReturnURL           // 返回地址;支付成功后,浏览器内跳转地址
	p.Subject = AliPayConfig.Subject + "标准版（半年）（90.00/6月-1端）"
	p.OutTradeNo = OutTradeNo                //生成订单号
	p.TotalAmount = "90.00"                  //支付金额
	p.ProductCode = AliPayConfig.ProductCode //产品标题支付的时候显示

	url, err := AliPayClient.TradePagePay(p)
	if err != nil {
		log.Fatalf("生成支付url失败: %s", err.Error())
	}
	var payURL = url.String()
	db := model.Dbinit()
	var user model.User
	db.First(&user, UserID)
	timex := time.Now().Format("2006-01-02")
	if user.Vip < timex {
		tm2, _ := time.Parse("2006-01-02", timex)
		oldTime := tm2.AddDate(0, 6, 0).Format("2006-01-02")
		ctx.HTML(200, "noticepay.html", gin.H{
			"url":           payURL,
			"pay":           "提交订单",
			"Paymentamount": p.TotalAmount,
			"After":         15,
			"day":           0.5,
			"vip":           oldTime,
			"payment":       "标准版（半年）",
		})
		// return
	} else {
		tm2, _ := time.Parse("2006-01-02", user.Vip)
		oldTime := tm2.AddDate(0, 6, 0).Format("2006-01-02")
		ctx.HTML(200, "noticepay.html", gin.H{
			"url":           payURL,
			"pay":           "提交订单",
			"Paymentamount": p.TotalAmount,
			"After":         15,
			"day":           0.5,
			"vip":           oldTime,
			"payment":       "标准版（半年）",
		})
	}
	defer db.Close()
}

// GetPay 获取支付url 充值一年年的接口
func GetYear(ctx *gin.Context) {
	OutTradeNo := fmt.Sprintf("%s", strings.ReplaceAll(uuid.NewV4().String(), "-", "")) //生成订单号
	var p = alipay.TradePagePay{}
	UserID := session.GetSessionUserId(ctx)

	stringA := strconv.Itoa(int(UserID))           //将uint转换成文本类型给回调接口拼接
	p.NotifyURL = AliPayConfig.NotifyURL + stringA // 回调地址;用来通知我们支付结果的,好去修改状态
	p.ReturnURL = AliPayConfig.ReturnURL           // 返回地址;支付成功后,浏览器内跳转地址
	p.Subject = AliPayConfig.Subject + "标准版（年付）（180.00/1年-1端）"
	p.OutTradeNo = OutTradeNo                //生成订单号
	p.TotalAmount = "180.00"                 //支付金额
	p.ProductCode = AliPayConfig.ProductCode //产品标题支付的时候显示

	url, err := AliPayClient.TradePagePay(p)
	if err != nil {
		log.Fatalf("生成支付url失败: %s", err.Error())
	}
	var payURL = url.String()
	db := model.Dbinit()
	var user model.User
	db.First(&user, UserID)

	timex := time.Now().Format("2006-01-02")
	if user.Vip < timex {
		tm2, _ := time.Parse("2006-01-02", timex)
		oldTime := tm2.AddDate(1, 0, 0).Format("2006-01-02")
		ctx.HTML(200, "noticepay.html", gin.H{
			"url":           payURL,
			"pay":           "提交订单",
			"Paymentamount": p.TotalAmount,
			"After":         15,
			"day":           0.5,
			"vip":           oldTime,
			"payment":       "标准版（一年）",
		})
		// return
	} else {
		tm2, _ := time.Parse("2006-01-02", user.Vip)
		oldTime := tm2.AddDate(1, 0, 0).Format("2006-01-02")
		ctx.HTML(200, "noticepay.html", gin.H{
			"url":           payURL,
			"pay":           "提交订单",
			"Paymentamount": p.TotalAmount,
			"After":         15,
			"day":           0.5,
			"vip":           oldTime,
			"payment":       "标准版（一年）",
		})

	}

	defer db.Close()
}

// Notify 回调地址: 支付成功后会回调这里;我们可以用来修改订单状态等等
func Notify(c *gin.Context) {
	noti, err := AliPayClient.GetTradeNotification(c.Request)
	// c.String(http.StatusOK, "%s", "success")
	if err != nil {
		log.Println("获取回调信息失败", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Printf("订单号:%s;状态:%s\n", noti.OutTradeNo, noti.TradeStatus)
	c.String(http.StatusOK, "success")
	db := model.Dbinit()
	var user model.User
	id := c.Param("id")
	db.First(&user, id)

	timex := time.Now().Format("2006-01-02")
	// 这个是当用户一直不充值，再次充一个月的回调
	if user.Vip < timex {
		if noti.TotalAmount == "15.00" {
			tmx, _ := time.Parse("2006-01-02", timex)
			oldTime := tmx.AddDate(0, 1, 0).Format("2006-01-02")
			db.Model(&model.User{}).Where(id).Update(&model.User{
				Vip: oldTime,
			})
			fmt.Println("当前时间", timex)
			fmt.Println("过期充一个月支付状态", noti.TradeStatus)
			return
		}
	}
	// 这个是当用户一直不充值，再次充3月的回调
	if user.Vip < timex {
		if noti.TotalAmount == "45.00" {
			tmx, _ := time.Parse("2006-01-02", timex)
			oldTime := tmx.AddDate(0, 3, 0).Format("2006-01-02")
			db.Model(&model.User{}).Where(id).Update(&model.User{
				Vip: oldTime,
			})
			fmt.Println("当前时间", timex)
			fmt.Println("过期充3个月支付状态", noti.TradeStatus)
			return
		}
	}

	// 这个是当用户一直不充值，再次充6个月的回调
	if user.Vip < timex {
		if noti.TotalAmount == "90.00" {
			tmx, _ := time.Parse("2006-01-02", timex)
			oldTime := tmx.AddDate(0, 6, 0).Format("2006-01-02")
			db.Model(&model.User{}).Where(id).Update(&model.User{
				Vip: oldTime,
			})
			fmt.Println("当前时间", timex)
			fmt.Println("过期充半年支付状态", noti.TradeStatus)
			return
		}
	}

	// 这个是当用户一直不充值，再次充一年的回调
	if user.Vip < timex {
		if noti.TotalAmount == "180.00" {
			tmx, _ := time.Parse("2006-01-02", timex)
			oldTime := tmx.AddDate(1, 0, 0).Format("2006-01-02")
			db.Model(&model.User{}).Where(id).Update(&model.User{
				Vip: oldTime,
			})
			fmt.Println("当前时间", timex)
			fmt.Println("过期充1年支付状态", noti.TradeStatus)
			return
		}

	}

	// 没过期再充1月的
	if noti.TotalAmount == "15.00" {
		tm2, _ := time.Parse("2006-01-02", user.Vip)
		oldTime := tm2.AddDate(0, 1, 0).Format("2006-01-02")
		db.Model(&model.User{}).Where(id).Update(&model.User{
			Vip: oldTime,
		})
		return
	}

	// 没过期再充3月的
	if noti.TotalAmount == "45.00" {
		tm2, _ := time.Parse("2006-01-02", user.Vip)
		oldTime := tm2.AddDate(0, 3, 0).Format("2006-01-02")
		db.Model(&model.User{}).Where(id).Update(&model.User{
			Vip: oldTime,
		})
		fmt.Println("正常三个月支付状态", noti.TradeStatus)
		return

	}
	// 没过期再充6月的
	if noti.TotalAmount == "90.00" {
		tm2, _ := time.Parse("2006-01-02", user.Vip)
		oldTime := tm2.AddDate(0, 6, 0).Format("2006-01-02")
		db.Model(&model.User{}).Where(id).Update(&model.User{
			Vip: oldTime,
		})
		fmt.Println("正常半年支付状态", noti.TradeStatus)
		return
	}
	// 没过期再充一年的
	if noti.TotalAmount == "180.00" {
		tm2, _ := time.Parse("2006-01-02", user.Vip)
		oldTime := tm2.AddDate(1, 0, 0).Format("2006-01-02")
		db.Model(&model.User{}).Where(id).Update(&model.User{
			Vip: oldTime,
		})
		fmt.Println("正常一年支付状态", noti.TradeStatus)
		return
	}

	defer db.Close()

}

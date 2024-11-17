package main

import (
	"errors"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
	"os"
)

var (
	MchId                             = "1693369919"
	SerialNo                          = "432B01701EB1DC4070DB35C97557FC69AE3C2342"
	APIv3Key                          = "SiChuanFuDianYunSuanKeJi20230605 "
	PrivateKey                        = ""
	Success                           = 200
	mchID                      string = "1693369919"                               // 商户号
	mchCertificateSerialNumber string = "432B01701EB1DC4070DB35C97557FC69AE3C2342" // 商户证书序列号
	mchAPIv3Key                string = "SiChuanFuDianYunSuanKeJi20230605"         // 商户APIv3密钥
	//mchID                      string = "1693369919"                               // 商户号
	//mchCertificateSerialNumber string = "432B01701EB1DC4070DB35C97557FC69AE3C2342" // 商户证书序列号
	//mchAPIv3Key                string = "SiChuanFuDianYunSuanKeJi20230605"         // 商户APIv3密钥
	//https://blog.csdn.net/One_Rabbit2016/article/details/131574198?ops_request_misc=&request_id=&biz_id=102&utm_term=go%E5%AE%9E%E7%8E%B0%E5%BE%AE%E4%BF%A1%E5%B0%8F%E7%A8%8B%E5%BA%8F%E6%94%AF%E4%BB%98&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduweb~default-0-131574198.142^v100^pc_search_result_base1&spm=1018.2226.3001.4187
)

var (
	NotifyUrl = "https://www.cdszly.cn/slq-api/applet/common/payNotify"
)

var client *wechat.ClientV3

func main() {

	fmt.Println(fmt.Sprintf("-----init wxpay ------- "))

	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容

	err := errors.New("")

	client, err = wechat.NewClientV3(MchId, SerialNo, APIv3Key, ReadPem())
	if err != nil {
		xlog.Error(err)
		return
	}

	// 设置微信平台API证书和序列号（推荐开启自动验签，无需手动设置证书公钥等信息）
	//client.SetPlatformCert([]byte(""), "")

	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = client.AutoVerifySign()
	if err != nil {
		xlog.Error(err)
		return
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn
}

func ReadPem() string {
	privateKey, err := os.ReadFile("D:\\code\\sl2\\1693369919_20241107_cert\\apiclient_key.pem")
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(privateKey)
}

//
//func (b *BaseApi) CreateOrderPay(c *gin.Context) {
//	//// todo redis 分布式锁操作
//	//var order adminReq.OrderToPay
//	//err := c.ShouldBindJSON(&order)
//	//if err != nil {
//	//	response.FailWithMessage(err.Error(), c)
//	//	return
//	//}
//	//// 订单
//	//exitOrder, err := orderService.GetOrder(order.OrderId)
//	//if err != nil {
//	//	response.FailWithMessage(err.Error(), c)
//	//	return
//	//}
//	//if exitOrder.PayStatus == 1 {
//	//	response.FailWithMessage("订单已支付", c)
//	//	return
//	//}
//	//
//	//// 我方平台订单号
//	//tradeNo := exitOrder.Code
//	//// 过期时间 15分钟 rfc3339
//	//expire := time.Now().Add(15 * time.Minute).Format(time.RFC3339)
//	//// 订单总金额，单位为分
//	//total := int(*exitOrder.PaidAmount * 100)
//
//	// 初始化 BodyMap
//	bm := make(gopay.BodyMap)
//	bm.Set("sp_appid", AppId).
//		Set("sp_mchid", MchId).
//		Set("sub_mchid", MchId).
//		// 商品描述
//		Set("description", "Jsapi支付商品").
//		// 商户订单号
//		Set("out_trade_no", tradeNo).
//		// time_expire
//		Set("time_expire", expire).
//		// 通知地址
//		Set("notify_url", NotifyUrl).
//		SetBodyMap("amount", func(bm gopay.BodyMap) {
//			// 订单总金额，单位为分
//			bm.Set("total", total).
//				// 货币类型
//				Set("currency", "CNY")
//		}).
//		SetBodyMap("payer", func(bm gopay.BodyMap) {
//			// 用户在直连商户appid下的唯一标识。 下单前需获取到用户的Openid
//			bm.Set("sp_openid", order.OpenId)
//		})
//
//	// 向微信支付平台生成支付回调
//	wxRsp, err := client.V3TransactionJsapi(c, bm)
//	if err != nil {
//		xlog.Error(err)
//		return
//	}
//	if wxRsp.Code == Success {
//		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
//		return
//	}
//	xlog.Errorf("wxRsp:%s", wxRsp.Error)
//
//	// 下单后，获取微信小程序支付、APP支付、JSAPI支付所需要的 pay sign
//	applet, err := client.PaySignOfApplet(AppId, wxRsp.Response.PrepayId)
//
//	// 这里简化处理，直接返回
//	response.OkWithData(applet, c)
//}
//
//// PayNotify
//// @Tags      WeChat
//// @Summary   小程序-微信小程序支付回调  https://github.com/go-pay/gopay/blob/main/doc/wechat_v3.md
//// 微信官方文档 https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_5.shtml
//// @Security  ApiKeyAuth
//// @accept    application/json
//// @Produce   application/json
//// @Success   200   {object}  response.Response{msg=string}  "小程序-微信小程序支付回调"
//// @Router    /applet/common/payNotify [post]
//func (b *BaseApi) PayNotify(c *gin.Context) {
//	//WxPayNotify 解析微信回调请求的参数到 V3NotifyReq 结构体
//	fmt.Println("--------------------- WxPayNotify START ---------------------")
//	// c.Request 是 gin 框架的写法
//	notifyReq, err := wechat.V3ParseNotify(c.Request)
//	if err != nil {
//		fmt.Println("------ WxPayNotify V3ParseNotify ERR ------", err.Error())
//		c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: "回调内容异常"})
//		return
//	}
//
//	// 获取微信平台证书
//
//	// 验证异步通知的签名
//	err = notifyReq.VerifySignByPK(client.WxPublicKey())
//	if err != nil {
//		fmt.Println("------ WxPayNotify VerifySignByPKMap ERR ------", err.Error())
//		c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: "内容验证失败"})
//		return
//	}
//
//	// 通用通知解密（推荐此方法）
//	//result, err := notifyReq.DecryptCipherTextToStruct(APIv3Key, objPtr)
//	// 普通支付通知解密
//	result, rErr := notifyReq.DecryptPayCipherText(APIv3Key)
//
//	if rErr != nil {
//		fmt.Println("------ WxPayNotify DecryptCipherText Error ------", rErr.Error())
//		c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: "内容解密失败"})
//		return
//	}
//	if result != nil {
//		// 查询已存在订单
//		exitOrder, err := orderService.GetOrderByCode(result.OutTradeNo)
//		if err != nil {
//			c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()})
//			return
//		}
//		// 处理支付成功
//		if result.TradeState == "SUCCESS" {
//			fmt.Println("------ WxPayNotify PushMessToPayQueue START 【" + result.OutTradeNo + "】------")
//			var wxReq = make(map[string]interface{})
//			promotionAmount := 0
//			for i := range result.PromotionDetail {
//				promotionAmount += result.PromotionDetail[i].Amount
//			}
//			fmt.Println(fmt.Sprintf("------ WxPayNotify 优惠券总额 promotionAmount 【%d】------", promotionAmount))
//			//商户订单号:商户系统内部订单号
//			wxReq["pay_no"] = result.OutTradeNo
//			//微信支付订单号:微信支付系统生成的订单号。
//			wxReq["trade_no"] = result.TransactionId
//			//与支付宝同步
//			wxReq["trade_status"] = "TRADE_SUCCESS"
//			wxReq["notify_time"] = result.SuccessTime
//			wxReq["total_amount"] = result.Amount.Total
//			wxReq["receipt_amount"] = result.Amount.PayerTotal + promotionAmount
//			var mapData = make(map[string]interface{})
//			mapData["data_type"] = "PayNotify"
//			mapData["param"] = map[string]interface{}{"payType": "wxPay", "notifyReq": wxReq}
//			reqJson, _ := json.Marshal(mapData)
//
//			fmt.Println(string(reqJson))
//
//			// 订单存在
//			if *exitOrder.Id > 0 && exitOrder.PayStatus == 0 {
//				exitOrder.PayStatus = 1
//				exitOrder.PayNotify = string(reqJson)
//				exitOrder.PayNotifyNumber = result.TransactionId
//				// 更新
//				orderService.UpdateOrder(exitOrder)
//			} else if err != nil {
//				c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: "内容解密失败"})
//				return
//			}
//		} else {
//			// 处理支付失败
//			exitOrder.PayStatus = -1
//			reqJson, _ := json.Marshal(result)
//			exitOrder.PayNotify = string(reqJson)
//			// 更新
//			orderService.UpdateOrder(exitOrder)
//		}
//
//		//lib.PushMessToPayQueue(reqJson)
//		fmt.Println("------ WxPayNotify PushMessToPayQueue END 【" + result.OutTradeNo + "】------")
//	}
//
//	/*var wxReq = make(map[string]interface{})
//	//商户订单号:商户系统内部订单号
//	wxReq["pay_no"] = "PAY22112815152611300011"
//	//微信支付订单号:微信支付系统生成的订单号。
//	wxReq["trade_no"] = "2022061322001402060"
//	//与支付宝同步
//	wxReq["trade_status"] = "TRADE_SUCCESS"
//	wxReq["notify_time"] = "2022-11-28T14:00:20+08.00"
//	wxReq["total_amount"] = 100000
//	wxReq["receipt_amount"] = 100000
//	var mapData = make(map[string]interface{})
//	mapData["data_type"] = "PayNotify"
//	mapData["param"] = map[string]interface{}{"payType": "wxPay", "notifyReq": wxReq}
//	reqJson, _ := json.Marshal(mapData)
//	lib.PushMessToPayQueue(reqJson)*/
//
//	// 此写法是 gin 框架返回微信的写法
//	c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
//	fmt.Println("--------------------- WxPayNotify END ---------------------")
//	return
//}
//
//// WxJsAPI 支付，在微信支付服务后台生成预支付交易单
//func WxJsAPI(bm gopay.BodyMap) (string, string) {
//	wxRsp, err := client.V3TransactionJsapi(context.Background(), bm)
//	if err != nil {
//		fmt.Println(fmt.Sprintf("-----wxPay WxJsAPI() -------error:%s", err.Error()))
//		return "", err.Error()
//	}
//	return wxRsp.Response.PrepayId, wxRsp.Error
//}
//
//// WxNative 当面付 扫码支付，获取二维码
//func WxNative(bm gopay.BodyMap) string {
//	wxRsp, err := client.V3TransactionNative(context.Background(), bm)
//	if err != nil {
//		fmt.Println(fmt.Sprintf("-----wxPay WxNative() -------error:%s", err.Error()))
//		return ""
//	}
//	return wxRsp.Response.CodeUrl
//}
//
//// WxRefund 退款
//func WxRefund(bm gopay.BodyMap) (*wechat.RefundRsp, string) {
//	wxRsp, err := client.V3Refund(context.Background(), bm)
//	if err != nil {
//		fmt.Println(fmt.Sprintf("-----wxPay WxRefund() -------error:%s", err.Error()))
//		return nil, err.Error()
//	}
//	return wxRsp, wxRsp.Error
//}
//
//// WxQueryRefund 查询退款状态
//func WxQueryRefund(refundNo string) (*wechat.RefundQueryRsp, error) {
//	wxRsp, err := client.V3RefundQuery(context.Background(), refundNo, nil)
//	if err != nil {
//		fmt.Println(fmt.Sprintf("-----wxPay WxQueryRefund() -------error:%s", err.Error()))
//	}
//	return wxRsp, err
//}
//
//// WxTestV3Query 交易查询
//func WxTestV3Query(no string) *wechat.QueryOrderRsp {
//	wxRsp, err := client.V3TransactionQueryOrder(context.Background(), wechat.OutTradeNo, no)
//	if err != nil {
//		fmt.Println(fmt.Sprintf("-----wxPay TestV3QueryOrder() -------error:%s", err.Error()))
//		return nil
//	}
//	return wxRsp
//}

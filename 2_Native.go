package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func main() {

	var err error
	//初始化wechat v3 client
	ctx := context.Background()
	//商户私钥
	publicKey, err := utils.LoadPublicKeyWithPath("D:\\code\\sl2\\sl-management\\server\\resource\\wxPay\\pub_key.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	privateKey, err := utils.LoadPrivateKeyWithPath("D:\\code\\sl2\\sl-management\\server\\resource\\wxPay\\apiclient_key.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	opts := []core.ClientOption{
		option.WithWechatPayPublicKeyAuthCipher(
			"1693369919",
			"432B01701EB1DC4070DB35C97557FC69AE3C2342", privateKey,
			"PUB_KEY_ID_0116933699192024110700557300000107", publicKey),
	}
	clientV3, err := core.NewClient(ctx, opts...)
	if err != nil {
		fmt.Println(err)
		return
	}

	svc := native.NativeApiService{Client: clientV3}

	//创建
	resp, result, err := svc.Prepay(ctx,
		native.PrepayRequest{
			Appid:       core.String("wxc8153a453687bfb7"),                                    //
			Mchid:       core.String("1693369919"),                                            //
			Description: core.String("Image形象店-深圳腾大-QQ公仔"),                                    //
			OutTradeNo:  core.String("1217752501201407033233368040"),                          //
			TimeExpire:  core.Time(time.Now().Add(time.Minute)),                               //
			Attach:      core.String("自定义数据说明"),                                               //
			NotifyUrl:   core.String("https://www.cdszly.cn/slq-api/applet/common/payNotify"), //
			// LimitPay:      []string{"LimitPay_example"},
			SupportFapiao: core.Bool(false),
			Amount: &native.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(1),
			}},
	)
	if err != nil {
		// 处理错误
		fmt.Println("call Prepay err:", err)
		fmt.Println("status", result.Response.StatusCode)
		fmt.Println("resp", resp)
	} else {
		// 处理返回结果
		fmt.Println("status", result.Response.StatusCode)
		fmt.Println("resp", resp.CodeUrl)
		//log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}

	// 处理通知内容
	//fmt.Println(notifyReq.Summary)
	//fmt.Println(transaction.TransactionId)
}

//func Decrypt(c *gin.Context) (notifyReq *notify.Request, err error) {
//	var (
//		mchID                      = global.GVA_CONFIG.WeiXin.MchID                      // 商户号
//		mchCertificateSerialNumber = global.GVA_CONFIG.WeiXin.MchCertificateSerialNumber // 商户证书序列号
//		mchAPIv3Key                = global.GVA_CONFIG.WeiXin.MchAPIv3Key                // 商户APIv3密钥
//		privateKeyPath             = "./resource/wxPay/apiclient_key.pem"                // 商户私钥文件
//	)
//
//	ctx := c             //这个参数是context.Background()
//	request := c.Request //这个值是*http.Request
//
//	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
//	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(privateKeyPath)
//	if err != nil {
//		return nil, err
//	}
//
//	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
//	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, mchCertificateSerialNumber, mchID, mchAPIv3Key)
//	if err != nil {
//		return nil, err
//	}
//	// 2. 获取商户号对应的微信支付平台证书访问器
//	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
//	// 3. 使用证书访问器初始化 `notify.Handler`
//	handler, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
//	if err != nil {
//		return nil, err
//	}
//	transaction := new(payments.Transaction)
//	notifyReq, err = handler.ParseNotifyRequest(ctx, request, transaction)
//	// 如果验签未通过，或者解密失败
//	if err != nil {
//		fmt.Println(err)
//		//return
//	}
//	// 处理通知内容
//	fmt.Println(notifyReq.Summary)
//	fmt.Println(transaction.TransactionId)
//	// 如果验签未通过，或者解密失败
//	if err != nil {
//		return nil, err
//	}
//
//	return notifyReq, nil
//}

//---------------
// 发送请求，以下载微信支付平台证书为例
// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
//svc := certificates.CertificatesApiService{Client: client}
//resp, result, err := svc.DownloadCertificates(ctx)
//log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
// svc := native.NativeApiService{Client: client}
// svc := native.NativeApiService{Client: client}
// resp, result, err := svc.QueryOrderById(ctx,
// 	native.QueryOrderByIdRequest{
// 		TransactionId: core.String("1217752501201407033233368019"),
// 		Mchid:         core.String("1693369919"),
// 	},
// )

// if err != nil {
// 	// 处理错误
// 	log.Printf("call QueryOrderById err:%s", err)
// } else {
// 	// 处理返回结果
// 	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
// }

//resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
//	native.QueryOrderByOutTradeNoRequest{
//		OutTradeNo: core.String("1217752501201407033233368021"),
//		Mchid:      core.String("1693369919"),
//	},
//)
//
//if err != nil {
//	// 处理错误
//	log.Printf("call QueryOrderByOutTradeNo err:%s", err)
//} else {
//	// 处理返回结果
//	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
//}

//client.V3TransactionNative()

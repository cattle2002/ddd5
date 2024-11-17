package main

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"time"
)

func main() {

	var (
		wechatpayPublicKeyID string = "1" // 微信支付公钥ID
	)

	wechatpayPublicKey, err := utils.LoadPublicKeyWithPath("D:\\code\\sl2\\sl-management\\server\\resource\\wxPay\\pub_key.pem")
	if err != nil {
		panic(fmt.Errorf("load wechatpay public key err:%s", err.Error()))
	}
	priv, err := utils.LoadPrivateKeyWithPath("D:\\code\\sl2\\sl-management\\server\\resource\\wxPay\\apiclient_key.pem")
	if err != nil {
		panic(err)
	}
	//初始化 Client
	opts := []core.ClientOption{
		option.WithWechatPayPublicKeyAuthCipher(
			"2",
			"3", priv,
			wechatpayPublicKeyID, wechatpayPublicKey),
	}
	ctx := context.Background()
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		panic(err)
	}
	svc := native.NativeApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		native.PrepayRequest{
			Appid:       core.String("4"),
			Mchid:       core.String("5"),
			Description: core.String("Image形象店-深圳腾大-QQ公仔"),
			OutTradeNo:  core.String("1217752501201407033233368030"),
			TimeExpire:  core.Time(time.Now().Add(time.Hour)),
			Attach:      core.String("自定义数据说明"),
			NotifyUrl:   core.String("https://www.cdszly.cn/slq-api/applet/common/payNotify"),
			GoodsTag:    core.String("WXG"),
			// LimitPay:      []string{"LimitPay_example"},
			SupportFapiao: core.Bool(false),
			Amount: &native.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(1),
			},
			Detail: &native.Detail{
				CostPrice: core.Int64(608800),
				GoodsDetail: []native.GoodsDetail{native.GoodsDetail{
					GoodsName:        core.String("iPhoneX 256G"),
					MerchantGoodsId:  core.String("ABC"),
					Quantity:         core.Int64(1),
					UnitPrice:        core.Int64(828800),
					WechatpayGoodsId: core.String("1001"),
				}},
				InvoiceId: core.String("wx123"),
			},
			SettleInfo: &native.SettleInfo{
				ProfitSharing: core.Bool(false),
			},
			SceneInfo: &native.SceneInfo{
				DeviceId:      core.String("013467007045764"),
				PayerClientIp: core.String("14.23.150.211"),
				StoreInfo: &native.StoreInfo{
					Address:  core.String("广东省深圳市南山区科技中一道10000号"),
					AreaCode: core.String("440305"),
					Id:       core.String("0001"),
					Name:     core.String("腾讯大厦分店"),
				},
			},
		},
	)
	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

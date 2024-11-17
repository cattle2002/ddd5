package main

import (
	"context"
	"fmt"
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
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		native.QueryOrderByOutTradeNoRequest{
			OutTradeNo: core.String("20241112190321hlQgq8kN"),
			Mchid:      core.String("1693369919"),
		},
	)
	if err != nil {
		fmt.Println(resp, result.Response.StatusCode, err)
	}
	//if  resp.TradeState
	fmt.Println(resp, result.Response.StatusCode, err)

}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// 设置请求的URL
	url := "https://api.mch.weixin.qq.com/v3/pay/transactions/native"

	// 构建请求体
	requestBody := map[string]interface{}{
		"appid":          "wxc8153a453687bfb7",
		"mchid":          "1693369919",
		"description":    "Image形象店-深圳腾大-QQ公仔",
		"out_trade_no":   "1217752501201407033233368018",
		"time_expire":    "2025-06-08T10:34:56+08:00",
		"attach":         "自定义数据说明",
		"notify_url":     "https://www.weixin.qq.com/wxpay/pay.php",
		"goods_tag":      "WXG",
		"support_fapiao": false,
		"amount": map[string]interface{}{
			"total":    100,
			"currency": "CNY",
		},
		"detail": map[string]interface{}{
			"cost_price": 608800,
			"invoice_id": "微信123",
			"goods_detail": []interface{}{
				map[string]interface{}{
					"merchant_goods_id":  "1246464644",
					"wechatpay_goods_id": "1001",
					"goods_name":         "iPhoneX 256G",
					"quantity":           1,
					"unit_price":         528800,
				},
			},
		},
		"scene_info": map[string]interface{}{
			"payer_client_ip": "127.0.0.1",
			"device_id":       "013467007045764",
			"store_info": map[string]interface{}{
				"id":        "0001",
				"name":      "腾讯大厦分店",
				"area_code": "440305",
				"address":   "广东省深圳市南山区科技中一道10000号",
			},
		},
		"settle_info": map[string]interface{}{
			"profit_sharing": false,
		},
	}

	// 将请求体序列化为JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Add("Authorization", "WECHATPAY2-SHA256-RSA2048 mchid=\"1693369919\",...")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应体
	fmt.Println("Response:", string(body))

}

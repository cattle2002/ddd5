package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("order", func(context *gin.Context) {
		var mutex sync.Mutex
		mutex.Lock()
		fmt.Println("锁住")
		time.Sleep(time.Second * 3)
		mutex.Unlock()
		context.JSON(200, "下单成功")
	})
	r.Run(":9999")
}

package rate

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func DoRate() {
	//初始化 limiter 每秒10个令牌，令牌桶容量为20
	limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 20)

	for i := 0; i < 25; i++ {
		if limiter.Allow() {
			fmt.Println("success") //do something
		} else {
			fmt.Println("busy")
		}
	}

	//阻塞直到获取足够的令牌或者上下文取消
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	fmt.Println("start get token", time.Now())
	err := limiter.WaitN(ctx, 20)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("success get token", time.Now())
}

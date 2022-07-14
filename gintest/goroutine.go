package gintest

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func HasGo(c *gin.Context) {
	// 创建在 goroutine 中使用的副本
	cCp := c.Copy()
	go func() {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)

		// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
		log.Fatal("Done! in path " + cCp.Request.URL.Path)
	}()
}

func NoGo(c *gin.Context) {
	// 用 time.Sleep() 模拟一个长任务。
	time.Sleep(5 * time.Second)

	// 因为没有使用 goroutine，不需要拷贝上下文
	log.Println("Done! in path " + c.Request.URL.Path)
}

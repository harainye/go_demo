package routine

import (
	"fmt"
	"time"
)

func TestSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer func() {
			fmt.Println("====读取完成====")
		}()
		fmt.Println("====读取通道====")
		for {
			select {
			case d := <-ch1:
				fmt.Println(d)
			case d := <-ch2:
				fmt.Println(d)
			}
		}
	}()

	fmt.Println("====写入通道====")
	ch1 <- 1
	ch1 <- 2
	ch2 <- 11
	ch2 <- 21
	fmt.Println("====写入完成====")
	fmt.Printf("====睡眠3秒====")
	time.Sleep(3 * time.Second)
}

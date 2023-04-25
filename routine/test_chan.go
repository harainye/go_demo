package routine

import (
	"fmt"
	"time"
)

func TestChan() {
	c := make(chan int, 10)
	go send(c)
	go recv(c)

	time.Sleep(3 * time.Second)
	//close(c)
}

// 只能接收channel中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received ", i)
	}
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send readey ", i)
		c <- i
		fmt.Println("send ", i)
	}
}

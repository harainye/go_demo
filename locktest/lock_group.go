package locktest

import (
	"fmt"
	"sync"
)

func TestLock() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	fmt.Println("111")
	wg.Wait()
}

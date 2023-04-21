package _defer

import "fmt"

// 参考：https://www.jb51.net/article/264777.htm
// DeferCall defer遇见panic，但是并不捕获异常的情况
func DeferCall() {
	defer func() {
		fmt.Println("defer: panic 之前1")
	}()

	defer func() {
		fmt.Println("defer: panic 之前2")
	}()

	panic("异常内容") //触发defer出栈

	defer func() {
		fmt.Println("defer: panic 之后，永远执行不到")
	}()
}

// DeferCall2 defer遇见panic，并捕获异常
func DeferCall2() {
	defer func() {
		fmt.Println("defer: panic 之前1, 捕获异常")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()
	panic("异常内容") //触发defer出栈
	defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()
}

// DeferCall3 panic仅有最后一个可以被revover捕获
func DeferCall3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int)
	select {
	case <-ch1:
		fmt.Println("c1...")
	case ch2 <- 123:
		fmt.Println("c2...")
	default:
		fmt.Println("none...")
	}

	// 到计时
	t := time.NewTimer(time.Second)
	<-t.C // 从t取出数据，但是没有，等待1秒后往chan放入数据，才能执行
	fmt.Println("hello world!")
}

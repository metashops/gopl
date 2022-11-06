package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go Producer(ch)
	Consumer(ch) // 主协程作为消费者
}

// Producer 生产数据
func Producer(out chan<- string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(fmt.Sprintf("Producer:%v", i))
		out <- fmt.Sprintf("in:%v", i)
	}
	close(out)
}

// Consumer 消费数据
func Consumer(in <-chan string) {
	for {
		if val, ok := <-in; ok {
			fmt.Printf("Consumer:%v\n", val)
		} else {
			fmt.Println("正在忙碌中。。。")
			break
		}
	}

	// 不需要同步
	// for data := range in {
	// 	fmt.Printf("Consumer:%v\n", data)
	// }
}

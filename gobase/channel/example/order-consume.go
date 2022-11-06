package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan OrderInfo, 3)
	go ProducerOrder(ch)
	go ConsumerOrder(ch)
	time.Sleep(time.Second * 10)

}

type OrderInfo struct {
	Id   int
	Name string
}

// ProducerOrder 生产数据
func ProducerOrder(out chan<- OrderInfo) {
	for i := 1; i <= 10; i++ {
		order := OrderInfo{Id: i, Name: "诸葛亮"}
		fmt.Println(fmt.Sprintf("producer order:%v", order.Id))
		out <- order
	}
	close(out)
}

// ConsumerOrder 消费数据
func ConsumerOrder(in <-chan OrderInfo) {
	for val := range in {
		fmt.Printf("consumer order:%v\n", val.Id)
	}
}

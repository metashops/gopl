package main

import (
	`fmt`
)

func main() {
	naturals := make(chan int)
	squarers := make(chan int)
	go counter(naturals)
	go squarer(squarers, naturals)
	printer(squarers)
}

func counter(out chan<- int) { // 只发送的通道
	for x := 1; x <= 10; x++ {
		out <- x
	}
	close(out)
}

// in <-chan int 只能接收
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) { // 只接收
	for v := range in {
		fmt.Println("v=", v)
	}
}

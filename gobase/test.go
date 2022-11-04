package main

import (
	"fmt"
	"strconv"
	"sync"
)

func go1(ch1 chan int, ch2 chan int, wg *sync.WaitGroup, n int) {
	defer wg.Done()

	for i := 1; i <= n; i += 2 {
		<-ch1
		fmt.Println("go1: " + strconv.Itoa(i))
		ch2 <- 1
	}

	<-ch1
	close(ch2)
}

func go2(ch1 chan int, ch2 chan int, wg *sync.WaitGroup, n int) {
	defer wg.Done()

	ch1 <- 1
	for i := 2; i <= n; i += 2 {
		<-ch2
		fmt.Println("go2: " + strconv.Itoa(i))
		ch1 <- 1
	}

	<-ch2
	close(ch1)
}

func main1() {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)
	n := 10

	wg.Add(2)
	go go1(ch1, ch2, &wg, n)
	go go2(ch1, ch2, &wg, n)

	wg.Wait()
}

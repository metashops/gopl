package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go go1(ch1, ch2, &wg, 28)
	go go2(ch1, ch2, &wg, 'A')
	wg.Wait()
}

func go1(ch1, ch2 chan bool, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		<-ch1
		fmt.Print(i)
		fmt.Print(i + 1)
		ch2 <- true
	}
	<-ch1
	close(ch2)
}

func go2(ch1, ch2 chan bool, wg *sync.WaitGroup, n rune) {
	defer wg.Done()
	ch1 <- true
	for i := n; i <= 'Z'; i += 2 {
		<-ch2
		fmt.Print(string(i))
		fmt.Print(string(i+1), " ")
		ch1 <- true
	}
	<-ch2
	close(ch1)
}

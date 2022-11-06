package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	n := 28
	c := 'Z'
	wg.Add(2)
	go go1(ch1, ch2, &wg, n)
	go go2(ch1, ch2, &wg, c)
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

func go2(ch1, ch2 chan bool, wg *sync.WaitGroup, c rune) {
	defer wg.Done()
	ch1 <- true
	for i := 'A'; i <= c; i += 2 {
		<-ch2
		fmt.Print(string(i))
		fmt.Print(string(i+1), " ")
		ch1 <- true
	}
	<-ch2
	close(ch1)
}

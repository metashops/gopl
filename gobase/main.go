package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 12AB 34CD 56EF 78GH 910IJ 1112KL 1314MN 1516OP
// 1718QR 1920ST 2122UV 2324WX 2526YZ 2728

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	n := 28

	wg.Add(2)
	go g1(ch1, ch2, &wg, n)
	go g2(ch1, ch2, &wg, nil)

	wg.Wait()
}

func g1(ch1 chan bool, ch2 chan bool, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		<-ch1
		fmt.Print(strconv.Itoa(i))
		fmt.Print(strconv.Itoa(i + 1))
		ch2 <- true
	}
	<-ch1
	close(ch2)
}

func g2(ch1 chan bool, ch2 chan bool, wg *sync.WaitGroup, n []byte) {
	defer wg.Done()

	ch1 <- false
	for i := 'A'; i <= 'Z'; i += 2 {
		<-ch2
		fmt.Print(string(i))
		fmt.Print(string(i+1), " ")
		ch1 <- false
	}

	<-ch2
	close(ch1)
}

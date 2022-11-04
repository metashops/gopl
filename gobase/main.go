package main

import (
	`fmt`
	`strconv`
	`sync`
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	n := 3

	wg.Add(2)
	go g1(ch1, ch2, &wg, n)
	go g2(ch1, ch2, &wg, n)

	wg.Wait()
}

func g1(ch1 chan bool, ch2 chan bool, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		<-ch1
		fmt.Println("go1 => " + strconv.Itoa(i))
		ch2 <- true
	}
	<-ch1
	close(ch2)
}

func g2(ch1 chan bool, ch2 chan bool, wg *sync.WaitGroup, n int) {
	defer wg.Done()

	ch1 <- false
	for i := 2; i <= n; i += 2 {
		<-ch2
		fmt.Println("go2 => " + strconv.Itoa(i))
		ch1 <- false
	}

	<-ch2
	close(ch1)
}

package main

import (
	`fmt`
	`os`
	`time`
)

func main() {
	abort := make(chan struct{})
	fmt.Println("Commencing countdown...")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}
	go func() {
		os.Stdin.Read(make([]byte, 1)) // Read a char
		abort <- struct{}{}
	}()
}

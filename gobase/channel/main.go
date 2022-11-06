package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 1)
	defer close(ch)
	ch <- true
	c := <-ch
	fmt.Println("ch", c)
	time.Sleep(1 * time.Hour)
}

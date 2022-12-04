package main

import (
	"fmt"
)

func sum(a, b int) *int {
	sum := 0
	sum = a + b
	return &sum
}

func main() {
	a, b := 1, 2
	fmt.Printf("%v", *sum(a, b))
}

package main

import (
	"fmt"
)

func main() {
	s0 := []int{1, 2}
	s1 := append(s0, 3)
	fmt.Printf("s0:%p\n", s0) // s0:0xc0000ac010
	fmt.Printf("s1:%p\n", s1) // s1:0xc0000b6000

	s1[0] = 88

	fmt.Println("s0:", s0) // s0: [1 2]
	fmt.Println("s1:", s1) // s1: [88 2 3]
}

package main

import (
	"fmt"
)

func main() {
	s := []string{"诸葛亮", "关羽"}
	fmt.Printf("adder1:%p\n", s)
	s = append(s, "曹操")
	fmt.Printf("adder2:%p\n", s)
	s = append(s, "曹操")
	s = append(s, "曹操")
	fmt.Printf("adder3:%p\n", s)
	for key, val := range s {
		fmt.Println(key, "=>", val)
	}
}

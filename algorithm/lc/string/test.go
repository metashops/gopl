package main

import (
	"fmt"
)

func main() {
	str := "HellOWorlD" // 返回str is all lower char
	b := make([]byte, len(str))
	for i, _ := range str {
		s := str[i]
		if s >= 'A' && s <= 'Z' {
			s = s - 'A' + 'a'
		}
		b[i] = s
	}
	fmt.Println(str)      // 返回hellOWorlD
	fmt.Printf("%s\n", b) // 返回helloworld
}

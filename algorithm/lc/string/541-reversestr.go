package main

import (
	`fmt`
)

func main() {
	str := "abcdefg" // bacdge
	s := reverseStr(str, 2)
	fmt.Println(s)
}

func reverseStr(s string, k int) string {
	str := []byte(s)
	var begin int
	for begin < len(str) {
		left := begin
		right := begin + k - 1
		if right >= len(str) {
			right = len(str) - 1
		}
		for left < right && left < len(str) {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}
		begin = begin + 2*k
	}
	return string(str)
}

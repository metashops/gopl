package main

import (
	`fmt`
)

func main() {
	fmt.Println(Sqrt(4))
	l, r := 0, 4
	mid := l + (r-l)/2
	fmt.Println(mid)
}

func Sqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
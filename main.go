package main

import (
	`fmt`
)

func main() {
	// var m1 = make(map[string]int)
	var m1 = map[string]int{}
	m1["name"] = 123
	if m1 == nil {
		fmt.Println("m1 is nil")
	} else {
		fmt.Println(m1)
	}
}

func demo(nums []int, target int) []int {
	for k, _ := range nums {
		if k == 1 {
			nums[k] = target
		}
	}
	return nums
}

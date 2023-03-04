package main

import (
	`fmt`
)

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		if left <= right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}

func main() {
	n := []int{1, 2, 3, 3, 2, 5}
	v := 2
	elem := removeElement(n, v)
	fmt.Println(elem)
}

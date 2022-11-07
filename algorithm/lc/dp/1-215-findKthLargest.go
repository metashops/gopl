package main

import (
	"math"
)

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 注: 你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 本题希望我们返回数组排序之后的倒数第 k 个位置。

func main() {

}

// 朴素解法，直接维护一个 K 大的数组，替换掉最小的数
func findKthLargest(nums []int, k int) int {
	arr := append([]int{}, nums[:k]...)
	findMin := func(data []int) int {
		minIndex := -1
		min := math.MaxInt64
		for i, n := range data {
			if n < min {
				minIndex = i
				min = n
			}
		}
		return minIndex
	}
	index := findMin(arr)
	for i := k; i < len(nums); i++ {
		if arr[index] < nums[i] {
			arr[index] = nums[i]
			index = findMin(arr)
		}
	}

	return arr[findMin(arr)]
}

package main

import (
	"fmt"
)

func main() {
	nums := []int{-4, 8, -1, 0, 88, 3, 10}
	// sortnums := sort(sortedSquares(nums))
	// fmt.Println(sortnums) // 平方后，数组变为 [16,1,0,9,100] => 排序后，数组变为 [0,1,9,16,100]

	fmt.Println(sortedSquares2(nums))
}

func sortedSquares(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, nums[i]*nums[i])
	}
	return res
}

// 快速排序
func sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[0]
	var left, right []int
	for _, val := range nums[1:] {
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	return append(sort(left), append([]int{pivot}, sort(right)...)...)
}

// ========== 977. 有序数组的平方
func sortedSquares2(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)
	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			ans[k] = lm
			i++
		} else {
			ans[k] = rm
			j--
		}
		k--
	}
	return ans
}

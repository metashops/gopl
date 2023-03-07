package main

func minSubArrayLen(nums []int, target int) int {
	i := 0
	l := len(nums) // 数组长度
	sum := 0       // 子数组之和
	result := l + 1
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLen := j - i + 1
			if subLen < result {
				result = subLen
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}

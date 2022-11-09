package main

import (
	"fmt"
)

/**
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。
一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
请你计算并返回达到楼梯顶部的最低花费。

*/

func main() {
	min := minCostClimbingStairs([]int{10, 15, 20, 5})
	fmt.Println(min)
}

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int, length)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < length; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

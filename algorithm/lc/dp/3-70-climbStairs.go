package main

// func main() {
// 	res := climbStairs(4)
// 	fmt.Println(res)
// }

// 动态规划
func climbStairs(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 背包
func climbStairs1(n int) int {
	// 定义
	dp := make([]int, n+1)
	// 初始化
	dp[0] = 1
	// 本题物品只有两个1,2
	m := 2
	// 遍历顺序
	for j := 1; j <= n; j++ { // 先遍历背包
		for i := 1; i <= m; i++ { // 再遍历物品
			if j >= i {
				dp[j] += dp[j-i]
			}
			// fmt.Println(dp)
		}
	}
	return dp[n]
}

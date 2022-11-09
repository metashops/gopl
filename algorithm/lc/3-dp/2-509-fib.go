package main

/**
description:斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。
	公式：
		F(0) = 0，F(1) = 1
		F(n) = F(n - 1) + F(n - 2)，其中 n > 1
		给定 n ，请计算 F(n) 。
	示例 1：
		输入：n = 4
		输出：3
		解释：F(4) = F(3) + F(2) = 2 + 1 = 3
*/

// func main() {
// 	n := 8
// 	f1 := fib1(n)
// 	f2 := fib2(n)
// 	fmt.Println("fib1:", f1)
// 	fmt.Println("fib2:", f2)
// }

func fib1(n int) int {
	if n < 2 {
		return n
	}
	a := 0
	b := 1
	c := 0
	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

// 题目：n 项数列 [1，2，4，5，10，11，22]
//
//	当 n 是奇数时，是前面的平方，当n偶数时，前面值加1
//	例子：当 n = 8时，输出： 23
func fib2(n int) int {
	if n < 2 {
		return n
	}
	a := 1
	c := 0 // n=2=>c=2,a=2|n=3=>c=4,a=4|n=4,c=5,a=5
	for i := 2; i <= n; i++ {
		if i%2 != 0 {
			c = a * 2
			a = c
		} else {
			c = a + 1
			a = c
		}
	}
	return c
}

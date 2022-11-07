package main

import (
	"fmt"
)

// 给定两个从大到小排好序的数组A和B，给定K，找到A和B中第K大元素。

func main() {
	a1 := []int{1, 8, 3}
	a2 := []int{3, 5, 7}
	FindTopK(a1, a2, 3)
}

func FindTopK(A, B []int, K int) (int, error) {
	tmp := make([]int, 0)
	tmp = append(tmp, A...)
	tmp = append(tmp, B...)
	fmt.Println(tmp)
	return K, nil
}

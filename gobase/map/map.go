package main

import (
	"fmt"
)

// （1）map 循环是有序的还是无序的
// 为什么？：for range map 在开始处理循环逻辑的时候，就做了随机播种

func main() {
	m := make(map[string]int, 10)
	fmt.Println(m)
}

package main

/*
int sum(int a, int b) {
	return a + b;
}
*/
import (
	"C"
)

import (
	"fmt"
)

func main() {
	// 假如 sum 函数是 C 语言实现，在Go中如何调用呢？
	fmt.Println(C.sum(2, 3))
}

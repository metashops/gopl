// Code generated by cmd/cgo; DO NOT EDIT.

//line /Users/apple/manongyufu/study/gopl/gobase/cgo/cgo.go:1:1
package main

/*
int sum(int a, int b) {
	return a + b;
}
*/
import (
	_ "unsafe"
)

import (
	"fmt"
)

func main() {
	// 假如 sum 函数是 C 语言实现，在Go中如何调用呢？
	fmt.Println(( /*line :18:14*/ _Cfunc_sum /*line :18:18*/)(2, 3))
}

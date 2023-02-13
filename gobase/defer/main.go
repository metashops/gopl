package main

import (
	"fmt"
)

func main() {
	// 必须要先声明defer，否则不能捕获到panic异常
	defer func() {
		fmt.Println("into")
		if err := recover(); err != nil {
			// 这里的err其实就是panic传入的内容
			fmt.Println(err)
		}
		fmt.Println("drop out")
	}()
	f()                    // 开始调用f
	fmt.Println("end f()") // 这里开始下面代码不会再执行
}

func f() {
	fmt.Println("begin go...")
	panic("异常信息")
	fmt.Println("end go...") // 这里开始下面代码不会再执行
}

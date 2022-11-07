package main

import (
	"fmt"
)

// 面试题：for range 的时候它的地址会发生变化吗？如何解决？

type Student struct {
	Name string
	Age  int
}

func main() {
	studs := []Student{
		{Name: "John", Age: 21},
		{Name: "Tom", Age: 23},
		{Name: "JD", Age: 20},
	}
	s := make(map[string]*Student)

	// for _, v := range studs {
	// 	s[v.Name] = &v
	// 	// fmt.Printf("%s addr: %v\n", v.Name, &v.Age)
	// }

	// for k, v := range s {
	// 	fmt.Println(k, "=>", v.Age)
	// }
	// 运行结果；
	// John => 20
	// Tom => 20
	// JD => 20
	// 分析原因：
	// for range的时候，地址并没有发生变化。
	// 在循环时，会创建一个变量，每次都会把地址赋给同一个变量，
	// 导致循环结束后，拷贝的是最后一个

	// 解决：在每次循环时，创建一个临时变量

	for _, v := range studs {
		tmp := v
		s[tmp.Name] = &tmp
	}

	for k, v := range s {
		fmt.Println(k, "=>", v.Age)
	}

	// 结论：go中for range的时候，地址没有发生变化
}

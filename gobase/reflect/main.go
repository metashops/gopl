package main

import (
	"fmt"
	"reflect"
)

func MyAdd1(a, b int) int {
	return a + b
}

func MyAdd2(a, b int) int {
	return a - b
}

func CallAdd(f func(a, b int) int) {
	// 函数包装为反射值对象
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		return
	}
	// 构造函数参数, 传入两个整型值
	// argv := make([]reflect.Value, 2)
	// argv[0] = reflect.ValueOf(1)
	// argv[1] = reflect.ValueOf(2)
	argv := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}

	// 反射调用函数
	result := v.Call(argv)
	fmt.Println(result[0].Int())
}

func main() {
	CallAdd(MyAdd1)
	CallAdd(MyAdd2)
}

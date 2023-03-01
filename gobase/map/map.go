package main

import (
	"fmt"
)

// （1）map 循环是有序的还是无序的
// 为什么？：for range map 在开始处理循环逻辑的时候，就做了随机播种

func main() {
	// m := map[string]float64{"pi": 3.13}
	// if v, ok := m["pi"]; ok {
	// 	fmt.Println(v)
	// }
	// Get slices of keys and values from a map

	// myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	//
	// keys := make([]string, 0, len(myMap))
	// vls := make([]int, 0, len(myMap))
	//
	// for k, v := range myMap {
	// 	keys = append(keys, k)
	// 	vls = append(vls, v)
	// }
	//
	// for _, v := range keys {
	// 	fmt.Print(v, " ")
	// }

	var temp []int
	mp := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	for k := range mp {
		temp = append(temp, k)
	}

	for _, v := range temp {
		fmt.Println(v, mp[v])
	}
}

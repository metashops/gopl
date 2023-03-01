package main

import (
	`fmt`
)

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	ints := merge1(a, b, []int{}, 4, 4, len(a)+len(b))
	fmt.Println(ints)
}

func merge2(a,b []int) []int {
	c := make([]int, len(a)+len(b))

	for k:= 0; k<len(c);k++ {
		if
	}

}

// 输入递增有序数组a和b，合并后输出有序数组c。（注：a、b、c的空间已分配，不能分配额外数组空间）
func merge1(a, b, c []int, al, bl, cl int) []int {
	cl = al + bl
	c = make([]int, cl)
	copy(c, a)
	copy(c[len(a):], b)
	return c
}

func merge(a, b, c []int, aLen, bLen, cLen int) {
	// 实现
	i, j, k := 0, 0, 0 // 分别指向a、b、c的第一位
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[k] = a[i] // 当前元素较小的放入c中
			k++
			i++
		} else {
			c[k] = b[j] // 当前元素较大放入c中
			k++
			j++
		}
	}
	// 剩余内容复制到c中
	copy(c[k:], a[i:])
	copy(c[k:], b[j:])
}

func MergeArrays(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	i, j := 0, 0
	for k := 0; k < len(c); k++ {
		if i >= len(a) {
			c[k] = b[j]
			j++
			continue
		}

		if j >= len(b) {
			c[k] = a[i]
			i++
			continue
		}

		if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
	}

	return c
}

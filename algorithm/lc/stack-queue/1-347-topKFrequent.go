package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 1, 1, 2, 3, 3, 3, 3}
	brr := []int{5, 5, 5, 5, 5, 6, 7, 8, 8, 8, 8, 9, 8, 8, 0}
	multi := append(append(arr, brr...))
	fmt.Println(multi)
	tf := topKFrequent(multi, 2)
	fmt.Println(tf)
}

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

// (1)要统计元素出现频率;(2)对频率排序;(3)找出前K个高频元素
// 大顶堆（堆头是最大元素），小顶堆（堆头是最小元素）
// 使用小根堆, 先用map统计词频, 然后放到小根堆里面, 但是要优于O(N * logN),
// 所以, 小根堆只开辟k个大小, 这样小根堆中push和pop时时间复杂度就是O(logk)

func topKFrequent(nums []int, k int) []int {
	map_num := make(map[int]int)
	// 记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++ // map_num[item] = map_num[item] + 1
	}

	// 放到小根堆里面
	h := &NewHeapMin{}
	heap.Init(h)
	topK := min(k, len(map_num))

	size := 0
	for k, v := range map_num {
		if size < topK {
			heap.Push(h, &HeapMin{
				data: k,
				size: v,
			})
			size++
		} else {
			if v > (*h)[0].size {
				heap.Pop(h)
				heap.Push(h, &HeapMin{
					data: k,
					size: v,
				})
			}
		}
	}
	// 3.收集答案
	res := make([]int, 0, topK)
	for i := 0; i < topK; i++ {
		res = append(res, heap.Pop(h).(*HeapMin).data)
	}
	return res
}

// HeapMin 构建小顶堆
type HeapMin struct {
	data int
	size int
}

type NewHeapMin []*HeapMin

func (h NewHeapMin) Len() int {
	return len(h)
}

func (h NewHeapMin) Less(i, j int) bool {
	return h[i].size < h[j].size
}

func (h NewHeapMin) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NewHeapMin) Push(x interface{}) {
	*h = append(*h, x.(*HeapMin))
}

func (h *NewHeapMin) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}

func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

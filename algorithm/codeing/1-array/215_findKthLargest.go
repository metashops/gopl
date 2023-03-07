package main

/**
215. 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

// TODO:
func findKthLargestK(nums []int, k int) int {
	nums = heapSort(nums)
	return nums[k-1]
}

func heapSort(nums []int) []int {
	lens := len(nums)
	// 将一个数组调整为小顶堆
	// 建堆	lens/2 后面都是叶子节点，不需要调整down()
	for i := lens / 2; i >= 0; i-- {
		down(nums, i, lens)
	}
	// 将小根堆堆顶排到切片末尾(降序)
	for i := lens - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		lens--
		down(nums, 0, lens)
	}
	return nums
}

// 小根堆
// []int 是待调整的数组
// i 表示非叶子结点在数组中索引
// lens 表示对多少个元素继续调整（lens是逐渐减少的）
func down(nums []int, i, lens int) {
	min := i         // i 父节点
	left := 2*i + 1  // 左孩子
	right := 2*i + 2 // 右孩子
	if left < lens && nums[left] < nums[min] {
		min = left
	}
	if right < lens && nums[right] < nums[min] {
		min = right
	}
	if min != i {
		Swap(nums, min, i)
		down(nums, min, lens)
	}
}

func Swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

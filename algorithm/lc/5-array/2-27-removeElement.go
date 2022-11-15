package main

// func main() {
// 	nums := []int{3, 2, 2, 2, 3}
// 	fmt.Println(removeElement(nums, 3))
// }

// 移除元素
func removeElement(nums []int, val int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

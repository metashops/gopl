package main

// func main() {
// 	arr := []int{1, 43, 5, 94, 90}
// 	BubbleSort(arr)
// 	fmt.Println(arr)
//
// }

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

package main

import (
	"fmt"
)

// func main() {
// 	// s := "anagram"
// 	// t := "nagaram"
// 	// anagram := isAnagram(s, t)
// 	// fmt.Println(anagram)
//
// }

// 242. 有效的字母异位词
func isAnagram(s string, t string) bool {
	record := [26]int{}
	for _, r := range s {
		record[r-rune('a')]++
	}
	fmt.Println(record)
	for _, r := range t {
		record[r-rune('a')]--
	}
	return record == [26]int{}
}

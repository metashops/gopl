package main

import (
	"fmt"
)

func main() {
	s := "abab" // true
	fmt.Println(repeatedSubstringPattern(s))
}

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if len(s) == 0 {
		return false
	}
	next := make([]int, len(s))
	j := -1
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	if next[n-1] != -1 && n%(n-(next[n-1]+1)) == 0 {
		return true
	}
	return false
}

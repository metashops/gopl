package main

import (
	"fmt"
)

/**
description: 给定一个字符串，逐个翻转字符串中的每个单词。
example：
输入：s = "the sky is blue"
输出："blue is sky the"
*/

// func main() {
// 	s := " the sky is blue"
// 	fmt.Println(s)
// 	words := reverseWords(s)
// 	fmt.Println("a", words)
//
// }

func reverseWords(s string) string {
	// 1.使用双指针删除冗余的空格
	slowIndex := 0
	fastIndex := 0
	b := []byte(s)

	// （1）删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	fmt.Println("fast:", fastIndex)

	// （2）删除 单词间 冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}

	// （3）删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}

	// 2.反转整个字符串
	reverse(&b, 0, len(b)-1)
	// 3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

// 写法二
func reverseWords3(s string) (res string) {
	s = " " + s + " "
	l, r := len(s)-1, len(s)-1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == ' ' {
			l, r = i, l
			if r > l+1 {
				res = res + s[l+1:r] + " "
			}
		}
	}
	return res[:len(res)-1]
}

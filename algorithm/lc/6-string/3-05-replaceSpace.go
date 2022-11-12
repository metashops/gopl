package main

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例：
输入：s = "We are happy."
输出："We%20are%20happy."
*/

// func main() {
// 	s := "We are happy."
// 	space := replaceSpace(s)
// 	fmt.Println(space)
// }

func replaceSpace(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}

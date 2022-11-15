package main

func reverseWords2(s string) (res string) {
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

// func main() {
// 	s := "this is go language"
// 	res := reverseWords2(s)
// 	fmt.Println(res)
//
// 	t := " " + s + " "
// 	fmt.Println("t1:", len(t)-1)
// 	fmt.Println("t2:", len(s)-1)
// }

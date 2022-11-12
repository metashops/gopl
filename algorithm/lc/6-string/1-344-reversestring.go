package main

func ReverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
func ReverseString2(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left] ^= s[right]
		s[right] ^= s[left]
		s[left] ^= s[right]
		left++
		right--
	}
}

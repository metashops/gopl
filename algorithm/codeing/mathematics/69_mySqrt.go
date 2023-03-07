package mathematics

func Sqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

// 精确到3位小数
func mySqrt(x float64) float64 {
	l, r := 0.0, x
	for l <= r {
		mid := (l + r) / 2
		if x < mid*mid {
			r = mid - 1e-3
		} else {
			l = mid + 1e-3
		}
	}
	return r
}

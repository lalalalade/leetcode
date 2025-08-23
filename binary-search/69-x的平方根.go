package binary_search

func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r) >> 1
		if mid*mid >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l*l == x {
		return l
	}
	return l - 1
}

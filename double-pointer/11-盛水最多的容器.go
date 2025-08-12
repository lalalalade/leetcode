package double_pointer

func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

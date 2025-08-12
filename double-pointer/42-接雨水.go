package double_pointer

func trap(height []int) int {
	n := len(height)
	preMax := make([]int, n)
	preMax[0] = height[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	sufMax := make([]int, n)
	sufMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}

	res := 0
	for i, x := range height {
		res += min(preMax[i], sufMax[i]) - x
	}
	return res
}

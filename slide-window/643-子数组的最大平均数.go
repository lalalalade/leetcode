package slide_window

import "math"

func findMaxAverage(nums []int, k int) float64 {
	res := math.MinInt
	s := 0
	for i, x := range nums {
		s += x
		if i < k-1 {
			continue
		}
		res = max(res, s)
		s -= nums[i-k+1]
	}
	return float64(res) / float64(k)
}

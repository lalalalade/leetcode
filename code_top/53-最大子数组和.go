package code_top

import "math"

func maxSubArray(nums []int) int {
	res := math.MinInt
	f := 0
	for _, x := range nums {
		f = max(f, 0) + x
		res = max(res, f)
	}
	return res
}

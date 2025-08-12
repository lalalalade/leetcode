package prefix_sum

import "math"

func maxSubArray(nums []int) int {
	res := math.MinInt
	minPreSum := 0
	preSum := 0
	for _, x := range nums {
		preSum += x
		res = max(res, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}
	return res
}

package slide_window

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	res := 0
	left := 0
	prod := 1
	for right, x := range nums {
		prod *= x
		for prod >= k {
			prod /= nums[left]
			left += 1
		}
		res += right - left + 1
	}
	return res
}

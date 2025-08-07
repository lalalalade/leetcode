package slide_window

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := n + 1
	sum, left := 0, 0
	for right, x := range nums {
		sum += x
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res <= n {
		return res
	}
	return 0
}

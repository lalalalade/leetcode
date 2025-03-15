package code_top

// findDuplicates 数组中重复的数据
func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		} else {
			res = append(res, x)
		}
	}
	return res
}

package double_pointer

func twoSum(nums []int, target int) []int {
	var res []int
	l, r := 0, len(nums)-1
	for l < r {
		s := nums[l] + nums[r]
		if s == target {
			res = append(res, l+1, r+1)
			break
		}
		if s > target {
			r--
		} else {
			l++
		}
	}
	return res
}

package double_pointer

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range n {
		for 1 <= nums[i] && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i := range n {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

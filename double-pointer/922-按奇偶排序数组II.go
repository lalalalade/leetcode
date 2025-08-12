package double_pointer

func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	n := len(nums)
	for even < n && odd < n {
		if nums[n-1]&1 == 1 {
			nums[odd], nums[n-1] = nums[n-1], nums[odd]
			odd += 2
		} else {
			nums[even], nums[n-1] = nums[n-1], nums[even]
			even += 2
		}
	}
	return nums
}

package double_pointer

func removeDuplicatesII(nums []int) int {
	n := len(nums)
	stackSize := 2
	for i := 2; i < n; i++ {
		if nums[i] != nums[stackSize-2] {
			nums[stackSize] = nums[i]
			stackSize++
		}
	}
	return min(stackSize, n)
}

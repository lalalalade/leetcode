package double_pointer

func removeDuplicates(nums []int) int {
	n := len(nums)
	k := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

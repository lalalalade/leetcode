package double_pointer

func removeElement(nums []int, val int) int {
	n := len(nums)
	k := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

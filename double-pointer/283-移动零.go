package double_pointer

func moveZeroes(nums []int) {
	j := 0
	for i, x := range nums {
		if x != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

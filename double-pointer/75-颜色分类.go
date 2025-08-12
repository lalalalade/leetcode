package double_pointer

func sortColors(nums []int) {
	p0, p1 := 0, 0
	// 假设从nums[0] ~ nums[i-1]有序
	for i, x := range nums {
		nums[i] = 2
		if x <= 1 {
			nums[p1] = 1
			p1++
		}
		if x == 0 {
			nums[p0] = 0
			p0++
		}
	}
}

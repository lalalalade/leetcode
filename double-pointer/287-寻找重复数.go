package double_pointer

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

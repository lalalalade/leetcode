package double_pointer

import "slices"

func triangleNumber(nums []int) int {
	res := 0
	slices.Sort(nums)
	for k := 2; k < len(nums); k++ {
		c := nums[k]
		i := 0
		j := k - 1
		for i < j {
			if nums[i]+nums[j] > c {
				res += j - i
				j--
			} else {
				i++
			}
		}
	}
	return res
}

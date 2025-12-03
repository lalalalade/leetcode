package codetop4

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		res[0] = l
	}
	l, r = 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if nums[l] == target {
		res[1] = l
	}
	return res
}

package codetop1

// search 搜索旋转排序数组
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, n-1
	// 二分找到左半段的最后一个值
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] >= nums[0] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if target >= nums[0] {
		l = 0
	} else {
		l = r + 1
		r = n - 1
	}
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[r] == target {
		return r
	}
	return -1
}

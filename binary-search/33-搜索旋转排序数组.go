package binary_search

func search(nums []int, target int) int {
	// 找最大值
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] >= nums[0] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	idx := l
	if target == nums[0] {
		return 0
	}
	if target > nums[0] {
		l, r = 0, idx
	} else {
		l, r = idx+1, n-1
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

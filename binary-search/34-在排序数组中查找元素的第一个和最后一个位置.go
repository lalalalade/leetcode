package binary_search

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	l, r := 0, len(nums)-1
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

// 二分模板
func check(mid int) bool {
	return true
}

// SL 查找左边界
func SL(l, r int) int {
	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// SR 查找右边界限
func SR(l, r int) int {
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

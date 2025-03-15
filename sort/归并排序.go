package sort

func sortArray1(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) >> 1
	// 递归处理子问题
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, r)
	tmp := make([]int, r-l+1)
	i, j, k := l, m+1, 0
	for i <= m && j <= r {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= m {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= r {
		tmp[k] = nums[j]
		j++
		k++
	}
	copy(nums[l:r+1], tmp)
}

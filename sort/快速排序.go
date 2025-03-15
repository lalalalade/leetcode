package sort

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	x := nums[(l+r)>>1]
	i, j := l-1, r+1
	for i < j {
		for i++; nums[i] < x; i++ {
		}
		for j--; nums[j] > x; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}

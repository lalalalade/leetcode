package codetop1

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickSelect(nums, 0, n-1, k-1)
}

func quickSelect(nums []int, l, r, k int) int {
	if l >= r {
		return nums[k]
	}
	p := nums[(l+r)/2]
	i, j := l-1, r+1
	for i < j {
		for i++; nums[i] > p; i++ {
		}
		for j--; nums[j] < p; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickSelect(nums, l, j, k)
	} else {
		return quickSelect(nums, j+1, r, k)
	}
}

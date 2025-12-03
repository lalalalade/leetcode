package codetop3

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	q := []int{}

	for i, x := range nums {
		// 1. 右边入
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}

		q = append(q, i)

		left := i - k + 1
		if q[0] < left {
			q = q[1:]
		}

		if left >= 0 {
			res[left] = nums[q[0]]
		}
	}
	return res
}

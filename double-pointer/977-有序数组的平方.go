package double_pointer

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	i, j := 0, n-1
	for p := n - 1; p >= 0; p-- {
		x := nums[i] * nums[i]
		y := nums[j] * nums[j]
		if x > y {
			res[p] = x
			i++
		} else {
			res[p] = y
			j--
		}
	}
	return res
}

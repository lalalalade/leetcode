package backtrack

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	path := make([]int, n)
	onPath := make([]bool, n)
	var dfs func(int)
	// i 代表要填path的第几个数
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int{}, path...))
			return
		}
		for j, on := range onPath {
			if !on {
				path[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return res
}

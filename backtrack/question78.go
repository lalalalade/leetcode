package backtrack

import "slices"

// subsets 子集
func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, slices.Clone(path))
			return
		}

		// 不选nums[i]
		dfs(i + 1)

		// 选nums[i]
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1] // 恢复
	}
	dfs(0)
	return res
}

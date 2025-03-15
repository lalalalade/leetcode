package backtrack

import "slices"

// combinationSum 组合总和
func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int, int)
	// i 代表枚举值的数组下标， t 代表剩余的和
	dfs = func(i, t int) {
		if t == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if t < candidates[i] {
			return
		}
		// 从小到大枚举
		for j := i; j < len(candidates); j++ {
			path = append(path, candidates[j])
			dfs(j, t-candidates[j])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return res
}

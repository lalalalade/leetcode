package backtrack

import "slices"

// combinationSum2 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int, int)
	dfs = func(i, t int) {
		if t == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for j := i; j < len(candidates) && candidates[j] <= t; j++ {
			// 去重
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			path = append(path, candidates[j])
			dfs(j+1, t-candidates[j])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return res
}

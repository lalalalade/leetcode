package codetop4

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if i == len(candidates) || left < 0 {
			return
		}

		dfs(i+1, left)

		path = append(path, candidates[i])
		dfs(i, left-candidates[i])
		path = path[:len(path)-1]
	}
	dfs(0, target)
	return res
}

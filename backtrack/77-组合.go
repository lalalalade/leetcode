package backtrack

// combine 组合
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	// 某一路径的结果
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(i int) {
		// 还要选 d 个数
		d := k - len(path)
		if d == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		// j 至少要是d 否则不够选
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j - 1)
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return res
}

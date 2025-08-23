package backtrack

func generateParenthesis(n int) []string {
	m := 2 * n
	res := make([]string, 0)
	path := make([]byte, m)
	var dfs func(int, int)
	dfs = func(i int, open int) {
		if i == m {
			res = append(res, string(path))
			return
		}
		if open < n {
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open {
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return res
}

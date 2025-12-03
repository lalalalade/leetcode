package codetop3

func generateParenthesis(n int) []string {
	path := make([]byte, 2*n)
	res := make([]string, 0)

	var dfs func(int, int)
	dfs = func(left, right int) {
		if right == n {
			res = append(res, string(path))
			return
		}
		if left < n {
			path[left+right] = '('
			dfs(left+1, right)
		}
		if right < left {
			path[left+right] = ')'
			dfs(left, right+1)
		}
	}
	dfs(0, 0)
	return res
}

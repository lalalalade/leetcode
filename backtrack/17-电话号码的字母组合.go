package backtrack

var mapping = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

// letterCombinations 电话号码的字母组合
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	// 要枚举几个数字
	n := len(digits)
	if n == 0 {
		return res
	}
	// 路径长度确定，一开始就是n
	path := make([]byte, n)
	var dfs func(int)
	// i 代表正在枚举第几个数字
	dfs = func(i int) {
		if i == n {
			res = append(res, string(path))
			return
		}
		for _, c := range mapping[digits[i]-'0'] {
			// 直接覆盖
			path[i] = byte(c)
			dfs(i + 1)
		}
	}
	dfs(0)
	return res
}

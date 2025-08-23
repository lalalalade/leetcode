package backtrack

// partition 分割回文串
func partition(s string) [][]string {
	res := make([][]string, 0)
	n := len(s)
	path := make([]string, 0)

	var dfs func(int)
	// i代表枚举起点
	dfs = func(i int) {
		if i == n {
			// 完成了一次分割 收集结果
			res = append(res, append([]string{}, path...))
			return
		}
		// 枚举子串结束位置
		for j := i; j < n; j++ {
			if isPalindrome(s, i, j) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

// 回溯三问：

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

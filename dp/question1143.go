package dp

// longestCommonSubsequence 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i, x := range text1 {
		for j, y := range text2 {
			if x == y {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[n][m]
}

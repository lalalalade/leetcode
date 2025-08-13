package code_top

// longestPalindrome 最长回文子串
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	res := s[0:1]
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for length := 2; length <= len(s); length++ {
		for start := 0; start < len(s)-length+1; start++ {
			end := start + length - 1
			if s[start] != s[end] {
				continue
			} else if length < 3 {
				dp[start][end] = true
			} else {
				dp[start][end] = dp[start+1][end-1]
			}
			if dp[start][end] && (end-start+1) > len(res) {
				res = s[start : end+1]
			}
		}
	}
	return res
}

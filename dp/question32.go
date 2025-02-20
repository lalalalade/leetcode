package dp

func longestValidParentheses(s string) int {
	n := len(s)
	// dp[i]表示以s[i]结尾的最长有效括号子串的长度
	dp := make([]int, n+1)
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					dp[i] = 2 + dp[i-2]
				} else {
					dp[i] = 2
				}
			} else {
				length := dp[i-1]
				if i-1-length >= 0 && s[i-1-length] == '(' {
					if i-1-length > 0 {
						dp[i] = dp[i-1] + 2 + dp[i-2-length]
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
	}
	res := 0
	for _, num := range dp {
		if num > res {
			res = num
		}
	}
	return res
}

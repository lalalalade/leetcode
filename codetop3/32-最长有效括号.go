package codetop3

func longestValidParentheses(s string) int {
	res := 0
	n := len(s)
	dp := make([]int, n+1)
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				length := dp[i-1]
				if i-1-length >= 0 && s[i-1-length] == '(' {
					if i-2-length >= 0 {
						dp[i] = dp[i-1] + 2 + dp[i-2-length]
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
	}
	for _, num := range dp {
		if num > res {
			res = num
		}
	}
	return res
}

package code_top

// lengthOfLongestSubstring 最长公共子串
func lengthOfLongestSubstring(s string) int {
	cnt := [128]int{}
	left := 0
	res := 0
	for right, c := range s {
		cnt[c]++
		for cnt[c] > 1 {
			cnt[s[left]]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

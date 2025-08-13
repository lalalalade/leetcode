package code_top

// lengthOfLongestSubstring 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	// 记录字符出现的次数 每个字符为rune类型，即int32
	cnt := [128]int{}
	left := 0
	res := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		cnt[c]++
		for cnt[c] > 1 {
			cnt[s[left]]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

package slide_window

func lengthOfLongestSubstring(s string) int {
	res := 0
	cnt := [128]int{}
	left := 0
	for right, v := range s {
		cnt[v]++
		for cnt[v] > 1 {
			cnt[s[left]]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

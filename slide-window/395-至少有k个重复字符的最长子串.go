package slide_window

func longestSubstring(s string, k int) int {
	cnts := [128]int{}
	res := 0
	for require := 1; require <= 26; require++ {
		cnts = [128]int{}
		l, collect, satisfy := 0, 0, 0
		for r, x := range s {
			cnts[x]++
			if cnts[x] == 1 {
				collect++
			}
			if cnts[x] == k {
				satisfy++
			}
			// 种类超过了
			for collect > require {
				if cnts[s[l]] == 1 {
					collect--
				}
				if cnts[s[l]] == k {
					satisfy--
				}
				cnts[s[l]]--
				l++
			}
			if satisfy == require {
				res = max(res, r-l+1)
			}
		}
	}
	return res
}

package slide_window

import "math"

func minWindow(s string, t string) string {
	cnts := [128]int{}
	for _, v := range t {
		cnts[v]--
	}
	res := math.MaxInt
	start := 0
	debt := len(t)
	l := 0
	for r, v := range s {
		cnts[v]++
		if cnts[v]-1 < 0 {
			debt--
		}
		if debt == 0 {
			for cnts[s[l]] > 0 {
				cnts[s[l]]--
				l++
			}
			if (r - l + 1) < res {
				res = r - l + 1
				start = l
			}
		}
	}
	if res == math.MaxInt {
		return ""
	}
	return s[start : start+res]
}

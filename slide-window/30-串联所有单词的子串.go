package slide_window

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	windowLen := wordLen * len(words)
	res := make([]int, 0)
	for start := range wordLen {
		cnts := map[string]int{}
		for _, w := range words {
			cnts[w]--
		}
		debt := len(words)
		l := start
		for r := start + wordLen - 1; r < len(s); r += wordLen {
			inWord := s[r-wordLen+1 : r+1]
			cnts[inWord]++
			if cnts[inWord]-1 < 0 {
				debt--
			}
			if r-l+1 < windowLen {
				continue
			}
			if debt == 0 {
				res = append(res, l)
			}
			outWord := s[l : l+wordLen]
			cnts[outWord]--
			l += wordLen
			if cnts[outWord] < 0 {
				debt++
			}
		}
	}
	return res
}

package slide_window

func maxVowels(s string, k int) int {
	res := 0
	vowel := 0
	for i, in := range s {
		// 1. 进入窗口
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			vowel++
		}
		if i < k-1 { // 窗口大小不足 k
			continue
		}
		// 2. 更新答案
		res = max(res, vowel)
		// 3. 离开窗口，为下一个循环做准备
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			vowel--
		}
	}
	return res
}

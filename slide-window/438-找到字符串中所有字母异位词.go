package slide_window

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	cnts := [128]int{}
	for _, v := range p {
		cnts[v]--
	}
	debt := len(p)
	for i, in := range s {
		// 进入窗口
		cnts[in]++
		if cnts[in]-1 < 0 {
			debt--
		}
		// 更新
		if i < len(p)-1 {
			continue
		}
		if debt == 0 {
			res = append(res, i-len(p)+1)
		}
		// 离开窗口
		out := s[i-len(p)+1]
		cnts[out]--
		if cnts[out] < 0 {
			debt++
		}
	}
	return res
}

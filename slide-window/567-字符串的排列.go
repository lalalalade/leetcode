package slide_window

func checkInclusion(s1 string, s2 string) bool {

	cnts := [128]int{}
	for _, v := range s1 {
		cnts[v]--
	}
	debt := len(s1)
	for i, x := range s2 {
		// 进入窗口
		cnts[x]++
		if cnts[x]-1 < 0 {
			debt--
		}
		if i < len(s1)-1 {
			continue
		}
		if debt == 0 {
			return true
		}
		// 离开窗口
		out := s2[i-len(s1)+1]
		cnts[out]--
		if cnts[out] < 0 {
			debt++
		}
	}
	return false
}

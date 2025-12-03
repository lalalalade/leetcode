package codetop2

import "slices"

func merge1(intervals [][]int) [][]int {
	res := make([][]int, 0)
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })
	for _, p := range intervals {
		m := len(res)
		if m > 0 && p[0] <= res[m-1][1] {
			res[m-1][1] = max(res[m-1][1], p[1])
		} else {
			res = append(res, p)
		}
	}
	return res
}

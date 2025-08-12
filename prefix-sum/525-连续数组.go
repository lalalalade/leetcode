package prefix_sum

func findMaxLength(nums []int) int {
	res := 0
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		if x == 0 {
			s[i+1] = s[i] - 1
		} else {
			s[i+1] = s[i] + 1
		}
	}
	// idx表示某个子数组的和最早出现的位置
	idx := map[int]int{}
	// s[j] - s[i] = 0 -> s[j] = s[i]
	for j, sj := range s {
		if v, ok := idx[sj]; ok {
			res = max(res, j-v)
		} else {
			idx[sj] = j
		}
	}
	return res
}

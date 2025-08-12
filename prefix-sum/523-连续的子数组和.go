package prefix_sum

func checkSubarraySum(nums []int, k int) bool {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = (s[i] + x) % k
	}
	idx := make(map[int]int)
	for j, x := range s {
		if v, ok := idx[x]; ok {
			if j-v > 1 {
				return true
			}
		} else {
			idx[x] = j
		}
	}
	return false
}

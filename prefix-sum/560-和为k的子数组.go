package prefix_sum

func subarraySum(nums []int, k int) int {
	// 前缀和数组
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	res := 0
	cnt := make(map[int]int, len(s))
	for _, sj := range s {
		res += cnt[sj-k]
		cnt[sj]++
	}
	return res
}

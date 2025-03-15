package backtrack

import "slices"

// permuteUnique 全排列 II
func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	n := len(nums)
	path := make([]int, n)
	onPath := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int{}, path...))
			return
		}
		for j, on := range onPath {
			if on || j > 0 && nums[j] == nums[j-1] && !onPath[j-1] {
				continue
			}
			path[i] = nums[j]
			onPath[j] = true
			dfs(i + 1)
			onPath[j] = false
		}
	}
	dfs(0)
	return res
}

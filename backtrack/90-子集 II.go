package backtrack

import "slices"

// subsetsWithDup 子集 II
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	n := len(nums)
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(i int) {
		res = append(res, append([]int{}, path...))
		for j := i; j < n; j++ {
			// 去重
			if j > i && nums[j] == nums[j-1] {
				continue
			}
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

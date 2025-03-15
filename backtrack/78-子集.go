package backtrack

// subsets 子集
func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		// 收集所有节点
		res = append(res, append([]int{}, path...))
		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

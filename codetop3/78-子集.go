package codetop3

import "slices"

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

// 子集型回溯
// 每个位置选谁？
// 1. 当前操作：枚举一个下标j>=i的数字，加入path
// 2. 子问题：从下标>=i的数字中构造子集
// 3. 下一个子问题：从下标>=j+1的数字中构造子集

func subsets2(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	n := len(nums)
	var dfs func(int)
	// dfs(i)代表 当前考虑到 nums[i] 选或不选。
	dfs = func(i int) {
		if i == n {
			res = append(res, slices.Clone(path))
			return
		}
		// 不选
		dfs(i + 1)
		// 选
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}
	dfs(0)
	return res
}

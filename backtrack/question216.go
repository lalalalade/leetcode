package backtrack

// combinationSum3 组合总和 III
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int, int)
	// i 代表枚举的值， t 代表剩余的和
	dfs = func(i, t int) {
		// 还要选d个数
		d := k - len(path)
		// 从i开始往前枚举d个数 最大的总和为 (i*2-d+1)*d/2 ==> i, i-1, i-2, ... i-d+1
		if t < 0 || t > (i*2-d+1)*d/2 {
			return
		}
		if d == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j-1, t-j)
			path = path[:len(path)-1]
		}
	}
	dfs(9, n)
	return res
}

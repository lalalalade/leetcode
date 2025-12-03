package codetop2

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lVal := dfs(root.Left)
		rVal := dfs(root.Right)
		res = max(res, lVal+rVal+root.Val)
		return max(max(lVal, rVal)+root.Val, 0)
	}
	dfs(root)
	return res
}

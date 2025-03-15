package binary_tree

import "math"

// maxPathSum 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var dfs func(*TreeNode) int
	// dfs返回从某个节点到当前节点的路径中节点值之最大和 如果最大和为负数，返回0，不要这条链
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 可能从左子树过来，也可能从右子树过来
		lVal := dfs(node.Left)  // 左子树最大链和
		rVal := dfs(node.Right) // 右子树最大链和
		// 更新直径值
		res = max(res, lVal+rVal+node.Val)
		return max(max(lVal, rVal)+node.Val, 0)
	}
	dfs(root)
	return res
}

package binary_tree

// hasPathSum 路径总和 自顶向下DFS
func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(root *TreeNode, target int) bool
	dfs = func(root *TreeNode, target int) bool {
		// 只有叶子节点不满足条件才会走到这 直接返回false即可
		if root == nil {
			return false
		}
		// 到达叶子节点时刚好满足条件 说明找到了路径 如果到叶子节点还没满足，那么该条路不是一个合法路径
		if root.Left == root.Right && root.Val == target {
			return true
		}
		return dfs(root.Left, target-root.Val) || dfs(root.Right, target-root.Val)
	}
	return dfs(root, targetSum)
}

package codetop4

// maxDepth 二叉树的最大深度 自底向上递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	return max(lDepth, rDepth) + 1
}

// maxDepth1 自顶向下
func maxDepth1(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		depth++
		res = max(res, depth)
		dfs(root.Left, depth)
		dfs(root.Right, depth)
	}
	dfs(root, 0)
	return res
}

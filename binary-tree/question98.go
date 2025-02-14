package binary_tree

import "math"

// isValidBST 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}
	x := node.Val
	return left < x && x < right &&
		dfs(node.Left, left, x) &&
		dfs(node.Right, x, right)
}

package binary_tree

// sumNumbers 求根节点到叶节点数字之和 自顶向下DFS
func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, x int) {
		if node == nil {
			return
		}
		x = x*10 + node.Val
		// 到达叶子节点
		if node.Left == node.Right {
			res += x
			return
		}
		dfs(node.Left, x)
		dfs(node.Right, x)
	}
	dfs(root, 0)
	return res
}

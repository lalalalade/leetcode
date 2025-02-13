package binary_tree

// rightSideView 二叉树的右视图 自顶向下DFS
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		// 先递归右子树，再递归左子树
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return res
}

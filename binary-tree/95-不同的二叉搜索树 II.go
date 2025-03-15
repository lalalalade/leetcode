package binary_tree

// generateTrees 不同的二叉搜索树 II
func generateTrees(n int) []*TreeNode {
	var dfs func(start, end int) []*TreeNode
	dfs = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		res := make([]*TreeNode, 0)
		// 枚举根结点
		for i := start; i <= end; i++ {
			// 枚举左子树根节点
			for _, x := range dfs(start, i-1) {
				// 枚举右子树根节点
				for _, y := range dfs(i+1, end) {
					root := &TreeNode{Val: i}
					root.Left = x
					root.Right = y
					res = append(res, root)
				}
			}
		}
		return res
	}
	return dfs(1, n)
}

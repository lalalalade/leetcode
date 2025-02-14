package binary_tree

// lowestCommonAncestor1 二叉搜索树的最近公共祖先
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	x := root.Val
	if p.Val < x && q.Val < x {
		return lowestCommonAncestor1(root.Left, p, q)
	}
	if p.Val > x && q.Val > x {
		return lowestCommonAncestor1(root.Right, p, q)
	}
	return root
}

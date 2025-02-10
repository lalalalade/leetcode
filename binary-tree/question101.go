package binary_tree

func isSameTree2(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree2(p.Left, q.Right) && isSameTree2(p.Right, q.Left)
}

// isSymmetric 对称二叉树
func isSymmetric(root *TreeNode) bool {
	return isSameTree2(root.Left, root.Right)
}

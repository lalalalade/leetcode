package binary_tree

// invertTree 翻转二叉树 递归法
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 交换左右子树
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// invertTree1 层序遍历 迭代法
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

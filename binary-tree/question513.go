package binary_tree

// findBottomLeftValue 找树左下角的值
func findBottomLeftValue(root *TreeNode) int {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	node := root
	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return node.Val
}

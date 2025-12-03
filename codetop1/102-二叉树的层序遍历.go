package codetop1

// levelOrder 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		path := make([]int, size)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			path[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, path)
	}
	return res
}

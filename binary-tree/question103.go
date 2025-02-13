package binary_tree

// zigzagLevelOrder 锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		vals := make([]int, size)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if len(res)%2 > 0 {
				vals[size-1-i] = node.Val
			} else {
				vals[i] = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, vals)
	}
	return res
}

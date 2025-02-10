package code_top

// zigzagLevelOrder 二叉树的锯齿形层次遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		// 本层的节点数
		vals := make([]int, n)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if len(res)%2 > 0 {
				vals[n-1-i] = node.Val
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

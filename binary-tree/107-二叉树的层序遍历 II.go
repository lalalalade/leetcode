package binary_tree

import "slices"

// levelOrderBottom 二叉树的层序遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		vals := make([]int, size)
		for i := range vals {
			node := q[0]
			q = q[1:]
			vals[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, vals)
	}
	slices.Reverse(res)
	return res
}

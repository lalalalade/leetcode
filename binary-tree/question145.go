package binary_tree

// postorderTraversal 后序遍历 递归
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		res = append(res, root.Val)
	}
	postorder(root)
	return res
}

// postorderTraversal2 三色标记法
func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*Pair, 0)
	stack = append(stack, &Pair{root, 0})

	for len(stack) > 0 {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pair.TreeNode == nil {
			continue
		}
		switch pair.int {
		case 0:
			stack = append(stack, &Pair{pair.TreeNode, 1})
			stack = append(stack, &Pair{pair.Right, 0})
			stack = append(stack, &Pair{pair.Left, 0})
		case 1:
			res = append(res, pair.Val)
		}
	}
	return res
}

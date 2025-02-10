package binary_tree

// preorderTraversal 二叉树的前序遍历 递归
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return res
}

// preorderTraversal1 迭代
func preorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return res
}

// preorderTraversal2 三色标记法
func preorderTraversal2(root *TreeNode) []int {
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
			stack = append(stack, &Pair{pair.Right, 0})
			stack = append(stack, &Pair{pair.Left, 0})
			stack = append(stack, &Pair{pair.TreeNode, 1})
		case 1:
			res = append(res, pair.Val)
		}
	}
	return res
}

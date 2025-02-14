package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal 二叉树中序遍历 递归
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

// inorderTraversal1 二叉树中序遍历 迭代
func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		// 左边界进栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 可以出栈加入res了
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		// 走到右子树 再重复上面左边界进栈的步骤
		root = root.Right
	}
	return res
}

// inorderTraversal2 二叉树中序遍历 颜色标记法
// 使用颜色标记结点的状态，新节点为白色，已访问的结点为灰色
// 如果遇到的结点为白色，则将其标记为灰色，然后将其右子节点，自身，左子节点依次入栈
// 如果遇到的结点为灰色，则将结点的值输出

type Pair struct {
	*TreeNode
	int
}

// inorderTraversal2 二叉树中序遍历 颜色标记法
func inorderTraversal2(root *TreeNode) []int {
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
			stack = append(stack, &Pair{pair.TreeNode, 1})
			stack = append(stack, &Pair{pair.Left, 0})
		case 1:
			res = append(res, pair.Val)
		}
	}
	return res
}

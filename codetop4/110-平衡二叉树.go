package codetop4

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(node.Right)
	if rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

package codetop4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, x int) {
		if node == nil {
			return
		}
		x = x*10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += x
			return
		}
		dfs(node.Left, x)
		dfs(node.Right, x)
	}
	dfs(root, 0)
	return res
}

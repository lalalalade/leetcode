package codetop2

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		// 先递归右子树，再递归左子树
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return res
}

// rightSideView1 二叉树的右视图 层序遍历
func rightSideView1(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			// 收集每层最后一个元素
			if i == size-1 {
				res = append(res, node.Val)
			}
		}
	}
	return res
}

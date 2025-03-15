package binary_tree

import (
	"math"
)

// minDepth 二叉树的最小深度 自底向上
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// minDepth1 自顶向下
func minDepth1(root *TreeNode) int {
	res := math.MaxInt
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, cnt int) {
		if root == nil {
			return
		}
		cnt++
		if root.Left == nil && root.Right == nil {
			res = min(res, cnt)
			return
		}
		dfs(root.Left, cnt)
		dfs(root.Right, cnt)
	}
	dfs(root, 0)
	if root != nil {
		return res
	}
	return 0
}

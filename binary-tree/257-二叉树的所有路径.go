package binary_tree

import "strconv"

// binaryTreePaths 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		// 把节点值（字符串形式）加入path末尾
		path += strconv.Itoa(node.Val)
		// 叶子节点 直接把path加到答案
		if node.Left == node.Right {
			res = append(res, path)
			return
		}
		path += "->"
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, "")
	return res
}

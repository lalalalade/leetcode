package binary_tree

import "slices"

// buildTree 从前序遍历与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	leftSize := slices.Index(inorder, preorder[0])
	left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
	right := buildTree(preorder[1+leftSize:], inorder[leftSize+1:])
	return &TreeNode{preorder[0], left, right}
}

// 对上面进行优化：
// 1. 用一个哈希表（或者数组）预处理inorder每个元素的下标，这样就可以在O(1)查到preorder[0]在inorder的位置
// 2. 把递归参数改成子数组下标区间的左右端点，避免复制数组
// buildTree2 优化
func buildTree2(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	index := make(map[int]int, n)
	for i, v := range inorder {
		index[v] = i
	}
	var dfs func(int, int, int, int) *TreeNode
	dfs = func(preL, preR, inL, inR int) *TreeNode {
		if preL == preR {
			return nil
		}
		leftSize := index[preorder[preL]] - inL
		left := dfs(preL+1, preL+1+leftSize, inL, inL+leftSize)
		right := dfs(preL+1+leftSize, preR, inL+1+leftSize, inR)
		return &TreeNode{preorder[preL], left, right}
	}
	return dfs(0, n, 0, n)
}

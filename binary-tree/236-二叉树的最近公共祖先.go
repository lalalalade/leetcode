package binary_tree

// lowestCommonAncestor 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 找到p或者q就可以返回了，因为如果另一个节点在更深处，说明当前节点就是最近公共祖先
	// 如果另一个节点在另一棵子树上，则没必要继续查找，直接返回
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 左子树既找不到p也找不到q，那么最近公共祖先必定在右子树上
	if left == nil {
		return right
	}
	// 同理
	if right == nil {
		return left
	}
	// 左右子树都能找到 说明root是最近公共祖先
	return root
}

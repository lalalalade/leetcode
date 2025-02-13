package binary_tree

func widthOfBinaryTree(root *TreeNode) int {
	res := 0
	if root == nil {
		return 0
	}
	m := make(map[*TreeNode]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	m[root] = 1
	for len(q) > 0 {
		size := len(q)
		// 每层最左节点的下标
		start := m[q[0]]
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			index := m[node]
			if node.Left != nil {
				q = append(q, node.Left)
				m[node.Left] = 2 * index
			}
			if node.Right != nil {
				q = append(q, node.Right)
				m[node.Right] = 2*index + 1
			}
			// 遍历到本层最后一个节点了
			if i == size-1 {
				res = max(res, index-start+1)
			}
		}
	}
	return res
}

package linked_list

func deleteNode(node *ListNode) {
	*node = *node.Next
}

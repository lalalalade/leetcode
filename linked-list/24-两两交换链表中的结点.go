package linked_list

// swapPairs 两两交换链表中的结点 迭代
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	node0 := dummy
	node1 := head
	// 至少有两个结点
	for node1 != nil && node1.Next != nil {
		node2 := node1.Next
		node3 := node2.Next

		node0.Next = node2
		node2.Next = node1
		node1.Next = node3

		node0 = node1
		node1 = node3
	}
	return dummy.Next
}

// swapPairs2 两两交换链表中的结点 递归
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node1 := head
	node2 := head.Next
	node3 := node2.Next

	node1.Next = swapPairs2(node3)
	node2.Next = node1

	return node2
}

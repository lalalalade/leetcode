package linked_list

// reorderList 重排链表
func reorderList(head *ListNode) {
	mid := middleNode(head)
	head2 := reverseList(mid)

	for head2.Next != nil {
		nxt := head.Next
		nxt2 := head2.Next
		head.Next = head2
		head2.Next = nxt
		head = nxt
		head2 = nxt2
	}
}

// middleNode 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

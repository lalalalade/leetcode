package linked_list

// removeNthFromEnd 删除链表的倒数第N个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p, q := dummy, dummy
	for ; n > 0; n-- {
		q = q.Next
	}
	// q走到最后一个结点时，p就是要删除结点的前一个结点
	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}

package linked_list

// reverseBetween 反转链表2
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy

	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}

	var pre, cur *ListNode = nil, p0.Next
	for i := 0; i < right-left+1; i++ {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	p0.Next.Next = cur
	p0.Next = pre
	return dummy.Next
}

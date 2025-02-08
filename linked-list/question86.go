package linked_list

// partition 分隔链表
func partition(head *ListNode, x int) *ListNode {
	p1 := &ListNode{}
	h1 := p1
	p2 := &ListNode{}
	h2 := p2
	cur := head
	for cur != nil {
		if cur.Val < x {
			p1.Next = &ListNode{Val: cur.Val, Next: nil}
			p1 = p1.Next
		} else {
			p2.Next = &ListNode{Val: cur.Val, Next: nil}
			p2 = p2.Next
		}
		cur = cur.Next
	}
	p1.Next = h2.Next
	return h1.Next
}

package codetop1

// reverseKGroup K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	dummy := &ListNode{Next: head}
	p0 := dummy
	var pre, cur *ListNode = nil, p0.Next
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		// 翻转后的尾结点（翻转前的首节点）
		nxt := p0.Next
		nxt.Next = cur
		// pre : 翻转后的首结点（翻转前的尾结点）
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
}

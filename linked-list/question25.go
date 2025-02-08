package linked_list

// reverseKGroup K个一组翻转链表 迭代
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 统计结点个数
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	dummy := &ListNode{Next: head}
	// p0 指向翻转组的前一个结点
	p0 := dummy
	var pre, cur *ListNode = nil, p0.Next
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		// 获取翻转组翻转后的最后一个结点
		nxt := p0.Next
		// cur 是下一组的head
		nxt.Next = cur
		// pre 是翻转组的newHead
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
}

// reverseKGroup2 K个一组翻转链表 递归
func reverseKGroup2(head *ListNode, k int) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	if n < k {
		return head
	}
	// 翻转一组链表
	var pre, cur *ListNode = nil, head
	for i := 0; i < k; i++ {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	// 继续检索下一组
	head.Next = reverseKGroup2(cur, k)
	return pre
}

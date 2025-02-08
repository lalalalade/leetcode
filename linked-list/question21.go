package linked_list

// mergeTwoLists1 合并两个有序链表 迭代
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{Val: l1.Val, Next: nil}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val, Next: nil}
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

// mergeTwoLists2 合并两个有序链表 递归
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists2(l1, l2.Next)
	return l2
}

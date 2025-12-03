package codetop2

// mergeKLists 合并k个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[0 : m/2])
	right := mergeKLists(lists[m/2:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

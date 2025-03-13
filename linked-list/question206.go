package linked_list

// reverseList 反转链表 迭代
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// reverseList1 翻转链表 递归
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

package linked_list

// sortList 排序链表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head2 := middleNode2(head)
	// 分治
	head = sortList(head)
	head2 = sortList(head2)
	return mergeTwoLists(head, head2)
}

func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head
	// 链表中间结点的前一个结点
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mid
}

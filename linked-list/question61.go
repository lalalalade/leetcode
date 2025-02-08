package linked_list

// rotateRight 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 计算链表长度
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	k %= count
	if k == 0 {
		return head
	}
	slow, fast := head, head
	for ; k > 0; k-- {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 断开重连
	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}

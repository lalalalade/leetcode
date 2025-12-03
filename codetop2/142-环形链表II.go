package codetop2

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			// 一起再往前走，相遇时慢指针刚好到达入口
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}

package codetop2

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		// 出现了重复元素 至少两个
		if cur.Next.Next.Val == val {
			// 每次循环删一个重复元素
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

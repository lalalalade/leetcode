package linked_list

// deleteDuplicates2 删除排序链表中的重复元素2 一个不留 迭代
func deleteDuplicates2(head *ListNode) *ListNode {
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

// deleteDuplicates3 删除排序链表中的重复元素2 递归
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates3(head.Next)
		return head
	}
	for head.Next != nil && head.Val == head.Next.Val {
		head = head.Next
	}
	return deleteDuplicates3(head.Next)
}

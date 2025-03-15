package linked_list

// isPalindrome 回文链表
func isPalindrome(head *ListNode) bool {
	mid := middleNode(head)
	head2 := reverseList(mid)
	for head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

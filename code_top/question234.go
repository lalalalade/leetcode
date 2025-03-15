package code_top

// isPalindrome 回文链表
func isPalindrome(head *ListNode) bool {
	mid := middleNode(head)
	newHead := reverseList(mid)
	for newHead != nil {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
	}
	return true
}

// middleNode 找到中间节点
// 若链表长度为奇数，返回中间那个
// 若链表长度为偶数，返回第二段开头那个
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

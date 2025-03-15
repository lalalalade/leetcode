package code_top

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

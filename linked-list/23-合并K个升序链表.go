package linked_list

import "container/heap"

type hp []*ListNode

func (h hp) Len() int {
	return len(h)
}

// Less 最小堆
func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}
func (h *hp) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return v
}

// mergeKLists1 合并K个升序链表
func mergeKLists1(lists []*ListNode) *ListNode {
	h := hp{}
	// 建堆
	for _, list := range lists {
		if list != nil {
			h.Push(list)
		}
	}
	heap.Init(&h)

	dummy := &ListNode{}
	cur := dummy
	for len(h) > 0 {
		// 最小元素出堆
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

// mergeTwoLists 合并两个有序链表 迭代
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

// mergeKLists2 合并k个升序链表 递归
func mergeKLists2(lists []*ListNode) *ListNode {
	m := len(lists)

	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}
	left := mergeKLists2(lists[:m/2])
	right := mergeKLists2(lists[m/2:])
	return mergeTwoLists(left, right)
}

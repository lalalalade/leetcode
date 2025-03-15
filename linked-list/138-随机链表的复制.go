package linked_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList 随机链表的复制
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 复制每个结点，把新节点直接插到原节点的后面
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
	}
	// 遍历交错链表中的原链表结点
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			// 要复制的 random 是 cur.Random 的下一个结点
			cur.Next.Random = cur.Random.Next
		}
	}
	// 分离链表
	newHead := head.Next
	cur := head
	for ; cur.Next.Next != nil; cur = cur.Next {
		clone := cur.Next
		cur.Next = clone.Next
		clone.Next = clone.Next.Next
	}
	cur.Next = nil
	return newHead
}

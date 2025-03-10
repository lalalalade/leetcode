package linked_list

type LRUCache struct {
	capacity   int
	m          map[int]*LNode
	head, tail *LNode
}

type LNode struct {
	Key       int
	Val       int
	Pre, Next *LNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &LNode{}, &LNode{}
	head.Next = tail
	tail.Next = head
	return LRUCache{
		capacity: capacity,
		m:        map[int]*LNode{},
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.m[key]; ok {
		l.moveToHead(node)
		return node.Val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.m[key]; ok {
		node.Val = value
		l.moveToHead(node)
		return
	}
	if l.capacity == len(l.m) {
		rmKey := l.removeTail()
		delete(l.m, rmKey)
	}
	newNode := &LNode{Key: key, Val: value}
	l.m[key] = newNode
	l.addToHead(newNode)
}

func (l *LRUCache) moveToHead(node *LNode) {
	l.deleteNode(node)
	l.addToHead(node)
}

func (l *LRUCache) deleteNode(node *LNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (l *LRUCache) addToHead(node *LNode) {
	l.head.Next.Pre = node
	node.Next = l.head.Next
	node.Pre = l.head
	l.head.Next = node
}

func (l *LRUCache) removeTail() int {
	node := l.tail.Pre
	l.deleteNode(node)
	return node.Key
}

package code_top

type LRUCache struct {
	capacity   int
	m          map[int]*Node
	head, tail *Node
}

type Node struct {
	Key, Value int
	Pre, Next  *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next, tail.Pre = tail, head
	return LRUCache{
		capacity: capacity,
		m:        map[int]*Node{},
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.m[key]; ok {
		l.moveToHead(node)
		return node.Value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.m[key]; ok {
		node.Value = value
		l.moveToHead(node)
		return
	}
	if len(l.m) == l.capacity {
		delete(l.m, l.removeTail())
	}
	node := &Node{Key: key, Value: value}
	l.m[key] = node
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *Node) {
	node.Pre.Next, node.Next.Pre = node.Next, node.Pre
}

func (l *LRUCache) moveToHead(node *Node) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) addToHead(node *Node) {
	node.Pre, node.Next = l.head, l.head.Next
	l.head.Next.Pre, l.head.Next = node, node
}

func (l *LRUCache) removeTail() int {
	node := l.tail.Pre
	l.removeNode(node)
	return node.Key
}

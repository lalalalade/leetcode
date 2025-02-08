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

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.moveToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.Val = value
		this.moveToHead(node)
		return
	}
	if this.capacity == len(this.m) {
		rmKey := this.removeTail()
		delete(this.m, rmKey)
	}
	newNode := &LNode{Key: key, Val: value}
	this.m[key] = newNode
	this.addToHead(newNode)
}

func (this *LRUCache) moveToHead(node *LNode) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) deleteNode(node *LNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) addToHead(node *LNode) {
	this.head.Next.Pre = node
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next = node
}

func (this *LRUCache) removeTail() int {
	node := this.tail.Pre
	this.deleteNode(node)
	return node.Key
}

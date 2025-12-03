package main

// 运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 set 。
// key 和 value 均为 string 类型
//
// 获取数据 get：
//
//	如果关键字 (key) 存在于缓存中，则获取关键字的值并返回0，否则返回-1
//
// 写入数据 set：
//
//	如果关键字已经存在，则变更其数据值
//	如果关键字不存在，则插入该组「关键字/值」
//	当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间

type Node struct {
	key, val  string
	pre, next *Node
}

type LRUCache struct {
	dummy     *Node
	keyToNode map[string]*Node
	capacity  int
}

func (l *LRUCache) Init(capacity int) *LRUCache {
	return &LRUCache{
		dummy:     &Node{},
		keyToNode: make(map[string]*Node, capacity),
		capacity:  capacity,
	}
}

func (l *LRUCache) Put(key string, val string) {
	node := &Node{val: val, key: key}
	if l.keyToNode[key] == nil {
		// 加入
		l.keyToNode[key] = node
		l.pushOnFront(node)
	}

	if len(l.keyToNode) == l.capacity {
		n := l.dummy.pre
		delete(l.keyToNode, n.key)
		l.remove(n)
	}
}

func (l *LRUCache) Get(key string) (int, string) {
	node := l.keyToNode[key]
	if node == nil {
		return -1, ""
	}
	l.remove(node)
	l.pushOnFront(node)
	return 1, node.val
}

func (l *LRUCache) remove(root *Node) {
	root.pre.next = root.next
	root.next.pre = root.pre
}

func (l *LRUCache) pushOnFront(root *Node) {
	root.pre = l.dummy
	root.next = l.dummy.next
	l.dummy.next = root
	root.next.pre = root
}

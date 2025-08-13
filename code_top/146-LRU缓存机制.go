package code_top

type LRUCache struct {
	capacity  int
	keyToNode map[int]*Node
	dummy     *Node
}

type Node struct {
	key, value int
	pre, next  *Node
}

func Constructor(capacity int) LRUCache {
	dummy := &Node{}
	dummy.pre = dummy
	dummy.next = dummy
	return LRUCache{
		capacity:  capacity,
		dummy:     dummy,
		keyToNode: map[int]*Node{},
	}
}

// 删除一个节点
func (c *LRUCache) remove(x *Node) {
	x.pre.next = x.next
	x.next.pre = x.pre
}

// 头插一个节点
func (c *LRUCache) pushFront(x *Node) {
	x.pre = c.dummy
	x.next = c.dummy.next
	x.pre.next = x
	x.next.pre = x
}

// 获取key对应的节点，同时把它移到头部
func (c *LRUCache) getNode(key int) *Node {
	node, ok := c.keyToNode[key]
	if !ok {
		return nil
	}
	c.remove(node)
	c.pushFront(node)
	return node
}

func (c *LRUCache) Get(key int) int {
	node := c.getNode(key)
	if node == nil {
		return -1
	}
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	node := c.getNode(key)
	if node != nil {
		node.value = value
		return
	}
	node = &Node{
		key:   key,
		value: value,
	}
	c.keyToNode[key] = node
	c.pushFront(node)
	if len(c.keyToNode) > c.capacity {
		backNode := c.dummy.pre
		delete(c.keyToNode, backNode.key)
		c.remove(backNode)
	}
}

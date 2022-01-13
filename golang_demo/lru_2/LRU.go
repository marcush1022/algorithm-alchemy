// package lru_2
// name LRU

package lru_2

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	keys       map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		keys:     make(map[int]*Node),
		head:     nil,
		tail:     nil,
	}
}

func (c *LRUCache) Remove(node *Node) {
	if node == c.head {
		c.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}

	if node == c.tail {
		c.tail = node.prev

		node.prev.next = nil

		node.prev = nil
		return
	}
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (c *LRUCache) AddToFront(node *Node) {
	node.prev = nil
	node.next = c.head
	if c.head != nil {
		c.head.prev = node
	}
	c.head = node
	if c.tail == nil {
		c.tail = node
		c.tail.next = nil
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.keys[key]; ok {
		// move to list front
		c.Remove(node)
		c.AddToFront(node)
		return node.val
	}
	return -1
}

func (c *LRUCache) Put(key, val int) {
	if node, ok := c.keys[key]; ok {
		node.val = val
		// move to list front
		c.Remove(node)
		c.AddToFront(node)
		return
	} else {
		node := &Node{
			key:  key,
			val:  val,
			prev: nil,
			next: nil,
		}
		c.keys[key] = node
	}
	if len(c.keys) > c.capacity {
		// remove from rear
		delete(c.keys, c.tail.key)
		c.Remove(c.tail)
	}
}

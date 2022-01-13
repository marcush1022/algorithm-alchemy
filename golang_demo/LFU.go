// package golang_demo
// name LFU

package golang_demo

import "container/list"

/*
LFU 更新和插入新页面可以发生在链表中任意位置，删除页面都发生在表尾。

LFU 同样要求查询尽量高效，O(1) 内查询。依旧选用 map 查询。修改和删除也需要 O(1) 完成，依旧选用双向链表，
继续复用 container 包中的 list 数据结构。LFU 需要记录访问次数，所以每个结点除了存储 key，value，需要再多存储 frequency 访问次数。

还有 1 个问题需要考虑，一个是如何按频次排序？相同频次，按照先后顺序排序。如果你开始考虑排序算法的话，思考方向就偏离最佳答案了。
排序至少 O(nlogn)。重新回看 LFU 的工作原理，会发现它只关心最小频次。其他频次之间的顺序并不关心。所以不需要排序。
用一个 min 变量保存最小频次，淘汰时读取这个最小值能找到要删除的结点。相同频次按照先后顺序排列，这个需求还是用双向链表实现，
双向链表插入的顺序体现了结点的先后顺序。相同频次对应一个双向链表，可能有多个相同频次，所以可能有多个双向链表。
用一个 map 维护访问频次和双向链表的对应关系。删除最小频次时，通过 min 找到最小频次，然后再这个 map 中找到这个频次对应的双向链表，
在双向链表中找到最旧的那个结点删除。这就解决了 LFU 删除操作。

LFU 的更新操作和 LRU 类似，也需要用一个 map 保存 key 和双向链表结点的映射关系。这个双向链表结点中存储的是 key-value-frequency
三个元素的元组。这样通过结点中的 key 和 frequency 可以反过来删除 map 中的 key。
*/

type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (c *LFUCache) Get(key int) int {
	value, ok := c.nodes[key]
	if !ok {
		return -1
	}

	currentNode := value.Value.(*node)
	// 从旧的 frequency list 中删除
	c.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	// 若新的 frequency 不存在则新建 list
	if _, ok := c.lists[currentNode.frequency]; !ok {
		c.lists[currentNode.frequency] = list.New()
	}

	newList := c.lists[currentNode.frequency]
	// 访问的节点加入新 list
	newNode := newList.PushFront(currentNode)
	// 加入 map
	c.nodes[key] = newNode
	// 若旧 frequency 为 min 且 list 已空
	if currentNode.frequency-1 == c.min && c.lists[currentNode.frequency-1].Len() == 0 {
		c.min++
	}
	return currentNode.value
}

func (c *LFUCache) Put(key, value int) {
	if c.capacity == 0 {
		return
	}

	// 若存在
	if currentValue, ok := c.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		// update value
		currentNode.value = value
		// 更新访问次数
		c.Get(key)
		return
	}

	// 若不存在且缓存满了
	if len(c.nodes) == c.capacity {
		minFrequencyList := c.lists[c.min]
		backNode := minFrequencyList.Back()
		// delete map
		delete(c.nodes, backNode.Value.(*node).key)
		// delete list
		minFrequencyList.Remove(backNode)
	}

	// 新节点更新 min = 1
	c.min = 1
	currentNode := &node{key: key, value: value, frequency: 1}
	// 若 freq = 1 的 list 不存在
	if _, ok := c.lists[1]; !ok {
		c.lists[1] = list.New()
	}
	newList := c.lists[1]
	// push list
	newNode := newList.PushFront(currentNode)
	// set map
	c.nodes[key] = newNode
}

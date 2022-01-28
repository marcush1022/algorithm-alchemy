// package clru
// name lru

package clru

import (
	"container/list"
	"hash/fnv"
	"sync"
)

type command int

const (
	MoveToFront command = iota
	PushFront
	Delete
)

type clear struct {
	done chan struct{}
}

type LRUCache struct {
	sync.Mutex
	cap         int
	list        *list.List
	buckets     []*bucket
	bucketMask  uint32
	deleteParis chan *list.Element
	moveParis   chan *list.Element
	control     chan interface{}
}

type Pair struct {
	key   string
	value interface{}
	cmd   command
}

func New(cap int) *LRUCache {
	c := &LRUCache{
		cap:        cap,
		list:       list.New(),
		buckets:    make([]*bucket, 1024),
		bucketMask: uint32(1024) - 1,
	}
	for i := 0; i < 1024; i++ {
		c.buckets[i] = &bucket{
			keys: make(map[string]*list.Element),
		}
	}
	c.restart()
	return c
}

func (c *LRUCache) Get(key string) interface{} {
	elem := c.getBucket(key).get(key)
	if elem == nil {
		return nil
	}
	c.move(elem)
	return elem.Value.(Pair).value
}

func (c *LRUCache) Put(key string, value interface{}) {
	elem, exist := c.getBucket(key).set(key, value)
	if exist != nil {
		c.deleteParis <- elem
	}
	c.move(elem)
}

func (c *LRUCache) Delete(key string) bool {
	elem := c.getBucket(key).delete(key)
	if elem != nil {
		c.deleteParis <- elem
		return true
	}
	return false
}

func (c *LRUCache) Clear() {
	done := make(chan struct{})
	c.control <- clear{done: done}
	// block
	<-done
}

func (c *LRUCache) Count() int {
	count := 0
	for _, b := range c.buckets {
		count += b.pairCount()
	}
	return count
}

/*
由于 LRU 逻辑的特殊性，它保证了移动结点和移除结点一定分开在双链表两端。也就是说在双链表两边同时操作，相互不影响。
双链表的临界区范围可以进一步的缩小，可以缩小到结点级。最终方案就定下来了。
用 2 个带缓冲的 channel，分别处理移动结点和删除结点，这两个 channel 可以在同一个协程中一起处理，互不影响。
*/

func (c *LRUCache) worker() {
	defer close(c.control)
	for {
		select {
		case el, ok := <-c.moveParis:
			if !ok {
				c.doDelete()
				return
			}
			if c.doMove(el) && c.list.Len() > c.cap {
				el := c.list.Back()
				c.list.Remove(el)
				keyToBeDeleted := el.Value.(Pair).key
				c.getBucket(keyToBeDeleted).delete(keyToBeDeleted)
			}
		case el := <-c.deleteParis:
			c.list.Remove(el)
		case control := <-c.control:
			switch msg := control.(type) {
			case clear:
				for _, bucket := range c.buckets {
					bucket.clear()
				}
				c.list = list.New()
				msg.done <- struct{}{}
			}
		}
	}
}

func (c *LRUCache) move(elem *list.Element) {
	select {
	case c.moveParis <- elem:
	default:
	}
}

func (c *LRUCache) doDelete() {
	for {
		select {
		case el := <-c.deleteParis:
			c.list.Remove(el)
		default:
			close(c.deleteParis)
			return
		}
	}
}

func (c *LRUCache) getBucket(key string) *bucket {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return c.buckets[hash.Sum32()&c.bucketMask]
}

func (c *LRUCache) doMove(el *list.Element) bool {
	if el.Value.(Pair).cmd == MoveToFront {
		c.list.MoveToFront(el)
		return false
	}

	elPair := el.Value.(Pair)
	newEl := c.list.PushFront(elPair)
	c.getBucket(elPair.key).update(elPair.key, newEl)
	return true
}

func (c *LRUCache) restart() {
	c.deleteParis = make(chan *list.Element, 128) // todo why is 128
	c.moveParis = make(chan *list.Element, 128)
	c.control = make(chan interface{})
	go c.worker()
}

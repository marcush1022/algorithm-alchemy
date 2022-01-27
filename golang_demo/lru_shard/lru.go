// package lru_con_safe
// name lru

package lru_shard

import (
	"container/list"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"strconv"
	"sync"
)

type (
	LRUCacheShard struct {
		lock sync.RWMutex
		Cap  int
		Keys map[int]*list.Element
		List *list.List
	}

	LRUCache struct {
		lock   sync.Mutex
		shards map[string]*LRUCacheShard
	}

	pair struct {
		K, V int
	}
)

func intToByte1(n int) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(n))
	return buf
}

func intToByte2(n int) []byte {
	return []byte(strconv.Itoa(n))
}

func (c *LRUCache) GetShard(key int) *LRUCacheShard {
	hasher := sha1.New()
	hasher.Write(intToByte2(key))
	shardKey := fmt.Sprintf("%x", hasher.Sum(nil)[0:2])
	fmt.Println(">>>>>>>>>>>>> shardKey", shardKey)
	return c.shards[shardKey]
}

func (c *LRUCache) Get(key int) int {
	shard := c.GetShard(key)

	shard.lock.RLock()
	defer shard.lock.RUnlock()
	if el, ok := shard.Keys[key]; ok {
		shard.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	shard := c.GetShard(key)

	shard.lock.Lock()
	defer shard.lock.Unlock()

	if el, ok := shard.Keys[key]; ok {
		el.Value = pair{K: key, V: value}
		shard.List.MoveToFront(el)
	} else {
		el := shard.List.PushFront(pair{K: key, V: value})
		shard.Keys[key] = el
	}

	if shard.List.Len() > shard.Cap {
		// pop back
		el := shard.List.Back()
		shard.List.Remove(el)
		delete(shard.Keys, el.Value.(pair).K)
	}
}

func Constructor(capacity int) LRUCache {
	shards := make(map[string]*LRUCacheShard, 256)
	for i := 0; i < 256; i++ {
		fmt.Println(">>>>>>>>>>>>> init", fmt.Sprintf("%02x", i))
		shards[fmt.Sprintf("%02x", i)] = &LRUCacheShard{
			Cap:  capacity,
			Keys: make(map[int]*list.Element),
			List: list.New(),
		}
	}
	return LRUCache{
		shards: shards,
	}
}

/*
不使用任何库函数，设计一个跳表。

跳表是在 O(log(n)) 时间内完成增加、删除、搜索操作的数据结构。跳表相比于树堆与红黑树，其功能与性能相当，
并且跳表的代码长度相较下更短，其设计思想与链表相似。

跳表中有很多层，每一层是一个短的链表。在第一层的作用下，增加、删除和搜索操作的时间复杂度不超过 O(n)。
跳表的每一个操作的平均时间复杂度是 O(log(n))，空间复杂度是 O(n)。

在本题中，你的设计应该要包含这些函数：

bool search(int target) : 返回target是否存在于跳表中。
void add(int num):插入一个元素到跳表。
bool erase(int num): 在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num ，删除其中任意一个即可。
了解更多 : https://en.wikipedia.org/wiki/Skip_list

注意，跳表中可能存在多个相同的值，你的代码需要处理这种情况。

样例:

Skiplist skiplist = new Skiplist();

skiplist.add(1);
skiplist.add(2);
skiplist.add(3);
skiplist.search(0);   // 返回 false
skiplist.add(4);
skiplist.search(1);   // 返回 true
skiplist.erase(0);    // 返回 false，0 不在跳表中
skiplist.erase(1);    // 返回 true
skiplist.search(1);   // 返回 false，1 已被擦除

约束条件:

0 <= num, target <= 20000
最多调用 50000 次 search, add, 以及 erase操作。
*/

package golang

import (
	"math"
	"math/rand"
)

const (
	MaxNumber       = 20000
	MinNumber       = 0
	MaxLevel        = 16      // level[0] 最大元素 50000
	MaxRandomNumber = 65535.0 // 随机层数最大值 2^16
)

// 获取随机层数
func GetRandomLevel() int {
	return MaxLevel - int(math.Log2(1.0+MaxRandomNumber*rand.Float64()))
}

type skipListNode struct {
	value int           // 节点值
	right *skipListNode // 右指针
	down  *skipListNode // 下指针
}

type SkipList struct {
	head *skipListNode
}

func Constructor() SkipList {
	leftFrame := make([]*skipListNode, MaxLevel)
	rightFrame := make([]*skipListNode, MaxLevel)

	// 初始化节点值
	for i := 0; i < MaxLevel; i++ {
		leftFrame[i] = &skipListNode{
			value: MinNumber - 1,
			right: nil,
			down:  nil,
		}

		rightFrame[i] = &skipListNode{
			value: MaxNumber + 1,
			right: nil,
			down:  nil,
		}
	}

	// 初始化节点指针
	for i := MaxLevel - 2; i >= 0; i-- {
		leftFrame[i].right = rightFrame[i]
		leftFrame[i].down = leftFrame[i+1]
		rightFrame[i].down = rightFrame[i+1]
	}
	leftFrame[MaxLevel-1].right = rightFrame[MaxLevel-1]
	return SkipList{head: leftFrame[0]}
}

// 查找 target 是否存在
func (sl *SkipList) Search(target int) bool {
	node := sl.head
	for node != nil {
		if node.right.value == target {
			return true
		}

		if node.right.value > target {
			node = node.down
		} else {
			node = node.right
		}
	}
	return false
}

// 插入一个元素
func (sl *SkipList) Add(value int) {
	// jumpedNode 跳过的点
	jumpedNodes := make([]*skipListNode, MaxLevel)
	node := sl.head
	index := 0
	for node != nil {
		if node.right.value >= value {
			jumpedNodes[index] = node
			node = node.down
			index++
		} else {
			node = node.right
		}
	}

	// 获取该元素的随机层数
	level := GetRandomLevel()
	tmpArr := make([]*skipListNode, level)
	tmpNode := &skipListNode{
		value: -1,
		right: nil,
		down:  nil,
	}
	for i, a := range tmpArr {
		frontPoint := jumpedNodes[MaxLevel-level+i]
		a = &skipListNode{
			value: value,
			right: frontPoint.right,
			down:  nil,
		}
		frontPoint.right = a
		tmpNode.down = a
		tmpNode = a
	}
}

// 删除一个元素
func (sl *SkipList) Erase(value int) (res bool) {
	node := sl.head
	for node != nil {
		if node.right.value > value {
			node = node.down
		} else if node.right.value < value {
			node = node.right
		} else {
			res = true
			node.right = node.right.right
			// 将整层元素删除
			node = node.down
		}
	}
	return
}

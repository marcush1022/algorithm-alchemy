package golang

import (
	"math"
	"math/rand"
)

/*
概率分层直接进行了一个随机数的对数计算，其他的跟传统跳表实现起来大同小异。

用的是纯跳表单元，只有值、向右、向下三个属性，没有使用数组存表，add 的时候用一个 prev 数组记录前一个跳跃点，
所以不需要在结构体里用超过两个方向的形式存储，空间上比数组存表略优一点吧。

对于python来说，用统计字典collections.defaultdict(int)或者桶数组[0] * 20000来提交这题当然很快，大约100ms+这样，
这些数据结构的底层都是封装在c扩展里的，纯python写的跳表比前者慢这么多很正常，我这个代码也不例外，大约300ms+这样，
最多不超过500ms，其他python代码超时肯定是哪里写错了，大概率是跳跃的时候没有跳出循环。

golang的实现，线上测试不管做不做弊时间都差不多，说明跳表相对传统映射数据结构还是有一定实现价值的。

不管的python还是golang，都可以对重复值进行优化，即让skipNonde多存一个count之类的属性，不过这样还得再增加不少代码，
暂时就算了。
*/

const (
	maxLevel = 16      // 5
	maxRand  = 65535.0 // 32
)

func randLevel() int {
	return maxLevel - int(math.Log2(1.0+maxRand*rand.Float64()))
}

type skipNode struct {
	value int       // 值
	right *skipNode // 向右
	down  *skipNode // 向下
}

type Skiplist struct {
	head *skipNode
}

func constructor() Skiplist {
	left := make([]*skipNode, maxLevel)
	right := make([]*skipNode, maxLevel)
	for i := 0; i < maxLevel; i++ {
		left[i] = &skipNode{-1, nil, nil}
		right[i] = &skipNode{20001, nil, nil}
	}
	for i := maxLevel - 2; i >= 0; i-- {
		left[i].right = right[i]
		left[i].down = left[i+1]
		right[i].down = right[i+1]
	}
	left[maxLevel-1].right = right[maxLevel-1]
	return Skiplist{left[0]}
}

func (this *Skiplist) Search(target int) bool {
	node := this.head
	for node != nil {
		if node.right.value > target {
			node = node.down
		} else if node.right.value < target {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	// prev 记录前一个跳跃点
	prev := make([]*skipNode, maxLevel)
	i := 0
	node := this.head
	for node != nil {
		if node.right.value >= num {
			prev[i] = node
			i++
			node = node.down
		} else {
			node = node.right
		}
	}
	n := randLevel()
	arr := make([]*skipNode, n)
	t := &skipNode{-1, nil, nil}
	for i, a := range arr {
		p := prev[maxLevel-n+i]
		a = &skipNode{num, p.right, nil}
		p.right = a
		t.down = a
		t = a
	}
}

func (this *Skiplist) Erase(num int) (ans bool) {
	node := this.head
	for node != nil {
		if node.right.value > num {
			node = node.down
		} else if node.right.value < num {
			node = node.right
		} else {
			ans = true
			node.right = node.right.right
			node = node.down
		}
	}
	return
}

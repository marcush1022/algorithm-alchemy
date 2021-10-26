// package golang_demo
// name dfs

package golang_demo

import (
	"math"
)

var MinSum = math.MaxInt64

type Node2 struct {
	name string
	follows []string
	weight int
}

func resetMinSum() {
	MinSum = math.MaxInt64
}

func popBack(slice []string) []string {
	// pop back and deep copy
	slice = slice[:len(slice)-1]
	tmp := make([]string, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}

func dfsGetPath(head string, nodes []Node2, nodeHasFollow map[string]struct{}, sum, index int, results *[][]string, path []string) {
	//if index == len(nodes) {
	//	return
	//}
	//
	//// if it has no follow nodes 无后续节点，即端点
	//followNodes, ok := nodeHasFollow[nodes[index].name]
	//
	//if !ok {
	//	// if got a solution
	//	if sum < MinSum {
	//		fmt.Println(">>>>>>>> path, sum, name", path, sum, nodes[index].name)
	//		*results = append(*results, path)
	//		MinSum = sum
	//	}
	//	return
	//}
	//
	//// has follow nodes
	//for i := 0; i < len(followNodes); i++ {
	//	sum += followNodes[i].weight
	//	path = append(path, followNodes[i].name)
	//	dfsGetPath(head, followNodes, nodeHasFollow, sum, i, results, path)
	//	// clean up
	//	sum -= followNodes[i].weight
	//	path = popBack(path)
	//}
}

func GetShortPath(nodes []Node2, nodeHasFollow map[string]struct{}) [][]string {
	results := &[][]string{}
	path := make([]string, 0)
	*results = append(*results, path)
	for head, _ := range nodeHasFollow {
		dfsGetPath(head, nodes, nodeHasFollow, 0, 0, results, path)
		resetMinSum()
	}
	return *results
}

// package golang_demo
// name dfs

package golang_demo

import (
	"fmt"
	"math"
)

var MinSum = math.MaxInt64

type Node2 struct {
	name    string
	follows []string
	weight  int
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

func dfsGetPath(node string, nodes map[string][]string, result *[][]string, path []string) {
	follows, _ := nodes[node]
	// if this node has no follow
	if len(follows) == 0 {
		// valid path
		*result = append(*result, path)
		fmt.Println(">>>>> result", result)
		return
	}

	for i := 0; i < len(follows); i++ {
		path = append(path, follows[i])
		dfsGetPath(follows[i], nodes, result, path)
		path = popBack(path)
	}
}

func GetPath3(nodes map[string][]string) [][]string {
	result := make([][]string, 0)
	path := make([]string, 0)
	for node, _ := range nodes {
		dfsGetPath(node, nodes, &result, path)
	}
	return result
}
// package golang_demo
// name dfs

package golang_demo

import (
	"fmt"
	"math"
)

var MinSum uint = math.MaxUint32

type Node2 struct {
	name    string
	follows []string
	weight  int
}

func resetMinSum() {
	MinSum = math.MaxUint32
}

func popBack(slice []string) []string {
	// pop back and deep copy
	slice = slice[:len(slice)-1]
	tmp := make([]string, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}

func dfsGetPath(node string, nodes map[string][]string, result *[][]string, path []string, sum uint, weights map[string]uint) {
	follows, _ := nodes[node]
	// if this node has no follow
	if len(follows) == 0 {
		// valid path
		if sum < MinSum {
			// found a solution
			*result = append(*result, path)
			fmt.Println(">>>>> result", result)
			MinSum = sum
		}
		return
	}

	for i := 0; i < len(follows); i++ {
		path = append(path, follows[i])
		sum += weights[follows[i]]
		dfsGetPath(follows[i], nodes, result, path, sum, weights)
		// do clean up
		path = popBack(path)
		sum -= weights[follows[i]]
	}
}

func GetPath3(nodes map[string][]string, weights map[string]uint) [][]string {
	result := make([][]string, 0)
	path := make([]string, 0)
	for node, _ := range nodes {
		path = append(path, node)
		dfsGetPath(node, nodes, &result, path, 0, weights)
		// reset min sum each node
		resetMinSum()
		path = make([]string, 0)
	}
	return result
}
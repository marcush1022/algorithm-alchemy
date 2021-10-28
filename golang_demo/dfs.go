// package golang_demo
// name dfs

package golang_demo

import (
	"fmt"
	"math"
	"strings"
)

var MinSum uint

type Node2 struct {
	name    string
	follows []string
	weight  int
}

func intiMinSum() {
	MinSum = math.MaxUint32
}

func popBack(slice []string) []string {
	// pop back and deep copy
	slice = slice[:len(slice)-1]
	tmp := make([]string, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}

func dfsGetPath(node string, nodes map[string][]string, path []string, shortPath map[string]string, sum uint, weights map[string]uint) {
	follows, _ := nodes[node]
	// if this node has no follow
	if len(follows) == 0 {
		// valid path
		if sum < MinSum {
			// found a solution
			shortPath[fmt.Sprintf("%v,%v", path[0], path[len(path)-1])] = strings.Join(path, ",")
			MinSum = sum
		}
		return
	}

	for i := 0; i < len(follows); i++ {
		path = append(path, follows[i])
		sum += weights[follows[i]]
		dfsGetPath(follows[i], nodes, path, shortPath, sum, weights)
		// clean up
		path = popBack(path)
		sum -= weights[follows[i]]
	}
}

func GetPath3(nodes map[string][]string, weights map[string]uint) {
	path := make([]string, 0)
	shortPath := make(map[string]string)
	intiMinSum()
	for node, _ := range nodes {
		path = append(path, node)
		dfsGetPath(node, nodes, path,  shortPath, 0, weights)
		fmt.Printf("after %v, shortPath = %v\n", node, shortPath)
		// reset min sum each node
		intiMinSum()
		path = make([]string, 0)
	}

	for k, v := range shortPath {
		fmt.Printf("start end = %v, shorted path = %v\n", k, v)
	}
	return
}
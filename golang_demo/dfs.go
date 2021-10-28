// package golang_demo
// name dfs

package golang_demo

import (
	"fmt"
	"math"
	"strings"
)

var MinSum uint

// dfsGetPath
// node: 节点名, nodes 节点依赖图 (serviceInvocations), path 当前路径, shortPath 同起点终点的最短路径 (结果集), sum 路径和, weights 节点权重
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

/*
思路：
	nodes 即 serviceInvocations
	1. 遍历 nodes 中每个节点 n
	2. 若 n 有后序节点 follows，继续 dfs 遍历每个 follows 节点
	3. 若 follows 节点依然有后序节点，继续遍历
	4. 重复此步骤 3，直至节点无 follows 节点，表明此节点为终点节点
	5. 比较路径 sum 与全局最小 MinSum，若小，则加入结果集 shortPath
	6. 此趟搜索结束，重置 MinSum，并返回步骤 1
*/

// GetPath3 nodes 节点依赖图 (serviceInvocations), weights 节点权重
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

func intiMinSum() {
	MinSum = math.MaxUint32
}

// popBack pop back and deep copy
func popBack(slice []string) []string {
	slice = slice[:len(slice)-1]
	tmp := make([]string, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}
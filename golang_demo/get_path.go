// package golang_demo
// name min_path

package golang_demo

import (
	"fmt"
	"math"
)

const MAXVEX int = 6
const MAX_WEIGHT int = math.MaxInt16

type Path struct {
	SumWeight   int
	PassedNodes string
}

//var shortTablePath = [MAXVEX]int{MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT}
var shortTablePath = [MAXVEX]Path{
	{SumWeight: MAX_WEIGHT, PassedNodes: ""},
	{SumWeight: MAX_WEIGHT, PassedNodes: ""},
	{SumWeight: MAX_WEIGHT, PassedNodes: ""},
	{SumWeight: MAX_WEIGHT, PassedNodes: ""},
	{SumWeight: MAX_WEIGHT, PassedNodes: ""},
	{SumWeight: MAX_WEIGHT, PassedNodes: ""},
}

func GetPath() {
	graph := NewGraph2()

	var TablePathMin int       //存放shortTablePath中,未遍历的最小结点的值
	var Vx int                 //存放shortTablePath中,未遍历的最小结点的下标
	var isgetPath [MAXVEX]bool //记录结点是否已经找到v0到vx的最小路径

	// 获取v0这一行的权值数组
	for v := 0; v < len(graph); v++ {
		shortTablePath[v].SumWeight = graph[0][v]
	}
	shortTablePath[0].SumWeight = 0
	isgetPath[0] = true

	//遍历v1 ~ v8
	for v := 1; v < len(graph); v++ {
		TablePathMin = MAX_WEIGHT

		//找出shortTablePath中,未遍历的最小结点的值
		for w := 0; w < len(graph); w++ {
			if !isgetPath[w] && shortTablePath[w].SumWeight < TablePathMin {
				Vx = w
				TablePathMin = shortTablePath[w].SumWeight
			}
		}
		isgetPath[Vx] = true
		for j := 0; j < len(graph); j++ {
			if !isgetPath[j] && TablePathMin+graph[Vx][j] < shortTablePath[j].SumWeight {
				shortTablePath[j].SumWeight = TablePathMin + graph[Vx][j]
				fmt.Println("vx j", Vx, j)
				// shortTablePath[v].PassedNodes += fmt.Sprintf("%v,", j)
			}
		}
		fmt.Println("遍历完V", v, "后:", shortTablePath)
	}

	//输出
	for i := 0; i < len(shortTablePath); i++ {
		fmt.Println("V0到V", i, "最小路径:", shortTablePath[i])
	}
}

func NewGraph2() [MAXVEX][MAXVEX]int {
	var graph [MAXVEX][MAXVEX]int
	var v0a = [MAXVEX]int{0, 20, 30, 40, MAX_WEIGHT, MAX_WEIGHT}
	var v1b = [MAXVEX]int{MAX_WEIGHT, 0, MAX_WEIGHT, 40, 15, MAX_WEIGHT}
	var v2c = [MAXVEX]int{MAX_WEIGHT, MAX_WEIGHT, 0, MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT}
	var v3d = [MAXVEX]int{MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT, 0, MAX_WEIGHT, MAX_WEIGHT}
	var v4e = [MAXVEX]int{MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT, 0, 20}
	var v5f = [MAXVEX]int{MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT, MAX_WEIGHT, 0}
	graph[0] = v0a
	graph[1] = v1b
	graph[2] = v2c
	graph[3] = v3d
	graph[4] = v4e
	graph[5] = v5f
	return graph
}

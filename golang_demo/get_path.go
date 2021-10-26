// package golang_demo
// name min_path

package golang_demo

import (
"fmt"
)

const MAXVEX int = 6
const MAXWEIGHT int = 1000

//var shortTablePath = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
var shortTablePath = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}

var NodeWeights = map[int]int{
	0:10,
	1: 20,
	2: 30,
	3: 40,
	4: 15,
	5: 20,
}


func GetPath() {
	graph := NewGraph2()

	var TablePathMin int       //存放shortTablePath中,未遍历的最小结点的值
	var Vx int                 //存放shortTablePath中,未遍历的最小结点的下标
	var isgetPath [MAXVEX]bool //记录结点是否已经找到v0到vx的最小路径

	// 获取v0这一行的权值数组
	for v := 0; v < len(graph); v++ {
		shortTablePath[v] = graph[0][v]
	}
	shortTablePath[0] = 0
	isgetPath[0] = true

	//遍历v1 ~ v8
	for v := 1; v < len(graph); v++ {
		TablePathMin = MAXWEIGHT

		//找出shortTablePath中,未遍历的最小结点的值
		for w := 0; w < len(graph); w++ {
			if !isgetPath[w] && shortTablePath[w] < TablePathMin {
				Vx = w
				TablePathMin = shortTablePath[w]
			}
		}
		isgetPath[Vx] = true
		for j := 0; j < len(graph); j++ {
			if !isgetPath[j] && TablePathMin+graph[Vx][j]+NodeWeights[j] < shortTablePath[j] {
				shortTablePath[j] = TablePathMin + graph[Vx][j]+NodeWeights[j]
			}
		}

		fmt.Println("遍历完V", v, "后:", shortTablePath)

	}

	//输出
	for i := 0; i < len(shortTablePath); i++ {
		fmt.Println("V0到V", i, "最小路径:", shortTablePath[i])
	}

}

//func NewGraph() [MAXVEX][MAXVEX]int {
//	var graph [MAXVEX][MAXVEX]int
//	var v0 = [MAXVEX]int{0, 1, 5, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
//	var v1 = [MAXVEX]int{1, 0, 3, 7, 5, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
//	var v2 = [MAXVEX]int{5, 3, 0, MAXWEIGHT, 1, 7, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
//	var v3 = [MAXVEX]int{MAXWEIGHT, 7, MAXWEIGHT, 0, 2, MAXWEIGHT, 3, MAXWEIGHT, MAXWEIGHT}
//	var v4 = [MAXVEX]int{MAXWEIGHT, 5, 1, 2, 0, 3, 6, 9, MAXWEIGHT}
//	var v5 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, 7, MAXWEIGHT, 3, 0, MAXWEIGHT, 5, MAXWEIGHT}
//	var v6 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 3, 6, MAXWEIGHT, 0, 2, 7}
//	var v7 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 9, 5, 2, 0, 4}
//	var v8 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 7, 4, 0}
//	graph[0] = v0
//	graph[1] = v1
//	graph[2] = v2
//	graph[3] = v3
//	graph[4] = v4
//	graph[5] = v5
//	graph[6] = v6
//	graph[7] = v7
//	graph[8] = v8
//	return graph
//}

func NewGraph2() [MAXVEX][MAXVEX]int {
	var graph [MAXVEX][MAXVEX]int
	var v0a = [MAXVEX]int{0, 0, 0, 0, MAXWEIGHT, MAXWEIGHT}
	var v1b = [MAXVEX]int{MAXWEIGHT, 0, MAXWEIGHT, 0, 0, MAXWEIGHT}
	var v2c = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, 0, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
	var v3d = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 0, MAXWEIGHT, MAXWEIGHT}
	var v4e = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 0, 0}
	var v5f = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 0}
	graph[0] = v0a
	graph[1] = v1b
	graph[2] = v2c
	graph[3] = v3d
	graph[4] = v4e
	graph[5] = v5f
	return graph
}

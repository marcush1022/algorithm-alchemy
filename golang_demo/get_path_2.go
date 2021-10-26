// package golang_demo
// name get_path_2

package golang_demo

import (
	"fmt"
	"math"
)

const VertexNum = 6 // V

var Dist [VertexNum][VertexNum]int
var MaxInt = math.MaxInt16 // 32767

var charMap = map[int]string{
	0: "a",
	1: "b",
	2: "c",
	3: "d",
	4: "e",
	5: "f",
}

func GetPath2() {
	initDist(MaxInt)
	// for k from 0 to V
	for k := 0; k < VertexNum; k++ {
		// for i from 0 to V
		for i := 0; i < VertexNum; i++ {
			// for j from 0 to V
			for j := 0; j < VertexNum; j++ {
				if i == j {
					continue
				}
				// if cost dist[i][j] > dist[i][k] + dist[k][j]
				if Dist[i][j] > Dist[i][k]+Dist[k][j] {
					Dist[i][j] = Dist[i][k] + Dist[k][j]
				}
			}
		}
	}
	for _, v := range Dist {
		fmt.Println(v)
	}
}

func initDist(max int) {
	Dist[0] = [VertexNum]int{0, 20, 30, 40, max, max}
	Dist[1] = [VertexNum]int{max, 0, max, 40, 15, max}
	Dist[2] = [VertexNum]int{max, max, 0, max, max, max}
	Dist[3] = [VertexNum]int{max, max, max, 0, max, max}
	Dist[4] = [VertexNum]int{max, max, max, max, 0, 20}
	Dist[5] = [VertexNum]int{max, max, max, max, max, 0}
}

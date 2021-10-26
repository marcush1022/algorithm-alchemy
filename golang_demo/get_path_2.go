// package golang_demo
// name get_path_2

package golang_demo

import (
	"errors"
	"fmt"
	"math"
)

var StdErrParam = errors.New("params error")
var MaxInt = math.MaxInt16 // 32767

var charMap = map[int]string{
	0: "a",
	1: "b",
	2: "c",
	3: "d",
	4: "e",
	5: "f",
}

var numMap = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
}

func GetPath2(start, end string, serviceInvocations map[string][]string, invocationsTime map[string]int) ([]string, error) {
	vertexNum := len(invocationsTime)
	if vertexNum == 0 {
		return nil, StdErrParam
	}
	dist, err := InitDist(serviceInvocations, invocationsTime, vertexNum)
	if err != nil {
		fmt.Println(err)
	}

	// pre is the last node between i -> j
	pre := InitPre(vertexNum)

	// for k from 0 to V
	for k := 0; k < vertexNum; k++ {
		// for i from 0 to V
		for i := 0; i < vertexNum; i++ {
			// for j from 0 to V
			for j := 0; j < vertexNum; j++ {
				if i == j {
					continue
				}
				// if cost dist[i][j] > dist[i][k] + dist[k][j]
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					pre[i][j] = pre[k][j]
				}
			}
		}
	}
	for _, v := range dist {
		fmt.Println(v)
	}

	//fmt.Println(getPath(convert2Num(start), convert2Num(end)))
	return getPath(convert2Num(start), convert2Num(end), pre, vertexNum), nil
}

func InitDist(serviceInvocations map[string][]string, invocationsTime map[string]int, vertexNum int) ([][]int, error) {
	dist := make([][]int, vertexNum)
	for i := range dist {
		dist[i] = make([]int, vertexNum)
	}

	for i := 0; i < vertexNum; i++ {
		for j := 0; j < vertexNum; j++ {
			if i == j {
				continue
			}
			dist[i][j] = MaxInt
		}
	}

	for node, links := range serviceInvocations {
		for _, l := range links {
			dist[convert2Num(node)][convert2Num(l)] = 0
		}
	}

	for i := 0; i < vertexNum; i++ {
		for j := 0; j < vertexNum; j++ {
			if i != j && dist[i][j] == 0 {
				dist[i][j] = invocationsTime[convert2Char(j)]
			}
		}
	}

	for _, v := range dist {
		fmt.Println(v)
	}
	return dist, nil
}

func InitPre(vertexNum int) [][]int {
	pre := make([][]int, vertexNum)
	for i := range pre {
		pre[i] = make([]int, vertexNum)
	}

	for i := 0; i < vertexNum; i++ {
		for j := 0; j < vertexNum; j++ {
			pre[i][j] = i
		}
	}
	return pre
}

// convert2Char convert a number to a letter
func convert2Char(n int) string {
	return charMap[n]
}

// convert2Num convert a letter to number
func convert2Num(c string) int {
	return numMap[c]
}

// getPath get path from pre in reverse order
func getPath(start, end int, pre [][]int, vertexNum int) []string {
	res := make([]string, 0, vertexNum)
	index := end
	for index != start {
		// fmt.Print(convert2Char(index), ",")
		res = append(res, convert2Char(index))
		index = pre[start][index]
	}
	// fmt.Println(convert2Char(index))
	res = append(res, convert2Char(index))
	reverseSlice(res)
	return res
}

func reverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//func initDist(max int) {
//	Dist[0] = [VertexNum]int{0, 20, 30, 40, max, max}
//	Dist[1] = [VertexNum]int{max, 0, max, 40, 15, max}
//	Dist[2] = [VertexNum]int{max, max, 0, max, max, max}
//	Dist[3] = [VertexNum]int{max, max, max, 0, max, max}
//	Dist[4] = [VertexNum]int{max, max, max, max, 0, 20}
//	Dist[5] = [VertexNum]int{max, max, max, max, max, 0}
//
//	for i := 0; i < VertexNum; i++ {
//		for j := 0; j < VertexNum; j++ {
//			Pre[i][j]= i
//		}
//	}
//}

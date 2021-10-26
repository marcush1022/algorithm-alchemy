// package golang_demo
// name get_path_2

package golang_demo

import (
	"fmt"
	"math"
)

const VertexNum = 6 // V

var Dist [VertexNum][VertexNum]int
// Pre is the last node between i -> j
var Pre [VertexNum][VertexNum]int
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

func GetPath2(start, end string) {
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
					Pre[i][j] = Pre[k][j]
				}
			}
		}
	}
	for _, v := range Dist {
		fmt.Println(v)
	}

	fmt.Println(getPath(convert2Num(start), convert2Num(end)))
}

func initDist(max int) {
	Dist[0] = [VertexNum]int{0, 20, 30, 40, max, max}
	Dist[1] = [VertexNum]int{max, 0, max, 40, 15, max}
	Dist[2] = [VertexNum]int{max, max, 0, max, max, max}
	Dist[3] = [VertexNum]int{max, max, max, 0, max, max}
	Dist[4] = [VertexNum]int{max, max, max, max, 0, 20}
	Dist[5] = [VertexNum]int{max, max, max, max, max, 0}

	for i := 0; i < VertexNum; i++ {
		for j := 0; j < VertexNum; j++ {
			Pre[i][j]= i
		}
	}
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
func getPath(start, end int) []string {
	res := make([]string, 0, VertexNum)
	index := end
	for index != start {
		// fmt.Print(convert2Char(index), ",")
		res = append(res, convert2Char(index))
		index= Pre[start][index]
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

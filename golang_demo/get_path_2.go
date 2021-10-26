// package golang_demo
// name get_path_2

package golang_demo

import "fmt"

//FMap is used to store point`s distrance
var FMap [4][4]int32

//PointNum is point num
var PointNum int = 4

func GetPath2() {
	var Inf int32
	Inf = 1 << 16
	initMap(Inf)
	//中间点
	for k := 0; k < PointNum; k++ {
		//任意两点间距离更新
		for i := 0; i < PointNum; i++ {
			for j := 0; j < PointNum; j++ {
				if i == j {
					continue
				}
				if FMap[i][j] > FMap[i][k]+FMap[k][j] {
					FMap[i][j] = FMap[i][k] + FMap[k][j]
				}
			}
		}
	}
	fmt.Println(FMap)
}

//init map
func initMap2(inf int32) {
	FMap[0][0] = 0
	FMap[0][1] = 2
	FMap[0][2] = 6
	FMap[0][3] = 4
	FMap[1][0] = inf
	FMap[1][1] = 0
	FMap[1][2] = 3
	FMap[1][3] = inf
	FMap[2][0] = 7
	FMap[2][1] = inf
	FMap[2][2] = 0
	FMap[2][3] = 1
	FMap[3][0] = 5
	FMap[3][1] = inf
	FMap[3][2] = 12
	FMap[3][3] = 0
}

//init map
func initMap(inf int32) {
	FMap[0][0] = 0
	FMap[0][1] = 2
	FMap[0][2] = 6
	FMap[0][3] = 4
	FMap[1][0] = inf
	FMap[1][1] = 0
	FMap[1][2] = 3
	FMap[1][3] = inf
	FMap[2][0] = 7
	FMap[2][1] = inf
	FMap[2][2] = 0
	FMap[2][3] = 1
	FMap[3][0] = 5
	FMap[3][1] = inf
	FMap[3][2] = 12
	FMap[3][3] = 0
}

package golang

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func GetRandomLevel2(maxLevel int, maxRand float64) int {
	return maxLevel - int(math.Log2(1.0+maxRand*rand.Float64()))
}

func TestGetRandomLevel(t *testing.T) {

	i := 0
	rand.Seed(time.Now().Unix())
	for i < 10 {
		res := GetRandomLevel2(5, 32)
		fmt.Println(">>>>>>>>", res)
		i++
	}
}

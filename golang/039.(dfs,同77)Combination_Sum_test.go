// package golang
// name 039.(dfs,同77)Combination_Sum_test

package golang

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	assert.Equal(t, "", compareCombinations(CombinationSum(7, []int{2, 3, 6, 7}), [][]int{
		{7},
		{2, 2, 3},
	}))

	assert.Equal(t, "", compareCombinations(CombinationSum(7, []int{7}), [][]int{
		{7},
	}))

	assert.Equal(t, "", compareCombinations(CombinationSum(7, []int{}), [][]int{}))
}

func compareCombinations(act, exp [][]int) string {
	fmt.Println("act: ", act)
	mapA := make(map[string]struct{})
	mapE := make(map[string]struct{})

	for _, v := range act {
		mapA[fmt.Sprint(v)] = struct{}{}
	}

	for _, v := range exp {
		mapE[fmt.Sprint(v)] = struct{}{}
	}

	for kA, vA := range mapA {
		if vE, ok := mapE[kA]; !ok {
			return fmt.Sprintf("k %v not in mapE", kA)
		} else if vE != vA {
			return fmt.Sprintf("vE %v != vA %v", vE, vA)
		}
	}

	for kE, vE := range mapE {
		if vA, ok := mapA[kE]; !ok {
			return fmt.Sprintf("k %v not in mapA", kE)
		} else if vE != vA {
			return fmt.Sprintf("vE %v != vA %v", vE, vA)
		}
	}
	return ""
}

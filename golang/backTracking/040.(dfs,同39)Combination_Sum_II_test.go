// package golang
// name 040.(dfs,同39)Combination_Sum_II_test

package backTracking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	assert.Equal(t, "", compareCombinations2(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8), [][]int{
		{1, 7},
		{1, 2, 5},
		{2, 6},
		{1, 1, 6},
	}))
}

func compareCombinations2(act, exp [][]int) string {
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

// package golang
// name 047.(dfs,剪枝)Permutations_II_test

package backTracking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutations2(t *testing.T) {
	assert.Equal(t, "", comparePermutations2(permuteUnique([]int{1, 1, 2}), [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}))

	assert.Equal(t, "", comparePermutations2(permuteUnique([]int{1, 1}), [][]int{
		{1, 1},
	}))

	assert.Equal(t, "", comparePermutations2(permuteUnique([]int{1}), [][]int{
		{1},
	}))

	assert.Equal(t, "", comparePermutations2(permuteUnique([]int{0, 1, 0, 0, 9}), [][]int{
		{0, 0, 0, 1, 9}, {0, 0, 0, 9, 1}, {0, 0, 1, 0, 9}, {0, 0, 1, 9, 0}, {0, 0, 9, 0, 1},
		{0, 0, 9, 1, 0}, {0, 1, 0, 0, 9}, {0, 1, 0, 9, 0}, {0, 1, 9, 0, 0}, {0, 9, 0, 0, 1},
		{0, 9, 0, 1, 0}, {0, 9, 1, 0, 0}, {1, 0, 0, 0, 9}, {1, 0, 0, 9, 0}, {1, 0, 9, 0, 0},
		{1, 9, 0, 0, 0}, {9, 0, 0, 0, 1}, {9, 0, 0, 1, 0}, {9, 0, 1, 0, 0}, {9, 1, 0, 0, 0},
	}))

	assert.Equal(t, "", comparePermutations2(permuteUnique([]int{0, 0, 0, 9}), [][]int{
		{0, 0, 0, 9},
		{0, 0, 9, 0},
		{0, 9, 0, 0},
		{9, 0, 0, 0},
	}))
}

func comparePermutations2(act, exp [][]int) string {
	if len(act) != len(exp) {
		return fmt.Sprintf("len(act) %v != len(exp) %v", len(act), len(exp))
	}

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

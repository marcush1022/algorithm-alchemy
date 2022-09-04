// package golang
// name 078.(dfs)Subsets_test

package backTracking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subsets(t *testing.T) {
	assert.Equal(t, "", compareSubsets(subsets([]int{1, 2, 3}),
		[][]int{
			{3},
			{1},
			{2},
			{1, 2, 3},
			{1, 3},
			{2, 3},
			{1, 2},
			{},
		},
	))

	assert.Equal(t, "", compareSubsets(subsets([]int{1, 2}),
		[][]int{
			{1},
			{2},
			{1, 2},
			{},
		},
	))

	assert.Equal(t, "", compareSubsets(subsets([]int{3}),
		[][]int{
			{3},
			{},
		},
	))

	assert.Equal(t, "", compareSubsets(subsets([]int{}),
		[][]int{
			{},
		},
	))
}

func compareSubsets(act, exp [][]int) string {
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

// package golang
// name 491.(dfs)Increasing_Subsequences_test

package golang

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrSub(t *testing.T) {
	assert.Equal(t, "", compareIncrSubSubSets([][]int{
		{4, 6},
		{4, 7},
		{4, 6, 7},
		{4, 6, 7, 7},
		{6, 7},
		{6, 7, 7},
		{7, 7},
		{4, 7, 7},
	}, IncrSub([]int{4, 6, 7, 7})))

	assert.Equal(t, "", compareIncrSubSubSets([][]int{
		{4, 4},
	}, IncrSub([]int{4, 4, 3, 2, 1})))
}

func compareIncrSubSubSets(act, exp [][]int) string {
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

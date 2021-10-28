// package golang_demo
// name dfs_test

package golang_demo

import (
	"fmt"
	"testing"
)

func TestGetShortPath(t *testing.T) {
	nodes := []Node2{
		{
			name:    "a",
			follows: []string{"b", "c", "d"},
			weight:  10,
		},
		{
			name:    "b",
			follows: []string{"d", "e"},
			weight:  20,
		},
		{
			name:    "c",
			follows: []string{},
			weight:  30,
		},
		{
			name:    "d",
			follows: []string{},
			weight:  40,
		},
		{
			name:    "e",
			follows: []string{"f"},
			weight:  15,
		},
		{
			name:    "f",
			follows: []string{},
			weight:  20,
		},
	}

	m := make(map[string][]string)
	for _, v := range nodes {
		if len(v.follows) > 0 {
			m[v.name] = v.follows
		}
	}

	weights := map[string]uint{
		"a": 10, "b": 20, "c": 30,
		"d": 40, "e": 15, "f": 20,
	}

	res := GetPath3(m, weights)
	fmt.Println(res)
}

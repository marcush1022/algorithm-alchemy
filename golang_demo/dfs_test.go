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
			name:   "a",
			follows:  []string{"b", "c", "d"},
			weight: 10,
		},
		{
			name:   "b",
			follows:  []string{"d", "e"},
			weight: 20,
		},
		{
			name:   "c",
			follows:  []string{},
			weight: 30,
		},
		{
			name:   "d",
			follows:  []string{},
			weight: 40,
		},
		{
			name:   "e",
			follows:  []string{"f"},
			weight: 15,
		},
		{
			name:   "f",
			follows:  []string{},
			weight: 20,
		},
	}

	m := make(map[string]struct{})
	for _, v := range nodes {
		if len(v.follows) > 0 {
			m[v.name] = struct{}{}
		}
	}
	res := GetShortPath(nodes, m)
	fmt.Println(res)
}

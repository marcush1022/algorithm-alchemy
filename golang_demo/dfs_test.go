// package golang_demo
// name dfs_test

package golang_demo

import (
	"testing"
)

func TestGetShortPath(t *testing.T) {
	nodes := map[string][]string{
		"a": {"b", "c", "d"},
		"b": {"d", "e"},
		//"c": {},
		"d": {"e"},
		"e": {"f"},
		//"f": {},
	}

	weights := map[string]uint{
		"a": 10, "b": 20, "c": 30,
		"d": 40, "e": 15, "f": 20,
	}

	GetPath3(nodes, weights)
}

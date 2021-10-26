// package golang_demo
// name get_path_test

package golang_demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPath(t *testing.T) {
	GetPath()
}

func TestGetPath2(t *testing.T) {
	var serviceInvocations = map[string][]string{
		"a": []string{"b", "c", "d"},
		"b": []string{"d", "e"},
		"c": []string{},
		"d": []string{},
		"e": []string{"f"},
		"f": []string{},
	}

	var invocationsTime = map[string]int{
		"a": 10, "b": 20, "c": 30, "d": 40, "e": 15, "f": 20,
	}
	res, err := GetPath2("a", "f", serviceInvocations, invocationsTime)
	assert.NoError(t, err, "err should be nil")
	fmt.Println(">>>>>> res:", res)
}

func TestInitDist(t *testing.T) {
	var serviceInvocations = map[string][]string{
		"a": []string{"b", "c", "d"},
		"b": []string{"d", "e"},
		"c": []string{},
		"d": []string{},
		"e": []string{"f"},
		"f": []string{},
	}

	var invocationsTime = map[string]int{
		"a": 10, "b": 20, "c": 30, "d": 40, "e": 15, "f": 20,
	}

	InitDist(serviceInvocations, invocationsTime, len(invocationsTime))
}

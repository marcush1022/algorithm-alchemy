// package golang
// name 022.(dfs)Generate_Parentheses_test

package golang

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	assert.Equal(t, "", compareParentheses(GenerateParentheses(3), []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}))

	assert.Equal(t, "", compareParentheses(GenerateParentheses(1), []string{
		"()",
	}))
}

func TestGenerateParentheses2(t *testing.T) {
	assert.Equal(t, "", compareParentheses(GenerateParentheses2(3), []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}))

	assert.Equal(t, "", compareParentheses(GenerateParentheses2(1), []string{
		"()",
	}))
}

func compareParentheses(act, exp []string) string {
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

// package golang
// name 151.(字符串处理,medium)Reverse Words in a String_test

package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// cba dc
// cd abc

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "cd abc", ReverseWords("abc cd"))
	assert.Equal(t, "   ", ReverseWords("   "))
	assert.Equal(t, "a", ReverseWords("a"))
	assert.Equal(t, "abc", ReverseWords("abc"))
	assert.Equal(t, "ab c", ReverseWords("c ab"))
	assert.Equal(t, "blue is sky the", ReverseWords("the sky is blue"))
}

// package golang
// name 008.(字符串处理)String_to_Integer(atoi)_test

package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, 4193, MyAtoi("4193 with words"), "4193 with words")
	assert.Equal(t, -2147483648, MyAtoi("-91283472332"), "-91283472332")
	assert.Equal(t, 42, MyAtoi("42"), "42")
	assert.Equal(t, 0, MyAtoi(""), "")
	assert.Equal(t, 0, MyAtoi(" "), " ")
	assert.Equal(t, -42, MyAtoi("   -42"), "-42")
	assert.Equal(t, 0, MyAtoi("words and 987"), "words and 987")
	assert.Equal(t, 12345678, MyAtoi("  0000000000012345678"), "  0000000000012345678")
}

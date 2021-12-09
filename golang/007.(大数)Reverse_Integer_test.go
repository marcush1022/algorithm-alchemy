// package golang
// name 007.(大数)Reverse_Integer_test

package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	assert.Equal(t, 321, ReverseInteger(123), 123)
	assert.Equal(t, -321, ReverseInteger(-123), -123)
	assert.Equal(t, 0, ReverseInteger(1534236469), 1534236469)
}

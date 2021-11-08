// package golang
// name 009.(大数)Palindrome_Number_test

package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, IsPalindrome(121))
	assert.Equal(t, true, IsPalindrome(1))
	assert.Equal(t, true, IsPalindrome(12321))
	assert.Equal(t, false, IsPalindrome(1212))
}

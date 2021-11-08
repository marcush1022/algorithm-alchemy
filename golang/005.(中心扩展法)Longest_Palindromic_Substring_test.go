// package golang
// name 005.(中心扩展法)Longest_Palindromic_Substring_test

package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "a", LongestPalindrome("a"), "a")
	assert.Equal(t, "aba", LongestPalindrome("caba"), "caba")
	assert.Equal(t, "aba", LongestPalindrome("cabad"), "cabad")
	assert.Equal(t, "ccabacc", LongestPalindrome("ccabacc"), "ccabacc")
	assert.Equal(t, "aaa", LongestPalindrome("abcaaabc"), "abcaaabc")
	assert.Equal(t, "bab", LongestPalindrome("babad"), "babad")
	assert.Equal(t, "bb", LongestPalindrome("cbbd"), "cbbd")
}

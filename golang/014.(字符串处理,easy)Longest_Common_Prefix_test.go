// package golang
// name 014.(字符串处理,easy)Longest_Common_Prefix_test

package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", LongestCommonPrefix([]string{"flower", "flow", "flight"}), []string{"flower", "flow", "flight"})
	assert.Equal(t, "", LongestCommonPrefix([]string{"dog", "racecar", "car"}), []string{"dog", "racecar", "car"})
}

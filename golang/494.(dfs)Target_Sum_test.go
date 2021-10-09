// package golang
// name 494.(dfs)Target_Sum_test

package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTargetSum(t *testing.T) {
	assert.Equal(t, 5, TargetSum([]int{1, 1, 1, 1, 1}, 3))
	assert.Equal(t, 0, TargetSum([]int{1, 1, 1, 1, 1}, 10))
	assert.Equal(t, 1, TargetSum([]int{1}, 1))
}

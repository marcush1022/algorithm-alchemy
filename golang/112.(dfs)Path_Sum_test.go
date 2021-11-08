// package golang
// name 112.(dfs)Path_Sum_test

package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathSum(t *testing.T) {
	node5 := TreeNode{Val: 5}
	node4_1 := TreeNode{Val: 4}
	node8 := TreeNode{Val: 8}
	node11 := TreeNode{Val: 11}
	node13 := TreeNode{Val: 13}
	node4_2 := TreeNode{Val: 4}
	node7 := TreeNode{Val: 7}
	node2 := TreeNode{Val: 2}
	node1 := TreeNode{Val: 1}

	node5.Left = &node4_1
	node5.Right = &node8
	node4_1.Left = &node11
	node11.Left = &node7
	node11.Right = &node2
	node8.Left = &node13
	node8.Right = &node4_2
	node4_2.Right = &node1

	assert.Equal(t, true, PathSum(&node5, 22))
}

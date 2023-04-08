// package twoPointers
// dir_path golang/twoPointers
// name 283.(Easy)Move_Zeroes_test

package twoPointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	assert.Equal(t, fmt.Sprint([]int{1, 3, 12, 0, 0}), fmt.Sprint(nums))

	nums = []int{0}
	moveZeroes(nums)
	assert.Equal(t, fmt.Sprint([]int{0}), fmt.Sprint(nums))
}

// package twoPointers
// dir_path golang/twoPointers
// name 167.(Medium)Two_Sum_II_Input_array_is_sorted_test

package twoPointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum2(t *testing.T) {
	assert.Equal(t, fmt.Sprint([]int{1, 2}), fmt.Sprint(twoSum2([]int{2, 7, 11, 15}, 9)))
	assert.Equal(t, fmt.Sprint([]int{1, 3}), fmt.Sprint(twoSum2([]int{2, 3, 4}, 6)))
	assert.Equal(t, fmt.Sprint([]int{1, 2}), fmt.Sprint(twoSum2([]int{-1, 0}, -1)))
}

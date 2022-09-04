// package twoPointers
// name 26.(Easy)Remove_Duplicates_from_Sorted_Array_test

package twoPointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	assert.Equal(t, "", checker26([]int{1, 1, 2}, []int{1, 2}, 2))
	assert.Equal(t, "", checker26([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5))
}

func checker26(nums, expNums []int, k int) string {
	res := removeDuplicates(nums)
	if res != k {
		return fmt.Sprintf("res != k, res: %v, k: %v", res, k)
	}

	for i := 0; i < k; i++ {
		if nums[i] != expNums[i] {
			return fmt.Sprintf("nums[i] != expNums[i], i: %v, nums[i]: %v, expNums[i]: %v", i, nums[i], expNums[i])
		}
	}
	return ""
}

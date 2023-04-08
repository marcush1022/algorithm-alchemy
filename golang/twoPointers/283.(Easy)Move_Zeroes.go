// package twoPointers
// dir_path golang/twoPointers
// name 283.(Easy)Move Zeroes

package twoPointers

import "fmt"

/**
283. Move Zeroes

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]


Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1
*/

func moveZeroes(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	slowIndex := 0
	for fastIndex := 0; fastIndex < length; fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	// nums[0]..nums[slow-1] is unique elements
	fmt.Println(">>>>>>> slowIndex", slowIndex)
	for ; slowIndex < length; slowIndex++ {
		nums[slowIndex] = 0
	}
}

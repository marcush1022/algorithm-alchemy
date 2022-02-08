// package golang
// name 015.3_Sum

package golang

import (
	"sort"
)

/*

题目 #
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets
in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:


Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	result := make([][]int, 0)
	left, right := 0, 0
	target := 0

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		left = i + 1
		right = len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// skip left duplicate numbers
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// skip right duplicate numbers
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}

		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}

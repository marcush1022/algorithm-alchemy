// package golang
// name 090.(dfs同78)Subsets.II.go

/**
Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:

Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
Example 2:

Input: nums = [0]
Output: [[],[0]]


Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10

https://labuladong.github.io/algo/1/9/
如果一个节点有多条值相同的树枝相邻，则只遍历第一条，剩下的都剪掉，不要去遍历
需要先进行排序，让相同的元素靠在一起，如果发现 nums[i] == nums[i-1]，则跳过
*/

package backTracking

import "sort"

func popBack90(slice []int) []int {
	slice = slice[:len(slice)-1]
	tmp := make([]int, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}

func dfs2(index int, result *[][]int, path, numbers []int) {
	if index == len(numbers) {
		return
	}

	for i := index; i < len(numbers); i++ {
		path = append(path, numbers[i])
		*result = append(*result, path)
		dfs2(i+1, result, path, numbers)
		path = popBack90(path)
		for i < len(numbers)-1 && numbers[i] == numbers[i+1] {
			i++
		}
	}
}

func GetSubsets2(numbers []int) [][]int {
	results := &[][]int{}
	path := make([]int, 0)
	*results = append(*results, path)
	sort.Ints(numbers)
	dfs2(0, results, path, numbers)
	return *results
}

func backtrack90(result *[][]int, track, numbers []int, start int) {
	// leaf node
	if start == len(numbers) {
		return
	}

	for i := start; i < len(numbers); i++ {
		track = append(track, numbers[i])
		*result = append(*result, track)
		backtrack90(result, track, numbers, i+1)
		track = popBack90(track)
		for i < len(numbers)-1 && numbers[i] == numbers[i+1] {
			i++
		}
	}
}

func subsetsWithDup(nums []int) [][]int {
	results := &[][]int{}
	track := make([]int, 0)
	*results = append(*results, []int{})
	sort.Ints(nums)
	backtrack90(results, track, nums, 0)
	return *results
}

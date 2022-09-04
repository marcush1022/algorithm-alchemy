// package golang
// name 078.(dfs)Subsets

/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]


Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.

https://labuladong.github.io/algo/1/9/

注意这棵树的特性：
如果把根节点作为第 0 层，将每个节点和根节点之间树枝上的元素作为该节点的值，那么第 n 层的所有节点就是大小为 n 的所有子集。

那么再进一步，如果想计算所有子集，那只要遍历这棵多叉树，把所有节点的值收集起来不就行了？

使用一个变量来记录路径，每遍历到一个元素即表示找到一条路径，将其加入子集中。
例如数组[1,2,3]
从1开始递归查询2,3，对于2，继续向下搜索，搜索完后将2删除。
*/

package backTracking

import (
	"fmt"
	"sort"
)

func popBack78(slice []int) []int {
	// pop back and deep copy
	slice = slice[:len(slice)-1]
	tmp := make([]int, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}

func backtrack78(result *[][]int, track []int, numbers []int, startIndex int) {
	// leaf node
	if startIndex == len(numbers) {
		return
	}

	// nums.length is tree's depth
	for i := startIndex; i < len(numbers); i++ {
		track = append(track, numbers[i])
		// each node of tree is subset
		*result = append(*result, track)
		fmt.Println(">>>>> result", result)
		// start+1
		backtrack78(result, track, numbers, i+1)
		// reset last node
		track = popBack78(track)
	}
}

func subsets(numbers []int) [][]int {
	results := &[][]int{}
	path := make([]int, 0)
	*results = append(*results, path)
	sort.Ints(numbers)
	backtrack78(results, path, numbers, 0)
	return *results
}

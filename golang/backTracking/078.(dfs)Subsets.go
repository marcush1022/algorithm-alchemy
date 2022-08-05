// package golang
// name 078.(dfs)Subsets

/*
Given a set of distinct integers, nums, return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,3], a solution is:

[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

使用一个变量来记录路径，每遍历到一个元素即表示找到一条路径，将其加入子集中。
例如数组[1,2,3]
从1开始递归查询2,3，对于2，继续向下搜索，搜索完后将2删除。
*/

package backTracking

import (
	"fmt"
	"sort"
)

func popBack(slice []int) []int {
	// pop back and deep copy
	slice = slice[:len(slice)-1]
	tmp := make([]int, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}

func dfs(result *[][]int, path []int, numbers []int, index int) {
	if index == len(numbers) {
		return
	}
	for i := index; i < len(numbers); i++ {
		path = append(path, numbers[i])
		*result = append(*result, path)
		fmt.Println(">>>>> result", result)
		dfs(result, path, numbers, i+1)
		path = popBack(path)
	}
}

func GetSubsets(numbers []int) [][]int {
	results := &[][]int{}
	path := make([]int, 0)
	*results = append(*results, path)
	sort.Ints(numbers)
	dfs(results, path, numbers, 0)
	return *results
}

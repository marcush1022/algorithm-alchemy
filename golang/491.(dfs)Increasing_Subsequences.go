// package golang
// name 491.(dfs)Increasing_Subsequences

package golang

import (
	"fmt"
)

/*
Given an integer array, your task is to find all the different possible increasing
subsequences of the given array, and the length of an increasing subsequence should be at least 2 .

Example:
Input: [4, 6, 7, 7]
Output: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]

Note: 典型dfs，set用来去重
*/

func popBackIncrSub(original []int) []int {
	// pop back and deep copy
	original = original[:len(original)-1]
	tmp := make([]int, len(original), len(original))
	copy(tmp, original)
	return tmp
}

func dfsIncrSub(results *[][]int, set map[string]struct{}, numbers, path []int, index int) {
	if len(path) > 1 {
		fmt.Println(">>>>>>> path", path)
		key := fmt.Sprint(path)
		if _, ok := set[key]; !ok {
			// this path is already deep copied no need copy again
			*results = append(*results, path)
			set[key] = struct{}{}
		}
	}

	for i := index; i < len(numbers); i++ {
		if len(path) == 0 || numbers[i] >= path[len(path)-1] {
			if len(path) != 0 {
				fmt.Println(">>>>>> numbers[i]", numbers[i], ",", path[len(path)-1], ",", path)
			}
			path = append(path, numbers[i])
		}
		dfsIncrSub(results, set, numbers, path, i+1)
		// pop back
		path = popBackIncrSub(path)
	}
}

func IncrSub(numbers []int) [][]int {
	results := make([][]int, 0)
	set := make(map[string]struct{})
	path := make([]int, 0)
	dfsIncrSub(&results, set, numbers, path, 0)
	fmt.Println(">>>>>>>>>>>> results", results)
	return results
}

// package golang
// name 040.(dfs,同39)Combination_Sum II

package backTracking

import (
	"fmt"
	"sort"
)

/**
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates
where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.



Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]


Constraints:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

func popBack40(original []int) []int {
	original = original[:len(original)-1]
	tmp := make([]int, len(original), len(original))
	copy(tmp, original)
	return tmp
}

func dfsCombination2(result *[][]int, index, sum, target int, path, numbers []int) {
	if sum > target {
		return
	}

	if sum == target {
		// valid path
		fmt.Println(">>>>>>>>> path", path)
		*result = append(*result, path)
		return
	}

	for i := index; i < len(numbers); i++ {
		if i > index && numbers[i] == numbers[i-1] {
			// disallow same number in same depth, each number in [numbers] is a depth
			continue
		}
		sum += numbers[i]
		path = append(path, numbers[i])
		dfsCombination2(result, i+1, sum, target, path, numbers)
		// clean up
		sum -= numbers[i]
		path = popBack40(path)
	}
}

func CombinationSum2(target int, numbers []int) [][]int {
	results := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(numbers)
	dfsCombination2(&results, 0, 0, target, path, numbers)
	fmt.Println(">>>>>>>>>>> results", results)
	return results
}

func backtrack40(result *[][]int, track, numbers []int, trackSum, target int, startIndex int) {
	if trackSum > target {
		return
	}

	if trackSum == target {
		*result = append(*result, track)
		return
	}

	for i := startIndex; i < len(numbers); i++ {
		trackSum += numbers[i]
		track = append(track, numbers[i])
		// go to next level
		backtrack40(result, track, numbers, trackSum, target, i+1)
		// reset last node
		track = popBack40(track)
		trackSum -= numbers[i]
		for i < len(numbers)-1 && numbers[i] == numbers[i+1] {
			i++
		}
	}
}

func combinationSum2(numbers []int, target int) [][]int {
	results := &[][]int{}
	track := make([]int, 0)
	sort.Ints(numbers)
	backtrack40(results, track, numbers, 0, target, 0)
	return *results
}

// package golang
// name 047.(dfs,剪枝)Permutations_II

package golang

import (
	"fmt"
	"sort"
)

/*
Given a collection of numbers that might contain duplicates, return all possible unique
permutations.

For example,
[1,1,2] have the following unique permutations:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

func swapPermutations2(slice []int, one, two int) {
	slice[one], slice[two] = slice[two], slice[one]
}

func dfsPermutations2(results *[][]int, numbers []int, index int) {
	if index == len(numbers) {
		// valid path
		path := make([]int, len(numbers), len(numbers))
		copy(path, numbers)
		*results = append(*results, path)
		fmt.Println(">>>>>>> path", path)
		return
	}

	for i := index; i < len(numbers); i++ {
		if i > index && numbers[i] == numbers[i-1] {
			// disallow same numbers in same depth
			continue
		}
		swapPermutations2(numbers, i, index)
		dfsPermutations2(results, numbers, index+1)
		// clean up
		swapPermutations2(numbers, i, index)
	}
}

func Permutations2(numbers []int) [][]int {
	results := make([][]int, 0)
	sort.Ints(numbers)
	dfsPermutations2(&results, numbers, 0)
	fmt.Println(">>>>>>>>>> results", results)
	return results
}

// package golang
// name 040.(dfs,同39)Combination_Sum II

package backTracking

import (
	"fmt"
	"sort"
)

func popBackCombination2(original []int) []int {
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
		path = popBackCombination2(path)
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

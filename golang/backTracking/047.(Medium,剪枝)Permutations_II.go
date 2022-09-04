// package golang
// name 047.(dfs,剪枝)Permutations_II

package backTracking

import (
	"fmt"
	"sort"
)

/**
Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.



Example 1:

Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]
Example 2:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


Constraints:

1 <= nums.length <= 8
-10 <= nums[i] <= 10

https://labuladong.github.io/algo/1/9/

标准全排列算法之所以出现重复，是因为把相同元素形成的排列序列视为不同的序列，但实际上它们应该是相同的；
而如果固定相同元素形成的序列顺序，当然就避免了重复

当出现重复元素时，比如输入 nums = [1,2,2',2'']，2' 只有在 2 已经被使用的情况下才会被选择，
同理，2'' 只有在 2' 已经被使用的情况下才会被选择，这就保证了相同元素在排列中的相对位置保证固定
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

func popBack47(slice []int) []int {
	// pop back and deep copy
	slice = slice[:len(slice)-1]
	tmp := make([]int, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}

func backtrack47(results *[][]int, track []int, numbers []int, used map[int]bool) {
	if len(track) == len(numbers) {
		tmp := make([]int, len(track), len(track))
		copy(tmp, track)
		*results = append(*results, tmp)
		return
	}

	for i := 0; i < len(numbers); i++ {
		if used[i] {
			continue
		}

		if i > 0 && numbers[i] == numbers[i-1] && !used[i-1] {
			continue
		}

		// use
		used[i] = true
		track = append(track, numbers[i])
		backtrack47(results, track, numbers, used)
		// reset last node
		track = popBack47(track)
		used[i] = false
	}
}

func permuteUnique(numbers []int) [][]int {
	results := make([][]int, 0)
	track := make([]int, 0)
	used := make(map[int]bool)
	sort.Ints(numbers)
	backtrack47(&results, track, numbers, used)
	return results
}

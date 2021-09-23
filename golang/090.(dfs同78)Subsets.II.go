// package golang
// name 090.(dfs同78)Subsets.II.go

package golang

import "sort"

func popBack2(slice []int) []int {
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
		path = popBack2(path)
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

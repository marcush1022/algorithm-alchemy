// package golang
// name 046.(dfs)Permutations

package backTracking

import "fmt"

/*
Given a collection of distinct numbers, return all possible permutations.

For example,
[1,2,3] have the following permutations:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

用index记录已确定的位置，每次确定一个元素的最终位置，当所有元素位置都确定时(index=nums.size()-1)输出
第一趟所有元素都和自己交换，所以输出原序列

注意：第一个字符与后面的某个位置的字符发生交换后，需要再次发生交换，不然顺序就会被打乱。
举个例子，在字符串abc中，在把第一个字符看成是a，后面的字符b、c看成一个整体的时候，
abc这个相对顺序不能改变，所以当b与c发生交换变成了acb之后，需要再次交换两个字符，重新回到abc。
*/

/*
path 1:
	original: [1, 2, 3]
	i = 0, index = 0, 0 swp 0: [1, 2, 3]
	i = 1, index = 1, 1 swp 1: [1, 2, 3]
	i = 2, index = 2, 2 swp 2: [1, 2, 3]
	index = 3, return

path 2:
	original: [1, 2, 3]
	i = 2, index = 1, 2 swp 1: [1, 3, 2]
	i = 2, index = 2, 2 swp 2: [1, 3, 2]
	index = 3, return
*/

func swapPermutations(slice []int, one, two int) {
	slice[one], slice[two] = slice[two], slice[one]
}

func dfsPermutations(results *[][]int, numbers []int, index int) {
	if index == len(numbers) {
		// valid path
		tmp := make([]int, len(numbers), len(numbers))
		copy(tmp, numbers)
		*results = append(*results, tmp)
		fmt.Println(">>>>>>> numbers", numbers)
		return
	}

	// numbers' length is tree depth
	// for search path in each depth
	for i := index; i < len(numbers); i++ {
		// swap index with i
		swapPermutations(numbers, i, index)
		// go to next depth
		dfsPermutations(results, numbers, index+1)
		// clean up
		swapPermutations(numbers, index, i)
	}
}

func Permutations(numbers []int) [][]int {
	results := make([][]int, 0)
	dfsPermutations(&results, numbers, 0)
	fmt.Println(">>>>>>>> results", results)
	return results
}

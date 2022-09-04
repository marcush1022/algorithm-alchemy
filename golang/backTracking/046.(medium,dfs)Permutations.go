// package golang
// name 046.(dfs)Permutations

package backTracking

import "fmt"

/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]


Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.

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

方法2：
https://labuladong.github.io/algo/1/9/

刚才讲的组合/子集问题使用 start 变量保证元素 nums[start] 之后只会出现 nums[start+1..] 中的元素，
通过固定元素的相对位置保证不出现重复的子集。

但排列问题本身就是让你穷举元素的位置，nums[i] 之后也可以出现 nums[i] 左边的元素，所以之前的那一套玩不转了，
需要额外使用 used 数组来标记哪些元素还可以被选择

我们用 used 数组标记已经在路径上的元素避免重复选择，然后收集所有叶子节点上的值，就是所有全排列的结果
*/

func swap46(slice []int, one, two int) {
	slice[one], slice[two] = slice[two], slice[one]
}

func dfsPermutations(results *[][]int, numbers []int, startIndex int) {
	if startIndex == len(numbers) {
		// valid path
		tmp := make([]int, len(numbers), len(numbers))
		copy(tmp, numbers)
		*results = append(*results, tmp)
		fmt.Println(">>>>>>> numbers", numbers)
		return
	}

	// numbers' length is tree depth
	// for search path in each depth
	for i := startIndex; i < len(numbers); i++ {
		// swap index with i
		swap46(numbers, i, startIndex)
		// go to next depth
		dfsPermutations(results, numbers, startIndex+1)
		// clean up
		swap46(numbers, startIndex, i)
	}
}

func Permutations(numbers []int) [][]int {
	results := make([][]int, 0)
	dfsPermutations(&results, numbers, 0)
	fmt.Println(">>>>>>>> results", results)
	return results
}

func popBack46(slice []int) []int {
	// pop back and deep copy
	slice = slice[:len(slice)-1]
	tmp := make([]int, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}

func backtrack46(results *[][]int, track []int, numbers []int, used map[int]bool) {
	// leaf node
	if len(track) == len(numbers) {
		tmp := make([]int, len(track), len(track))
		copy(tmp, track)
		*results = append(*results, tmp)
		return
	}

	for i := 0; i < len(numbers); i++ {
		// if used
		if v := used[numbers[i]]; v {
			continue
		}
		// use
		used[numbers[i]] = true
		track = append(track, numbers[i])
		backtrack46(results, track, numbers, used)
		// reset last node
		track = popBack46(track)
		used[numbers[i]] = false
	}
}

func permute(numbers []int) [][]int {
	results := make([][]int, 0)
	track := make([]int, 0)
	used := make(map[int]bool)
	backtrack46(&results, track, numbers, used)
	return results
}

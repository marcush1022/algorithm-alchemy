// package backTracking
// name 077.(dfs)Combinations

/**
Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].

You may return the answer in any order.



Example 1:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
Example 2:

Input: n = 1, k = 1
Output: [[1]]


Constraints:

1 <= n <= 20
1 <= k <= n

https://labuladong.github.io/algo/1/9/

还是以 nums = [1,2,3] 为例，刚才让你求所有子集，就是把所有节点的值都收集起来；
现在你只需要把第 2 层（根节点视为第 0 层）的节点收集起来，就是大小为 2 的所有组合
*/

package backTracking

import "fmt"

func popBack77(slice []int) []int {
	// pop back and deep copy
	slice = slice[:len(slice)-1]
	tmp := make([]int, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}

func backtrack77(results *[][]int, track []int, maxNumber, numberLimit, startIndex int) {
	// stop at 2 nodes level
	if len(track) == numberLimit {
		// valid path
		tmp := make([]int, len(track), len(track))
		copy(tmp, track)
		*results = append(*results, tmp)
		fmt.Println(">>>>>", track)
		return
	}

	// maxNumber is tree's depth
	// for search path in each depth
	for i := startIndex; i <= maxNumber; i++ {
		track = append(track, i)
		// start+1
		backtrack77(results, track, maxNumber, numberLimit, i+1)
		// reset last node
		track = popBack77(track)
	}
}

func combine(n int, k int) [][]int {
	results := &[][]int{}
	track := make([]int, 0)
	backtrack77(results, track, n, k, 1)
	return *results
}

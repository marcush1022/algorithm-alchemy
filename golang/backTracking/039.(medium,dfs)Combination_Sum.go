// package golang
// name 039.(dfs,同77)Combination_Sum

package backTracking

import (
	"fmt"
	"sort"
)

/*
Given an array of distinct integers candidates and a target integer target,
return a list of all unique combinations of candidates where the chosen numbers sum to target.
You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.



Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1
Output: []


Constraints:

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
All elements of candidates are distinct.
1 <= target <= 500

dfs模板级题目...
取数字和为给定值，可以重复，求有多少种组合

Note: 回溯算法递归算法框架：
int a[n];
 try(int i)
 {
    if(i>n)
        输出结果;
      else
     {
        for(j = 下界; j <= 上界; j=j+1)  // 枚举i所有可能的路径
        {
            if(fun(j))                 // 满足限界函数和约束条件
              {
                 a[i] = j;
               ...                         // 其他操作
                 try(i+1);
               回溯前的清理工作（如a[i]置空值等）;
               }
          }
      }
 }

这道题说是组合问题，实际上也是子集问题：candidates 的哪些子集的和为 target？
想解决这种类型的问题，也得回到回溯树上，我们不妨先思考思考，标准的子集/组合问题是如何保证不重复使用元素的？
答案在于 backtrack 递归时输入的参数 start
这个 i 从 start 开始，那么下一层回溯树就是从 start + 1 开始，从而保证 nums[start] 这个元素不会被重复使用
那么反过来，如果我想让每个元素被重复使用，我只要把 i + 1 改成 i 即可
这相当于给之前的回溯树添加了一条树枝，在遍历这棵树的过程中，一个元素可以被无限次使用
*/

func popBack39(original []int) []int {
	original = original[:len(original)-1]
	tmp := make([]int, len(original), len(original))
	copy(tmp, original)
	return tmp
}

func dfsSum(result *[][]int, target, sum, index int, numbers, path []int) {
	if sum > target {
		return
	}

	if sum == target {
		// valid path
		fmt.Println(">>>>>>>> path", path)
		*result = append(*result, path)
		return
	}

	for i := index; i < len(numbers); i++ {
		sum += numbers[i]
		path = append(path, numbers[i])
		dfsSum(result, target, sum, i, numbers, path)
		sum -= numbers[i]
		path = popBack40(path)
	}
}

func backtrack39(result *[][]int, track, candidates []int, trackSum, target, startIndex int) {
	if trackSum > target {
		return
	}

	if trackSum == target {
		*result = append(*result, track)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		trackSum+= candidates[i]
		track = append(track, candidates[i])
		backtrack39(result, track, candidates, trackSum, target, i)
		// reset last node
		track = popBack39(track)
		trackSum -= candidates[i]
	}
}

func combinationSum(candidates []int, target int) [][]int {
    results := &[][]int{}
	track := make([]int, 0)
	sort.Ints(candidates)
	backtrack39(results, track, candidates, 0, target, 0)
	return *results
}

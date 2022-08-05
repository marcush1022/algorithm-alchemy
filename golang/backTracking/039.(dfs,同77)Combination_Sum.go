// package golang
// name 039.(dfs,同77)Combination_Sum

package backTracking

import "fmt"

/*
Given a set of candidate numbers (C) (without duplicates) and a target number (T), find all unique
combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7,
A solution set is:
[
  [7],
  [2, 2, 3]
]

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
*/

func popBackSum(slice []int) []int {
	slice = slice[:len(slice)-1]
	tmp := make([]int, len(slice), len(slice))
	copy(tmp, slice)
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
		path = popBackSum(path)
	}
}

func CombinationSum(n int, numbers []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	dfsSum(&result, n, 0, 0, numbers, path)
	fmt.Println(">>>>>>>> result", result)
	return result
}

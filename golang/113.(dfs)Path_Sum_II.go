// package golang
// name 113.(dfs)Path_Sum_II

package golang

/*
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

For example:
Given the below binary tree and sum = 22,
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
return
[
   [5,4,11,2],
   [5,8,4,5]
]

与前一题同样的思路
*/

func dfsPathSum2(root *TreeNode, target int, path []int, result *[][]int)

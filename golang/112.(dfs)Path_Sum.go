// package golang
// name 112.(dfs)Path_Sum

package golang

/*
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding
up all the values along the path equals the given sum.

For example:
Given the below binary tree and sum = 22,
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
*/

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func dfsPathSum(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target && root.Left == nil && root.Right == nil {
		return true
	}

	return dfsPathSum(root.Left, target-root.Val) || dfsPathSum(root.Right, target-root.Val)
}

func PathSum(root *TreeNode, target int) bool {
	return dfsPathSum(root, target)
}

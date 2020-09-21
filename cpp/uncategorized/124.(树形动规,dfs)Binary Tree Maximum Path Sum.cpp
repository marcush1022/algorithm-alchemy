/*
Given a binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in 
the tree along the parent-child connections. The path must contain at least one node and does not need
to go through the root.

For example:
Given the below binary tree,

       1
      / \
     2   3
Return 6.

最大path sum指的是当前结点的值加左右到叶结点的路径和的最大值，用maxSum记录全局最大值
例如：[1,2,3,4,5,6]最大path sum为17

*/

class Solution {
public:
    int maxSum=0;  //maxSum记录全局最大值
    
    int dfs(TreeNode *root)
    {
        if(root==NULL)
            return 0;

        int leftSum=dfs(root->left);
        int rightSum=dfs(root->right);
        
        leftSum=max(leftSum, 0);   //因为题目中路径的定义不一定从叶子开始
        rightSum=max(rightSum, 0);
        
        int sum=leftSum+rightSum+root->val;
        maxSum=max(maxSum, sum);
        
        return root->val+max(leftSum, rightSum);  //递归调用返回左右路径和的较大值
    }
    
    int maxPathSum(TreeNode* root) {
        maxSum=root->val;
        dfs(root);
        return maxSum;
    }
};

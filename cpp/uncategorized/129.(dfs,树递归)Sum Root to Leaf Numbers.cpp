/*
Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path 1->2->3 which represents the number 123.

Find the total sum of all root-to-leaf numbers.

For example,

    1
   / \
  2   3
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.

Return the sum = 12 + 13 = 25.
*/

class Solution {
public:
    int sum=0;
    
    void dfs(TreeNode *root, int tmp)
    {
        if(root==NULL)
            return ;
        if(root->left==NULL && root->right==NULL)
            sum+=(tmp*10+root->val);
        if(root->left)
            dfs(root->left, tmp*10+root->val);
        if(root->right)
            dfs(root->right, tmp*10+root->val);
    }
    
    int sumNumbers(TreeNode* root) {
        dfs(root, 0);
        return sum;
    }
};

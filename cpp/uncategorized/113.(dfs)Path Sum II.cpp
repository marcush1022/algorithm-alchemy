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

class Solution {
public:
    void dfs(int sum, vector<int> tmp, vector<vector<int>> &ret, TreeNode *root, int target)
    {
        if(root==NULL)
            return ;
        tmp.push_back(root->val);
        if(sum+root->val==target && !root->left && !root->right)
        {
            ret.push_back(tmp);
            return ;
        }    
        dfs(sum+root->val, tmp, ret, root->left, target);
        dfs(sum+root->val, tmp, ret, root->right, target);
        tmp.pop_back();
    }
    
    void dfs2(vector<int> tmp, vector<vector<int>> &ret, TreeNode *root, int target)
    {
        if(root==NULL)
            return ;
        tmp.push_back(root->val);
        if(target==root->val && !root->left && !root->right)
        {
            ret.push_back(tmp);
            return ;
        }
        dfs2(tmp, ret, root->left, target-root->val);
        dfs2(tmp, ret, root->right, target-root->val);
    }
    
    vector<vector<int>> pathSum(TreeNode* root, int sum) {
        vector<vector<int>> ret;
        if(root==NULL)
            return ret;
        vector<int> tmp;
        //dfs(0, tmp, ret, root, sum);
        dfs2(tmp, ret, root, sum);
        return ret;
    }
};

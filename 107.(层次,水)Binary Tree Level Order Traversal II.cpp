/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values. 
(ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/

class Solution {
public:
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int>> ret;
        if(root==NULL)
            return ret;
        vector<int> path;
        queue<TreeNode*> que;
        int level=0, count=1;
        que.push(root);
        while(!que.empty())
        {
            level=0;
            path.clear();
            for(int i=0; i<count; i++)
            {
                root=que.front();
                path.push_back(root->val);
                que.pop();
                if(root->left)
                {
                    que.push(root->left);
                    level++;
                }
                if(root->right)
                {
                    que.push(root->right);
                    level++;
                }
            }
            count=level;
            ret.push_back(path);
        }
        reverse(ret.begin(), ret.end());
        return ret;
    }
};

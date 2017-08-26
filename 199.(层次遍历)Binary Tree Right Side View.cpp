/*
Given a binary tree, imagine yourself standing on the right side of it, return the 
values of the nodes you can see ordered from top to bottom.

For example:
Given the following binary tree,
   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
You should return [1, 3, 4].

层次遍历，返回每一层的最后一个元素
*/

class Solution {
public:
    vector<int> rightSideView(TreeNode *root)
    {
        vector<int> path;
        vector<int> ret;
        if(root==NULL)
            return ret;
        queue<TreeNode*> que;
        que.push(root);
        int level=0, count=1;
        while(!que.empty())
        {
            path.clear();
            level=0;
            for(int i=0; i<count; i++)
            {
                root=que.front();
                que.pop();
                path.push_back(root->val);
                if(root->left)
                {
                    que.push(root->left);
                    ++level;
                }
                if(root->right)
                {
                    que.push(root->right);
                    ++level;
                }
            }
            count=level;
            ret.push_back(path.back());
        }
        return ret;
    } 
};
















/*
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left 
for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
*/


class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>> ret;
		if(root==NULL)
			return ret;
		vector<int> path;
		queue<TreeNode*> que;
		int count=1, level=0;
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
		
		for(int i=0; i<ret.size(); i++)
		{
		    if(i&1)
		        reverse(ret[i].begin(), ret[i].end());
		}
		return ret;
    }
};










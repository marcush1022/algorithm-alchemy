/*
Given an integer n, generate all structurally unique BST's (binary search trees) that 
store values 1...n.

For example,
Given n = 3, your program should return all 5 unique BST's shown below.

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
   
Note: 由于1~n是升序列，因此建起来的树天然就是BST。
递归思想，依次选择根节点，对左右子序列再分别建树。

递推公式：dp[i] = dp[k] * dp[i-k-1] { 0<=k<=i-1 }
*/



class Solution {
public:
	vector<TreeNode*> doGenerate(int begin, int end)
	{
		vector<TreeNode*> ret;
		if(begin>end)
			ret.push_back(NULL);
		else if(begin==end)
		{
			TreeNode *root=new TreeNode(begin);
			ret.push_back(root);
		}
		else 
		{
			for(int i=begin; i<=end; i++)
			{
				vector<TreeNode*> left=doGenerate(begin, i-1);
				vector<TreeNode*> right=doGenerate(i+1, end);
				for(int l=0; l<left.size(); l++)
				{
					for(int r=0; r<right.size(); r++)
					{
						TreeNode *root=new TreeNode(i);
						root->left=left[l];
						root->right=right[r];
						ret.push_back(root);
					}
				}
			}
		}
		return ret;
	}
	
    vector<TreeNode*> generateTrees(int n) {
        vector<TreeNode*> ret;
        if(n==0)
            return ret;
        else
            return doGenerate(1, n);
    }
};

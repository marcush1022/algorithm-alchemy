/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth 
of the two subtrees of every node never differ by more than 1.
*/

class Solution {
public:
    int getDepth(TreeNode *root)
    {
        if(root==NULL)
            return 0;
        int nLeft=getDepth(root->left);
        int nRight=getDepth(root->right);
        return (nLeft>nRight) ? (nLeft+1) : (nRight+1);
    }
    
    bool isBalanced(TreeNode* root) {
        if(root==NULL)
            return true;
        int leftDepth=getDepth(root->left);
        int rightDepth=getDepth(root->right);
        int diff=leftDepth-rightDepth;
        if(diff>1 || diff<-1)
            return false;
        return isBalanced(root->left) && isBalanced(root->right);
    }
};

class Solution {
public:

	int isBalanced(TreeNode* root)
	{
		int depth=0;
		return judgeIsBalanced(root, &depth);
	}
	
	bool judgeIsBalanced(TreeNode* root, int* depth)
	{
		if(root==NULL)
		{
			*depth=0;
			return true;
		}
		
		int left, right;
		if(judgeIsBalanced(root->left, &left) && judgeIsBalanced(root->right, &right))
		{
			*depth=left-right;
			if(*depth<=1 && *depth>=-1)
			{
				*depth=1+(left>right ? left : right);
				return true;
			}
		}
		
		return false;
	}
};

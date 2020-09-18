/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

class Solution {
public:
    int getDepth(TreeNode *root)
    {
        if(root==NULL)
            return 0;
        int nLeft=getDepth(root->left);
        int nRight=getDepth(root->right);
        return nLeft>nRight ? (nLeft+1) : (nRight+1);
    }
    
    int maxDepth(TreeNode* root) {
        return getDepth(root);
    }
};

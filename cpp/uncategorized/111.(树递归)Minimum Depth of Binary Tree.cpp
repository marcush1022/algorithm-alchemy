/*
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

最小路径的定义是从根节点到叶子节点的最小路径。与最大路径不同的是需要考虑子树为0的情况，这时该节点的最小高度等于另一个子树的
最小高度。
*/

class Solution {
public:
    int minDepth(TreeNode* root) {
        if(root==NULL)
            return 0;
        int nLeft=minDepth(root->left);
        int nRight=minDepth(root->right);
        if(nLeft==0)
            return nRight+1;
        if(nRight==0)
            return nLeft+1;
        return min(nLeft+1, nRight+1);
    }
};

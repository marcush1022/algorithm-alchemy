/*
Given a complete binary tree, count the number of nodes.

这题的本质就是要找最深层有几个节点。那么要如何定位最后的那个节点呢。看到之前递归算法
写的 getDepth 方法，发现可以 O(logn) 去得到 depth，那就可以获取左右的路径的深度，
然后进行处理了：如果一样的话，这个子树就是满的，可以直接得到大小，如果不是一样的话，
就分别求左右子树的大小然后加1
*/

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int getLeftDepth(TreeNode *root)
    {
        int count=0;
        while(root)
        {
            ++count;
            root=root->left;
        }
        return count;
    }
    
    int getRightDepth(TreeNode *root)
    {
        int count=0;
        while(root)
        {
            ++count;
            root=root->right;
        }
        return count;
    }
    
    int countNodes(TreeNode* root) {
        if(root==NULL)
            return 0;
        int leftDepth=getLeftDepth(root);
        int rightDepth=getRightDepth(root);
        if(leftDepth==rightDepth)
            return (1<<leftDepth)-1;
        return countNodes(root->left)+countNodes(root->right)+1;
    }
};

/*
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note: 
You may assume k is always valid, 1 ? k ? BST's total elements.

*/

class Solution {
public:
    void inorder(vector<int> &ret, TreeNode *root)
    {
        if(root==NULL)
            return ;
        if(root->left)
            inorder(ret, root->left);
        ret.push_back(root->val);
        if(root->right)
            inorder(ret, root->right);
    }
    
    int kthSmallest(TreeNode* root, int k) {
        vector<int> ret;
        inorder(ret, root);
        return ret[k-1];
    }
};

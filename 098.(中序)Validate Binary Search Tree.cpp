/*
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
Example 1:
    2
   / \
  1   3
Binary tree [2,1,3], return true.
Example 2:
    1
   / \
  2   3
Binary tree [1,2,3], return false.
*/

class Solution {
public:
    void inorder(TreeNode *root, vector<int> &ret)
    {
        if(root==NULL)
            return ;
        if(root->left)
            inorder(root->left, ret);
        ret.push_back(root->val);
        if(root->right)
            inorder(root->right, ret);
    }
    
    bool isValidBST(TreeNode* root) {
        if(root==NULL)
            return true;
        vector<int> ret;
        inorder(root, ret);
        for(int i=0; i<ret.size()-1; i++)
            if(ret[i]>=ret[i+1])
                return false;
        return true;
    }
};

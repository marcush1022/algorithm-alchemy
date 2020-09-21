/*
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.
*/

class Solution {
public:
    TreeNode* build(vector<int> &preorder, vector<int> &inorder, int preIndex, int inIndex, int length)
    {
        if(length<1)
            return NULL;
        TreeNode *root=new TreeNode(preorder[preIndex]);
        for(int i=0; i<length; i++)
        {
            if(root->val==inorder[i+inIndex])
            {
                root->left=build(preorder, inorder, preIndex+1, inIndex, i);
                root->right=build(preorder, inorder, preIndex+i+1, inIndex+i+1, length-i-1);
                break;
            }
        }
        return root;
    }
    
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        TreeNode *root=build(preorder, inorder, 0, 0, preorder.size());
        return root;
    }
};

/*
Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.
*/

class Solution {
public:
    TreeNode* build(vector<int>& inorder, vector<int>& postorder, int inIndex, int postIndex, int len)
    {
        if(len<1)
            return NULL;
        TreeNode* root=new TreeNode(postorder[postIndex]);
        for(int i=0; i<len; i++)
        {
            if(root->val==inorder[inIndex-i])
            {
                root->left=build(inorder, postorder, inIndex-i-1, postIndex-i-1, len-i-1);
                root->right=build(inorder, postorder, inIndex, postIndex-1, i);
                break;
            }
        }
        return root;
    }
    
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        TreeNode *root=build(inorder, postorder, inorder.size()-1, inorder.size()-1, inorder.size());
        return root;
    }
};

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3
*/

class Solution {
public:
    bool judge(TreeNode *p, TreeNode *q)
    {
        if(p==NULL && q==NULL)
            return true;
        if(p==NULL || q==NULL)
            return false;
        if(p->val==q->val)
            return judge(p->left, q->right) && judge(p->right, q->left);
        else
            return false;
    }
    
    bool isSymmetric(TreeNode* root) {
        if(root==NULL)
            return true;
        return judge(root->left, root->right);
    }
};

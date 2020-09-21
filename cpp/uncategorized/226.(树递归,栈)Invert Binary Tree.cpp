/*
Invert a binary tree.

     4
   /   \
  2     7
 / \   / \
1   3 6   9
to
     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/

//递归
//逐层交换即可
class Solution {
public:
    void invert(TreeNode *root)
    {
        if(root==NULL)
            return ;
        TreeNode *tmp=root->left;
        root->left=root->right;
        root->right=tmp;
        invert(root->left);
        invert(root->right);
    }
    
    TreeNode* invertTree(TreeNode* root) {
        invert(root);
        return root;
    }
};

//非递归
//用栈，每次把栈顶元素的左右孩子入栈并交换
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        if(root==NULL)
            return root;
        stack<TreeNode*> stk;
        stk.push(root);
        while(!stk.empty())
        {
            TreeNode *p=stk.top();
            stk.pop();
            if(p)
            {
                stk.push(p->left);
                stk.push(p->right);
                swap(p->left, p->right);
            }
        }
        return root;
    }
};

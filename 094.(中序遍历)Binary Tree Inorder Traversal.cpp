
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> path;
        stack<TreeNode*> stk;
        while(root!=NULL || !stk.empty())
        {
            while(root!=NULL)
            {
                stk.push(root);
                root=root->left;
            }
            if(!stk.empty())
            {
                root=stk.top();
                path.push_back(root->val);
                stk.pop();
                root=root->right;
            }
        }
        return path;
    }
};

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
    
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> ret;
        inorder(ret, root);
        return ret;
    }
};

class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> ret;
        TreeNode *visited=NULL;
        stack<TreeNode*> stk;
        while(root || !stk.empty())
        {
            while(root)
            {
                stk.push(root);
                root=root->left;
            }
            root=stk.top();
            if(root->right==NULL || root->right==visited)
            {
                ret.push_back(root->val);
                visited=root;
                stk.pop();
                root=NULL;
            }
            else
                root=root->right;
        }
        return ret;
    }
};

//递归
class Solution {
public:
    void post(vector<int> &ret, TreeNode *root)
    {
        if(root==NULL)
            return ;
        post(ret, root->left);
        post(ret, root->right);
        ret.push_back(root->val);
    }
    
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> ret;
        post(ret, root);
        return ret;
    }
};

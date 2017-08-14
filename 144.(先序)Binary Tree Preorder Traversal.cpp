
class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> ret;
        if(root==NULL)
            return ret;
        stack<TreeNode*> stk;
        while(!stk.empty() || root)
        {
            if(root)
            {
                while(root)
                {
                    ret.push_back(root->val);
                    stk.push(root);
                    root=root->left;
                }
            }
            else
            {
                root=stk.top()->right;
                stk.pop();
            }
        }
        return ret;
    }
};

//递归
class Solution {
public:
    void dfs(TreeNode* root, vector<int> &ret)
    {
        if(root==NULL)
            return ;
        ret.push_back(root->val);
        dfs(root->left, ret);
        dfs(root->right, ret);
    }
    
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> ret;
        dfs(root, ret);
        return ret;
    }
};

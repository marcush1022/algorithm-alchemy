//Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

class Solution {
public:
    TreeNode* getTree(int left, int right, vector<int>& nums)
    {
        if(left<=right)   
        {
            int mid=(left+right)/2;
            TreeNode* root=new TreeNode(nums[mid]);
            root->left=getTree(left, mid-1, nums);
            root->right=getTree(mid+1, right, nums);
            return root;
        }
        else
            return NULL;
    }
    
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        TreeNode *root=getTree(0, nums.size()-1, nums);
        return root;
    }
};


class Solution {
public:
    TreeNode* getTree(int left, int right, vector<int>& nums)
    {
        if(left<right)
        {
            int mid=(left+right)/2;
            TreeNode* root=new TreeNode(nums[mid]);
            root->left=getTree(left, mid, nums);
            root->right=getTree(mid+1, right, nums);
            return root;
        }
        else
            return NULL;
    }
    
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        TreeNode *root=getTree(0, nums.size(), nums);
        return root;
    }
};

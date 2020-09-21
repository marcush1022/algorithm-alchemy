/*
Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.

Note: next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
*/

class BSTIterator {
    queue<int> que;
public:
    void inorder(TreeNode *root)
    {
        if(root==NULL)
            return ;
        inorder(root->left);
        que.push(root->val);
        inorder(root->right);
    }
    
    BSTIterator(TreeNode *root) {
        inorder(root);
    }

    //return whether we have a next smallest number
    bool hasNext() {
        return !que.empty();
    }

    //return the next smallest number
    int next() {
        int ret=que.front();
        que.pop();
        return ret;
    }
};

/**
 * Your BSTIterator will be called like this:
 * BSTIterator i = BSTIterator(root);
 * while (i.hasNext()) cout << i.next();
 */

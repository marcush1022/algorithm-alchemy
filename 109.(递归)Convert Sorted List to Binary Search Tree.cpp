/*
Given a singly linked list where elements are sorted in ascending order, 
convert it to a height balanced BST.
*/

class Solution {
public:
    TreeNode* sortedListToBST(ListNode* head) {
        if(head==NULL)
            return NULL;
        if(head->next==NULL)
            return new TreeNode(head->val);
        ListNode *fast=head, *slow=head, *pre=NULL;
        while(fast && fast->next)
        {
            fast=fast->next->next;
            pre=slow;
            slow=slow->next;
        }
        if(pre)
            pre->next=NULL;
        TreeNode *root=new TreeNode(slow->val);
        root->left=sortedListToBST(head);
        root->right=sortedListToBST(slow->next);
        
        return root;
    }
};


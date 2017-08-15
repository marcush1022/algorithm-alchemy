/*
Sort a linked list in O(n log n) time using constant space complexity.

Note: It means you should partition the list and use the merge sort.
*/

class Solution {
public:
    ListNode* merge(ListNode *l1, ListNode *l2)
    {
        if(l1==NULL)
            return l2;
        if(l2==NULL)
            return l1;
        ListNode *head;
        if(l1->val<l2->val)
        {
            head=l1;
            head->next=merge(l1->next, l2);
        }
        else
        {
            head=l2;
            head->next=merge(l1, l2->next);
        }
        return head;
    }
    
    ListNode* sortList(ListNode* head) {
        if(head==NULL || head->next==NULL)
            return head;
        ListNode *fast, *slow;
        fast=head, slow=head;
        while(fast->next && fast->next->next)
        {
            fast=fast->next->next;
            slow=slow->next;
        }
        ListNode *mhead=slow->next;
        slow->next=NULL;
        head=merge(sortList(head), sortList(mhead));
        return head;
    }
};

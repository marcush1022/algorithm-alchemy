/*
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You must do this in-place without altering the nodes' values.

For example,
Given {1,2,3,4}, reorder it to {1,4,2,3}.

思路：先分割成两部分，再合并
*/

class Solution {
public:
    ListNode* reverseList(ListNode *head)
    {
        if(head->next==NULL || head==NULL)
            return head;
        ListNode *pre=head, *p=head->next;
        ListNode *pNext;
        pre->next=NULL;
        while(p)
        {
            pNext=p->next;
            p->next=pre;
            pre=p;
            p=pNext;
        }
        return pre;
    }
    
    void reorderList(ListNode* head) {
        if(head==NULL || head->next==NULL)
            return ;
        ListNode *fast, *slow, *mhead;
        fast=head, slow=head;
        while(fast->next && fast->next->next)
        {
            fast=fast->next->next;
            slow=slow->next;
        }
        mhead=slow->next;
        slow->next=NULL;
        mhead=reverseList(mhead);
        while(head && mhead)
        {
            ListNode *tmp=head->next;
            head->next=mhead;
            mhead=mhead->next;
            head->next->next=tmp;
            head=tmp;
        }
    }
};







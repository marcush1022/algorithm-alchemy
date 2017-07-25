/*
Given a linked list and a value x, partition it such that all nodes less than x come 
before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

For example,
Given 1->4->3->2->5->2 and x = 3,
return 1->2->2->4->3->5.
*/

class Solution {
public:
    ListNode* partition(ListNode* head, int x) {
        if(head==NULL)
            return head;
        ListNode *nhead=new ListNode(0);
        nhead->next=head;
        ListNode *p=nhead, *pre, *q;
        while(p->next!=NULL && p->next->val<x)
            p=p->next;
        pre=p;
        q=p->next;
        while(p->next!=NULL)
        {
            if(p->next->val<x)
            {
                pre->next=p->next;
                pre=pre->next;
                p->next=p->next->next;
                pre->next=q;
            }
            else
                p=p->next;
        }
        return nhead->next;
    }
};

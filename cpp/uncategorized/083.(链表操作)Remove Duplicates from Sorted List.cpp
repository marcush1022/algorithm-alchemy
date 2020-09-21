/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.
*/

class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if(head==NULL || head->next==NULL)
            return head;
        ListNode *pre=head, *p;
        while(pre->next)
        {
            if(pre->next->val==pre->val)
            {
                p=pre->next;
                pre->next=p->next;
                p=NULL;
                delete p;
            }
            else
                pre=pre->next;
        }
        return head;
    }
};

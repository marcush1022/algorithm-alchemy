/*
Remove all elements from a linked list of integers that have value val.

Example
Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, val = 6
Return: 1 --> 2 --> 3 --> 4 --> 5
*/

class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        ListNode *nhead=new ListNode(0);
        nhead->next=head;
        ListNode *p=head, *pre=nhead;
        while(p)
        {
            if(p->val==val)
            {
                pre->next=p->next;
                p->next=NULL;
                free(p);
                p=pre->next;
            }
            else
            {
                pre=p;
                p=p->next;
            }
        }
        return nhead->next;
    }
};

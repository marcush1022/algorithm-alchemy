/*
Given a sorted linked list, delete all nodes that have duplicate numbers, 
leaving only distinct numbers from the original list.

For example,
Given 1->2->3->3->4->4->5, return 1->2->5.
Given 1->1->1->2->3, return 2->3.

p一直走到最后一个不重复的位置
*/

class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if(head==NULL || head->next==NULL)
            return head;
        ListNode *nhead=new ListNode(0);
        nhead->next=head;
        ListNode *pre=nhead, *p=nhead->next;
        while(p)
        {
            while(p->next!=NULL && pre->next->val==p->next->val)
                p=p->next;
            if(pre->next!=p)
            {
                pre->next=p->next;
                p->next=NULL;
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
